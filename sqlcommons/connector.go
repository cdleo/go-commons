package sqlcommons

import (
	"database/sql"
	"database/sql/driver"

	"github.com/cdleo/go-commons/logger"
)

type SQLConnector interface {
	Open(logger logger.Logger, translator SQLAdapter) (*sql.DB, error)
	GetNextSequenceQuery(sequenceName string) string
}

type MockSQLConnector interface {
	SQLConnector

	PatchBegin(err error)
	PatchCommit(err error)
	PatchRollback(err error)

	PatchExec(query string, err error, args ...driver.Value)
	PatchQuery(query string, columns []string, values []driver.Value, err error, args ...driver.Value)
	PatchQueryRow(query string, result map[string]string, err error)
}
