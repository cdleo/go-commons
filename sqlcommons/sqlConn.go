package sqlcommons

import "database/sql"

type SqlConn interface {
	Open() error
	Close()
	IsOpen() error
	sql.DB
}
