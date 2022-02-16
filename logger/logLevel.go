/*
Package logger its the package of the Logger Facade
*/

package logger

import (
	"fmt"
	"strings"
)

// Level defines log levels.
type LogLevel int8

const (
	// No log
	LogLevel_Disabled LogLevel = iota

	// Very low frequency messages that should always be displayed (such as copyright)
	LogLevel_Show

	// Errors that cause the application to stop
	LogLevel_Fatal

	// Errors that don't stop the app
	LogLevel_Error

	// Alert conditions that not generate an error
	LogLevel_Warning

	// Important details, such as the version of the components when lifting
	LogLevel_Info

	// Important business details, such as the identification of the transaction and the user who carried it out
	LogLevel_Business

	// Details of all the Business request and response exchanged with the outside
	LogLevel_Message

	//  Detailed execution's information
	LogLevel_Debug

	// Detail of executed querys with their input parameters
	LogLevel_Query

	// Maximum level of detail, such as the values returned by the querys, http trace, etc.
	LogLevel_Trace
)

func NewLogLevel(level string) (LogLevel, error) {

	level = strings.ToLower(level)
	switch level {
	case LogLevel_Disabled.String():
		return LogLevel_Disabled, nil
	case LogLevel_Show.String():
		return LogLevel_Show, nil
	case LogLevel_Fatal.String():
		return LogLevel_Fatal, nil
	case LogLevel_Error.String():
		return LogLevel_Error, nil
	case LogLevel_Warning.String():
		return LogLevel_Warning, nil
	case LogLevel_Info.String():
		return LogLevel_Info, nil
	case LogLevel_Business.String():
		return LogLevel_Business, nil
	case LogLevel_Message.String():
		return LogLevel_Message, nil
	case LogLevel_Debug.String():
		return LogLevel_Debug, nil
	case LogLevel_Query.String():
		return LogLevel_Query, nil
	case LogLevel_Trace.String():
		return LogLevel_Trace, nil
	default:
		return LogLevel_Info, fmt.Errorf("unknown LogLevel [%s]", level)
	}
}

func (l LogLevel) String() string {

	switch l {
	case LogLevel_Disabled:
		return "disabled"
	case LogLevel_Show:
		return "show"
	case LogLevel_Fatal:
		return "fatal"
	case LogLevel_Error:
		return "error"
	case LogLevel_Warning:
		return "warn"
	case LogLevel_Info:
		return "info"
	case LogLevel_Business:
		return "business"
	case LogLevel_Message:
		return "message"
	case LogLevel_Debug:
		return "debug"
	case LogLevel_Query:
		return "query"
	case LogLevel_Trace:
		return "trace"
	default:
		return ""
	}
}

func (me LogLevel) IsLogAllowed(other LogLevel) bool {
	return (me >= other)
}
