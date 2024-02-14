package route

import (
	"context"
	"github.com/ThreeDP/poll-maker/db"
	_ "github.com/lib/pq"
	"database/sql"
)

type IQueries interface {
	CreatePoll(context.Context, db.CreatePollParams) error
	GetPoll(context.Context, string) (string, error)
	WithTx(*sql.Tx) *db.Queries
}