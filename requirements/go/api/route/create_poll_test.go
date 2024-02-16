package route

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"bytes"
	"encoding/json"
	"github.com/google/uuid"
)

type tBody struct {
	Id string `json:"id"`
}

func TestCreatePoll(t *testing.T) {
	gin.SetMode(gin.ReleaseMode)
	DBConnc = TQueries{dbError: false}
	
	t.Run("Test pass a correct Json Inform", func (t *testing.T) {
		reqBody := `{"title": "Test Poll"}`
		w := httptest.NewRecorder() 
		cG, _ := gin.CreateTestContext(w)
		cG.Request = httptest.NewRequest(http.MethodPost, "/poll/create", bytes.NewBuffer([]byte(reqBody)))
		CreatePollRequest(cG)

		if w.Code != http.StatusCreated {
			t.Errorf("Expected Status [ 201 ], but has %d", w.Code)
		}

		var responseStruct tBody
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
		w := httptest.NewRecorder() 
		cG, _ := gin.CreateTestContext(w)
		cG.Request = httptest.NewRequest(http.MethodPost, "/poll/create", bytes.NewBuffer([]byte(reqBody)))
		CreatePollRequest(cG)
		if w.Code != http.StatusBadRequest {
			t.Errorf("Expected Status [ 400 ], but has %d", w.Code)
		}
	})

	DBConnc = TQueries{dbError: true}
	t.Run("Test BD error", func(t *testing.T) {
		reqBody := `{"title": "Test Poll"}`
		w := httptest.NewRecorder() 
		cG, _ := gin.CreateTestContext(w)
		cG.Request = httptest.NewRequest(http.MethodPost, "/poll/create", bytes.NewBuffer([]byte(reqBody)))
		CreatePollRequest(cG)

		if w.Code != http.StatusInternalServerError {
			t.Errorf("Expected Status [ 500 ], but has %d", w.Code)
		}
	})
}