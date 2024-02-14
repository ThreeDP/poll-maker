package main

import (
	"context"
	"database/sql"
	"github.com/ThreeDP/poll-maker/db"
	"github.com/ThreeDP/poll-maker/route"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func connectDB() (*db.Queries, error) {
	dbConnection, err := sql.Open("postgres", "postgres://user:user@localhost:5432/poll-maker?sslmode=disable")
	if err != nil {
		return nil, err
	}
	return db.New(dbConnection), nil
}

func setupRouter() *gin.Engine {
	r := gin.Default()
	ctx := context.Background()
	dt, err := connectDB()
	if err != nil {
		panic(err.Error())
	}
	
	r.POST("/poll", func(c *gin.Context) {
		route.CreatePollRequest(c, dt, ctx)
	})

	r.GET("poll/:pollId", func(c *gin.Context) {
		route.GetPollRequest(c, dt, ctx)
	})

	return r
}

func main() {
	r := setupRouter()
	err := r.Run(":8080")
	if err != nil {
		return
	}
}