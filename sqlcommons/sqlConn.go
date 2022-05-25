package sqlcommons

import "database/sql"

type EngineAdapter interface {
	Open() (*sql.DB, error)
	ErrorHandler(err error) error
}

type SQLSintaxTranslator interface {
	Translate(query string) string
}
