package sqlcommons

import (
	"context"
	"database/sql"
)

//Drivers
const (
	Oracle_DriverName     = "oracle"
	OracleTNS_DriverName  = "oracle-tns"
	Postgresql_DriverName = "postgres"
	SQLite3_DriverName    = "sqlite3"
	MockDB_DriverName     = "mockdb"
)

type SQLEngineAdapter interface {
	Open() (*sql.DB, error)
	ErrorHandler(err error) error
}

type SQLSyntaxTranslator interface {
	Translate(query string) string
}

type SQLClient interface {
	Open() error
	Close()
	IsOpen() error

	Begin() (SQLTx, error)
	BeginTx(ctx context.Context, opts *sql.TxOptions) (SQLTx, error)

	SQLTransactionable
}
