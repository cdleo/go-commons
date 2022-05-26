package sqlcommons

import (
	"context"
	"database/sql"
)

type SQLClient interface {
	Open() error
	Close()
	IsOpen() error

	Begin() (SQLTx, error)
	BeginTx(ctx context.Context, opts *sql.TxOptions) (SQLTx, error)

	SQLTransactionable
}
