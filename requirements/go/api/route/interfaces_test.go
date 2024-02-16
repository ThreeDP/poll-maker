package route

import (
	"context"
	"github.com/ThreeDP/poll-maker/db"
	_ "github.com/lib/pq"
	"database/sql"
	"errors"
)

type TQueries struct {
	dbError bool
}

func (tq TQueries) CreatePoll(ctx context.Context, dt db.CreatePollParams) error {
	if tq.dbError == false {
		return nil
	}
	return errors.New("error on insert poll")
}

func (tq TQueries) WithTx(sql *sql.Tx) *db.Queries {
	return &db.Queries{}
}

func (tq TQueries) GetPoll(ctx context.Context, id string) (string, error) {
	if tq.dbError == false {
		if id == "75711fb4-565c-4232-b491-28175a8cd8e9" {
			return "Olá bea", nil
		}
	}
	return "test", errors.New("erro on get poll")
}

func (tq TQueries) CreateUser(ctx context.Context, arg db.CreateUserParams) error {
	if tq.dbError == true {
		return errors.New("error")
	}
	return nil
}