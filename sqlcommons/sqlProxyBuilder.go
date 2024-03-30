package sqlcommons

import (
	"database/sql"
	"database/sql/driver"

	"github.com/cdleo/go-commons/logger"
)

type MockSQLEngineAdapter interface {
	SQLEngineAdapter

	PatchBegin(err error)
	PatchCommit(err error)
	PatchRollback(err error)

	PatchExec(query string, err error, args ...driver.Value)
	PatchQuery(query string, columns []string, values []driver.Value, err error, args ...driver.Value)
	PatchQueryRow(query string, result map[string]string, err error)
}

type SQLEngineAdapter interface {
	Open(logger logger.Logger, translator SQLSyntaxTranslator) (*sql.DB, error)
}

type SQLSyntaxTranslator interface {
	Translate(query string) string
	ErrorHandler(err error) error
}
