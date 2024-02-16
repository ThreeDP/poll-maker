package route

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"bytes"
)

func TestCreateUser(t *testing.T) {
	gin.SetMode(gin.ReleaseMode)
	DBConnc = TQueries{dbError: false}

	t.Run("Test Create a user passing the correct informations", func (t *testing.T) {
		reqBody := `{
			"name": "John",
			"surname": "Smith",
			"nickname": "joSmith",
			"email": "JoSmith@poll-maker.com",
			"password": "joSmith12345"}`
		w := httptest.NewRecorder() 
		cG, _ := gin.CreateTestContext(w)
		cG.Request = httptest.NewRequest(http.MethodPost, "/user/create", bytes.NewBuffer([]byte(reqBody)))
		
		CreateUserRequest(cG)

		if w.Code != http.StatusCreated {
			t.Errorf("Expected Status [ 201 ], but has %d", w.Code)
		}
	})
}