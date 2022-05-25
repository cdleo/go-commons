package sqlcommons

const (
	Oracle_DriverName     = "oracle"
	OracleTNS_DriverName  = "oracle-tns"
	Postgresql_DriverName = "postgresql"
	SQLite3_DriverName    = "sqlite3"
	MockDB_DriverName     = "mockdb"
)

type sqlEngineAdapter interface {
	QueryTranslator(fromEngine string, query string) string
	ErrorHandler(err error) error
}
