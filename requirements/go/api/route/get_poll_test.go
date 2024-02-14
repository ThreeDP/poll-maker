package route

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ThreeDP/poll-maker/db"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type TBody struct {
	Title string `json:"title"`
}

type TQueries struct {
	dbError bool
}

func (tq TQueries) CreatePoll(ctx context.Context, dt db.CreatePollParams) error {
	if tq.dbError == false {
		return nil
	}
	return errors.New("error on insert poll")
}

func (tq TQueries) WithTx(sql *sql.Tx) *db.Queries {
	return &db.Queries{}
}

func (tq TQueries) GetPoll(ctx context.Context, id string) (string, error) {
	if tq.dbError == false {
		if id == "75711fb4-565c-4232-b491-28175a8cd8e9" {
			return "Olá bea", nil
		}
	}
	return "test", errors.New("erro on get poll")
}

func TestGetPoll(t *testing.T) {
	gin.SetMode(gin.ReleaseMode)
	ctx := context.Background()

	t.Run("Test get a valid poll", func(t *testing.T) {
		dt := TQueries{dbError: false}
		w := httptest.NewRecorder()
		cG, _ := gin.CreateTestContext(w)
		cG.Request = httptest.NewRequest(http.MethodGet, "/poll/:pollId", nil)
		cG.Params = append(cG.Params, gin.Param{
				Key: "pollId",
				Value: "75711fb4-565c-4232-b491-28175a8cd8e9",
			})
		
		GetPollRequest(cG, dt, ctx)

		if w.Code != http.StatusOK {
			t.Errorf("Expected Status [ 200 ], but has %d", w.Code)
		}

		var responseStruct TBody
		if err := json.Unmarshal([]byte(w.Body.Bytes()), &responseStruct); err != nil {
			t.Errorf("decode Error: %s", err)
		}

		want := "Olá bea"
		if responseStruct.Title != want {
			t.Errorf("expected %s, but has %s", want, responseStruct.Title)
		}
	})

	t.Run("Test get a poll with a invalid uuid", func(t *testing.T) {
		dt := TQueries{dbError: false}
		w := httptest.NewRecorder()
		cG, _ := gin.CreateTestContext(w)
		cG.Request = httptest.NewRequest(http.MethodGet, "/poll/:pollId", nil)
		cG.Params = append(cG.Params, gin.Param{
			Key: "pollId",
			Value: "invalid-uuid",
		})
		
		GetPollRequest(cG, dt, ctx)

		if w.Code != http.StatusBadRequest {
			t.Errorf("Expected Status [ %d ], but has %d", http.StatusBadRequest, w.Code)
		}

		var responseStruct TBody
		if err := json.Unmarshal([]byte(w.Body.Bytes()), &responseStruct); err != nil {
			t.Errorf("decode Error: %s", err)
		}
	})

	t.Run("Test get a poll without a pollId", func(t *testing.T) {
		dt := TQueries{dbError: false}
		w := httptest.NewRecorder()
		cG, _ := gin.CreateTestContext(w)
		cG.Request = httptest.NewRequest(http.MethodGet, "/poll/:pollId", nil)
		
		GetPollRequest(cG, dt, ctx)

		if w.Code != http.StatusBadRequest {
			t.Errorf("Expected Status [ %d ], but has %d", http.StatusBadRequest, w.Code)
		}

		var responseStruct TBody
		if err := json.Unmarshal([]byte(w.Body.Bytes()), &responseStruct); err != nil {
			t.Errorf("decode Error: %s", err)
		}
	})

	t.Run("Test get a query that doesn't exist", func(t *testing.T) {
		dt := TQueries{dbError: false}
		w := httptest.NewRecorder()
		cG, _ := gin.CreateTestContext(w)
		cG.Request = httptest.NewRequest(http.MethodGet, "/poll/:pollId", nil)
		cG.Params = append(cG.Params, gin.Param{
			Key: "pollId",
			Value: "75711fb4-565c-4232-b491-28175a8cd8e6",
		})

		GetPollRequest(cG, dt, ctx)

		if w.Code != http.StatusNotFound {
			t.Errorf("Expected Status [ %d ], but has %d", http.StatusNotFound, w.Code)
		}

		var responseStruct TBody
		if err := json.Unmarshal([]byte(w.Body.Bytes()), &responseStruct); err != nil {
			t.Errorf("decode Error: %s", err)
		}
	})
}