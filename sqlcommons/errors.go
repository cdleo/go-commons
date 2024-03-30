package sqlcommons

import (
	"errors"
)

// Errors
var (
	DBNotSupported   = errors.New("DataBase not supported")
	DBNotInitialized = errors.New("DataBase not initialized")
	ConnectionFailed = errors.New("Connection Failed")
	ConnectionClosed = errors.New("Connection Closed")
	NextValueFailed  = errors.New("Unable to get next sequence value")
	OpNotSupported   = errors.New("Operation not supported")
	TxNotFoundInCtx  = errors.New("SQLTx not found in context")

	UniqueConstraintViolation     = errors.New("Unique constraint violation")
	IntegrityConstraintViolation  = errors.New("Integrity constraint violation")
	ValueTooLargeForColumn        = errors.New("Value too large for column")
	ValueLargerThanPrecision      = errors.New("Value larger than specified precision")
	CannotSetNullColumn           = errors.New("Cannot set a NULL value into a NOT NULL column")
	InvalidNumericValue           = errors.New("Invalid number")
	SubqueryReturnsMoreThanOneRow = errors.New("Subquery returns more than one row")
)
