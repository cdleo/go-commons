/*
Package logger its the package of the Logger Facade
*/

package logger

import (
	"io"
	"time"
)

type Logger interface {
	//Sets the current log level. (e.g. "debug")
	SetLogLevel(level string) error
	//Sets the log writer. (e.g. os.Stdout)
	SetOutput(w io.Writer)
	//Sets the function to write the log's timestamp. (e.g. time.Now)
	SetTimestampFunc(f func() time.Time)

	//Includes the ref field on the related log msg call
	WithRefID(refID string) Logger

	Show(msg string)
	Showf(msg string, v ...interface{})

	Fatal(err error, msg string)
	Fatalf(err error, msg string, v ...interface{})

	Error(err error, msg string)
	Errorf(err error, msg string, v ...interface{})

	Warn(msg string)
	Warnf(msg string, v ...interface{})

	Info(msg string)
	Infof(msg string, v ...interface{})

	Bus(msg string)
	Busf(msg string, v ...interface{})

	Msg(msg string)
	Msgf(msg string, v ...interface{})

	Dbg(msg string)
	Dbgf(msg string, v ...interface{})

	Qry(msg string)
	Qryf(msg string, v ...interface{})

	Trace(msg string)
	Tracef(msg string, v ...interface{})
}
