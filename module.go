package logger

import (
	"log/slog"

	"go.uber.org/fx"
)

func NewLoggerModule(moduleName string) fx.Option {
	return fx.Module(moduleName, fx.Provide(GetLogger))
}

func NewLogger() fx.Option {
	return fx.Provide(GetLogger)
}

func NewSubLoggerModule(moduleName, group string) fx.Option {
	return fx.Module(moduleName, fx.Provide(
		func(logger *slog.Logger) *slog.Logger {
			return GetSubLogger(logger, group)
		},
	))
}

func NewSubLogger(group string) fx.Option {
	return fx.Decorate(
		func(logger *slog.Logger) *slog.Logger {
			return GetSubLogger(logger, group)
		},
	)
}
