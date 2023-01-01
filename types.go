package wpgx

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

// PostExecFunc is the function that must be ran after a successful CRUD.
// NOTE: context should have been captured into the function.
type PostExecFunc = func() error

// TxFunc is the body of a transaction.
type TxFunc = func(WGConn) (any, error)

// WGConn is the abstraction over wrapped connections and transactions.
type WGConn interface {
	WQuery(
		ctx context.Context, name string, unprepared string, args ...interface{}) (pgx.Rows, error)
	WQueryRow(
		ctx context.Context, name string, unprepared string, args ...interface{}) pgx.Row
	WExec(
		ctx context.Context, name string, unprepared string, args ...interface{}) (pgconn.CommandTag, error)
	PostExec(f PostExecFunc) error
}
