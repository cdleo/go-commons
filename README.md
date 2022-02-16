# GO-FACADES

The repository go-facades it's a set of useful interfaces of widely used subsystems (like logging).

It's based on the principles of Clean Architecture, due to the underlying business rules (use cases), have dependencies of this interfaces instead of the concrete implementation.

## Looger

Abstraction to normalize the implementation of several loggers without change the client:
```go
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
```

**Log levels**
This is the list of log levels with their `weights` and meanings:
- **0 disabled**: No log
- **1 show**: Very low frequency messages that should always be displayed (such as copyright)
- **2 fatal**: Errors that cause the application to stop
- **3 error**: Errors that don't stop the app
- **4 warning**: Alert conditions that not generate an error
- **5 info**: Important details, such as the version of the components when lifting
- **6 business**: Important business details, such as the identification of the transaction and the user who carried it out
- **7 message**: Details of all the Business request and response exchanged with the outside
- **8 debug**: Detailed execution's information
- **9 query**: Detail of executed querys with their input parameters
- **10 trace**: Maximum level of detail, such as the values returned by the querys, http trace, etc.

**Note**: Remember that only messages with levels <= to the value set on function `SetLogLevel` will be show.