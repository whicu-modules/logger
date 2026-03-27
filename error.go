package logger

import (
	"errors"
	"fmt"
)

var (
	ErrInvalidLogLevel = errors.New("invalid log level")
)

func wrapLoggerError(err error) error {
	return fmt.Errorf("logger: %w", err)
}
