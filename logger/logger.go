package logger

import (
	"io"
	"time"
)

type Logger interface {
	SetLogLevel(level LogLevel) error
	SetOutput(w io.Writer)
	SetTimestampFunc(f func() time.Time)

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
