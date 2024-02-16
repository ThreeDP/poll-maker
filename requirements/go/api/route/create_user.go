package route

import (
	"net/http"

	"github.com/ThreeDP/poll-maker/db"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CreateUserBody struct {
	ID string 
	Name string `json:"name" binding:"required"`
	Surname string `json:"surname" binding:"required"`
	Nickname string `json:"nickname" binding:"required"`
	Email string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

func CreateUserRequest(cG *gin.Context) {
	var userBody CreateUserBody
	if err := cG.ShouldBindJSON(&userBody); err != nil {
		cG.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	userBody.ID = uuid.New().String()
	err := DBConnc.CreateUser(cG, db.CreateUserParams(userBody))
	if err != nil {
		cG.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return 
	}
	cG.JSON(http.StatusCreated, nil)
}