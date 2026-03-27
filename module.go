package logger

import (
	"log/slog"

	"go.uber.org/fx"
)

func NewModule(moduleName string) fx.Option {
	return fx.Module(moduleName, fx.Provide(GetLogger))
}

func NewSubLoggerModule(group string) fx.Option {
	return fx.Provide(
		func(logger *slog.Logger) *slog.Logger {
			return GetSubLogger(logger, group)
		},
	)
}
