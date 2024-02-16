package route

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	"github.com/ThreeDP/poll-maker/db"
	_ "github.com/lib/pq"
)

type IQueries interface {
	CreatePoll(context.Context, db.CreatePollParams) error
	GetPoll(context.Context, string) (string, error)
	CreateUser(ctx context.Context, arg db.CreateUserParams) error
	WithTx(*sql.Tx) *db.Queries
}

var DBConnc IQueries

func ConnectDB() error {
	dataSource := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable",
		os.Getenv("DBUSER"),
		os.Getenv("DBPASS"),
		os.Getenv("DBHOST"),
		os.Getenv("DBNAME"),
	)
	dbConnection, err := sql.Open("postgres", dataSource)
	if err != nil {
		return err
	}
	DBConnc = db.New(dbConnection)
	return nil
}
