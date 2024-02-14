package route

import (
	"context"
	_ "github.com/lib/pq"
	"net/http"
	"github.com/ThreeDP/poll-maker/db"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type body struct {
	ID		string
	Title	string `json:"title" binding:"required"`
}

func CreatePollRequest(cG *gin.Context, dt IQueries, ctx context.Context) {
	var reqBody body

	if err := cG.ShouldBindJSON(&reqBody); err != nil {
		cG.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	reqBody.ID = uuid.New().String()
	err := dt.CreatePoll(ctx, db.CreatePollParams(reqBody))
	if err != nil {
		cG.JSON(http.StatusInternalServerError, gin.H{"error": http.StatusText(http.StatusInternalServerError)})
		return
	}
	cG.JSON(http.StatusCreated, gin.H{
		"id": reqBody.ID,
	})
}