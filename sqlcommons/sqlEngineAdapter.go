package sqlcommons

import "database/sql"

const (
	Oracle_DriverName     = "oracle"
	OracleTNS_DriverName  = "oracle-tns"
	Postgresql_DriverName = "postgres"
	SQLite3_DriverName    = "sqlite3"
	MockDB_DriverName     = "mockdb"
)

type SQLSintaxTranslator interface {
	Translate(query string) string
}

type EngineAdapter interface {
	Open() (*sql.DB, error)
	ErrorHandler(err error) error
}
