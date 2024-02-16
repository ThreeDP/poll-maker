package route

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type TBody struct {
	Title string `json:"title"`
}

func TestGetPoll(t *testing.T) {
	gin.SetMode(gin.ReleaseMode)
	DBConnc = TQueries{dbError: false}
	
	t.Run("Test get a valid poll", func(t *testing.T) {
		w := httptest.NewRecorder()
		cG, _ := gin.CreateTestContext(w)
		cG.Request = httptest.NewRequest(http.MethodGet, "/poll/:pollId", nil)
		cG.Params = append(cG.Params, gin.Param{
				Key: "pollId",
				Value: "75711fb4-565c-4232-b491-28175a8cd8e9",
			})
		
		GetPollRequest(cG)

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
		w := httptest.NewRecorder()
		cG, _ := gin.CreateTestContext(w)
		cG.Request = httptest.NewRequest(http.MethodGet, "/poll/:pollId", nil)
		cG.Params = append(cG.Params, gin.Param{
			Key: "pollId",
			Value: "invalid-uuid",
		})
		
		GetPollRequest(cG)

		if w.Code != http.StatusBadRequest {
			t.Errorf("Expected Status [ %d ], but has %d", http.StatusBadRequest, w.Code)
		}
		var responseStruct TBody
		if err := json.Unmarshal([]byte(w.Body.Bytes()), &responseStruct); err != nil {
			t.Errorf("decode Error: %s", err)
		}
	})

	t.Run("Test get a poll without a pollId", func(t *testing.T) {
		w := httptest.NewRecorder()
		cG, _ := gin.CreateTestContext(w)
		cG.Request = httptest.NewRequest(http.MethodGet, "/poll/:pollId", nil)
		
		GetPollRequest(cG)

		if w.Code != http.StatusBadRequest {
			t.Errorf("Expected Status [ %d ], but has %d", http.StatusBadRequest, w.Code)
		}
		var responseStruct TBody
		if err := json.Unmarshal([]byte(w.Body.Bytes()), &responseStruct); err != nil {
			t.Errorf("decode Error: %s", err)
		}
	})

	t.Run("Test get a query that doesn't exist", func(t *testing.T) {
		w := httptest.NewRecorder()
		cG, _ := gin.CreateTestContext(w)
		cG.Request = httptest.NewRequest(http.MethodGet, "/poll/:pollId", nil)
		cG.Params = append(cG.Params, gin.Param{
			Key: "pollId",
			Value: "75711fb4-565c-4232-b491-28175a8cd8e6",
		})

		GetPollRequest(cG)

		if w.Code != http.StatusNotFound {
			t.Errorf("Expected Status [ %d ], but has %d", http.StatusNotFound, w.Code)
		}
		var responseStruct TBody
		if err := json.Unmarshal([]byte(w.Body.Bytes()), &responseStruct); err != nil {
			t.Errorf("decode Error: %s", err)
		}
	})
}