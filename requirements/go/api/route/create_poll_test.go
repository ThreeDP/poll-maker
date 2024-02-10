package route

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/gin-gonic/gin"
	"github.com/ThreeDP/poll-maker/db"
	_ "github.com/lib/pq"
	"database/sql"
	"bytes"
	"encoding/json"
	"github.com/google/uuid"
	"errors"
)

type TBody struct {
	Id string `json:"id"`
}

type TQueriesCreatePoll struct {
	dbError bool
}

func (tq TQueriesCreatePoll) CreatePoll(ctx context.Context, dt db.CreatePollParams) error {
	if tq.dbError == false {
		return nil
	}
	return errors.New("error on insert poll")
}

func (tq TQueriesCreatePoll) WithTx(sql *sql.Tx) *db.Queries {
	return &db.Queries{}
}

func TestCreatePoll(t *testing.T) {
	gin.SetMode(gin.ReleaseMode)
	ctx := context.Background()
	
	t.Run("Test pass a correct Json Inform", func (t *testing.T) {
		reqBody := `{"title": "Test Poll"}`
		dt := TQueriesCreatePoll{dbError: false}
		w := httptest.NewRecorder() 
		cG, _ := gin.CreateTestContext(w)
		cG.Request = httptest.NewRequest(http.MethodPost, PollRoute, bytes.NewBuffer([]byte(reqBody)))
		CreatePollRequest(cG, dt, ctx)

		if w.Code != http.StatusOK {
			t.Errorf("Expected Status [ 200 ], but has %d", w.Code)
		}

		var responseStruct TBody
		if err := json.Unmarshal([]byte(w.Body.Bytes()), &responseStruct); err != nil {
			t.Errorf("decode Error: %s", err)
		}

		_, err := uuid.Parse(responseStruct.Id)
		if err != nil {
			t.Errorf("Returned ID is not a valid UUID: %s", responseStruct.Id)
		}
	})

	t.Run("Test don't pass title", func(t *testing.T) {
		reqBody := `{}`
		dt := TQueriesCreatePoll{dbError: false}
		w := httptest.NewRecorder() 
		cG, _ := gin.CreateTestContext(w)
		cG.Request = httptest.NewRequest(http.MethodPost, PollRoute, bytes.NewBuffer([]byte(reqBody)))
		CreatePollRequest(cG, dt, ctx)
		if w.Code != http.StatusBadRequest {
			t.Errorf("Expected Status [ 200 ], but has %d", w.Code)
		}
	})

	t.Run("Test BD error", func(t *testing.T) {
		reqBody := `{"title": "Test Poll"}`
		dt := TQueriesCreatePoll{dbError: true}
		w := httptest.NewRecorder() 
		cG, _ := gin.CreateTestContext(w)
		cG.Request = httptest.NewRequest(http.MethodPost, PollRoute, bytes.NewBuffer([]byte(reqBody)))
		CreatePollRequest(cG, dt, ctx)

		if w.Code != http.StatusInternalServerError {
			t.Errorf("Expected Status [ 200 ], but has %d", w.Code)
		}
	})
}