package route

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

type tRequest struct {
	PollId string `uri:"pollId" binding:"required,uuid"`
}

func GetPollRequest(cG *gin.Context, dt IQueries, ctx context.Context) {
	var params tRequest
	if err := cG.ShouldBindUri(&params); err != nil {
		cG.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	title, err := dt.GetPoll(ctx, params.PollId)
	if err != nil {
		cG.JSON(http.StatusNotFound, gin.H{"error": err})
		return
	}
	cG.JSON(http.StatusOK, gin.H{"title": title})
}