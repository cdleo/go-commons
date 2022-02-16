package logger

import (
	"fmt"
	"strings"
)

// Level defines log levels.
type LogLevel int8

const (
	// Disabled disables the logger.
	LogLevel_Disabled LogLevel = iota

	// Mensajes de muy baja frecuencia que se deben mostrar siempre (como el copyright)
	LogLevel_Show

	// Errores que provocan el detenimiento de la aplicación
	LogLevel_Fatal

	// Errores que no detienen la app
	LogLevel_Error

	// Condiciones a tener en cuenta pero que no generan error
	LogLevel_Warning

	// Detalles importantes, como la versión de los componentes al levantar
	LogLevel_Info

	// Detalles importantes de negocio, como la identificación de la transacción y el user que la realizó
	LogLevel_Business

	// Detalle de todos los request y response de Negocio que se intercambien con el exterior (no incluye mensajes internos de la app, como comandos)
	LogLevel_Message

	// Información detallada de la ejecución
	LogLevel_Debug

	// Detalle de querys ejecutadas con sus parámetros de entrada
	LogLevel_Query

	// Máximo nivel de detalle, como los valores devueltos por las querys, trace de http, etc.
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
	}
}

func (me LogLevel) IsLogAllowed(other LogLevel) bool {
	return (me >= other)
}
