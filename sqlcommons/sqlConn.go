package sqlcommons

import "database/sql"

type SqlConn interface {
	Open() error
	Close()
	IsOpen() error
	*sql.DB
}

type EngineAdapter interface {
	Open() (*sql.DB, error)
	ErrorHandler(err error) error
}

type SQLSintaxTranslator interface {
	Translate(query string) string
}
