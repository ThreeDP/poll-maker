package main

import (
	"github.com/ThreeDP/poll-maker/route"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	err := route.ConnectDB()
	if err != nil {
		panic(err.Error())
	}
	
	r.POST("/poll/create", route.CreatePollRequest)
	r.GET("poll/:pollId", route.GetPollRequest)

	r.POST("/user/create", route.CreateUserRequest)

	return r
}

func main() {
	r := setupRouter()
	err := r.Run(":8080")
	if err != nil {
		return
	}
}