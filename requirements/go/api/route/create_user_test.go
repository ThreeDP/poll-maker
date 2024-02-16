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
			t.Errorf("Expected Status [ %d ], but has %d", http.StatusCreated, w.Code)
		}
	})

	t.Run("Test Create a user without name param", func (t *testing.T) {
		reqBody := `{
			"surname": "Smith",
			"nickname": "joSmith",
			"email": "JoSmith@poll-maker.com",
			"password": "joSmith12345"}`
		w := httptest.NewRecorder() 
		cG, _ := gin.CreateTestContext(w)
		cG.Request = httptest.NewRequest(http.MethodPost, "/user/create", bytes.NewBuffer([]byte(reqBody)))
		
		CreateUserRequest(cG)

		if w.Code != http.StatusBadRequest {
			t.Errorf("Expected Status [ %d ], but has %d", http.StatusBadRequest, w.Code)
		}
	})

	t.Run("Test Create a user without surname param", func (t *testing.T) {
		reqBody := `{
			"name": "John",
			"nickname": "joSmith",
			"email": "JoSmith@poll-maker.com",
			"password": "joSmith12345"}`
		w := httptest.NewRecorder() 
		cG, _ := gin.CreateTestContext(w)
		cG.Request = httptest.NewRequest(http.MethodPost, "/user/create", bytes.NewBuffer([]byte(reqBody)))
		
		CreateUserRequest(cG)

		if w.Code != http.StatusBadRequest {
			t.Errorf("Expected Status [ %d ], but has %d", http.StatusBadRequest, w.Code)
		}
	})

	t.Run("Test Create a user without nickname param", func (t *testing.T) {
		reqBody := `{
			"name": "John",
			"surname": "Smith",,
			"email": "JoSmith@poll-maker.com",
			"password": "joSmith12345"}`
		w := httptest.NewRecorder() 
		cG, _ := gin.CreateTestContext(w)
		cG.Request = httptest.NewRequest(http.MethodPost, "/user/create", bytes.NewBuffer([]byte(reqBody)))
		
		CreateUserRequest(cG)

		if w.Code != http.StatusBadRequest {
			t.Errorf("Expected Status [ %d ], but has %d", http.StatusBadRequest, w.Code)
		}
	})

	t.Run("Test Create a user without email param", func (t *testing.T) {
		reqBody := `{
			"name": "John",
			"surname": "Smith",
			"nickname": "joSmith",
			"password": "joSmith12345"}`
		w := httptest.NewRecorder() 
		cG, _ := gin.CreateTestContext(w)
		cG.Request = httptest.NewRequest(http.MethodPost, "/user/create", bytes.NewBuffer([]byte(reqBody)))
		
		CreateUserRequest(cG)

		if w.Code != http.StatusBadRequest {
			t.Errorf("Expected Status [ %d ], but has %d", http.StatusBadRequest, w.Code)
		}
	})

	// t.Run("Test Create a user with a invalid email param", func (t *testing.T) {
	// 	reqBody := `{
	// 		"name": "John",
	// 		"surname": "Smith",
	// 		"nickname": "joSmith",
	// 		"email": "",
	// 		"password": "joSmith12345",
	// 	}`
	// 	w := httptest.NewRecorder() 
	// 	cG, _ := gin.CreateTestContext(w)
	// 	cG.Request = httptest.NewRequest(http.MethodPost, "/user/create", bytes.NewBuffer([]byte(reqBody)))
		
	// 	CreateUserRequest(cG)

	// 	if w.Code != http.Status {
	// 		t.Errorf("Expected Status [ %d ], but has %d", http.StatusBadRequest, w.Code)
	// 	}
	// })

	t.Run("Test Create a user without password param", func (t *testing.T) {
		reqBody := `{
			"name": "John",
			"surname": "Smith",
			"nickname": "joSmith",
			"email": "JoSmith@poll-maker.com",
		}`
		w := httptest.NewRecorder() 
		cG, _ := gin.CreateTestContext(w)
		cG.Request = httptest.NewRequest(http.MethodPost, "/user/create", bytes.NewBuffer([]byte(reqBody)))
		
		CreateUserRequest(cG)

		if w.Code != http.StatusBadRequest {
			t.Errorf("Expected Status [ %d ], but has %d", http.StatusBadRequest, w.Code)
		}
	})
}

func TestUserCreateServerError(t *testing.T) {
	DBConnc = TQueries{dbError: true}
	t.Run("Test a server error", func (t *testing.T) {
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

		if w.Code != http.StatusInternalServerError {
			t.Errorf("Expected Status [ %d ], but has %d", http.StatusInternalServerError, w.Code)
		}
	})
}