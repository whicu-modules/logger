package logger

import (
	"log/slog"
	"os"
	"time"

	"github.com/lmittmann/tint"
)

func InitLogger(level slog.Level, addSource bool) slog.Handler {
	return tint.NewHandler(os.Stdout, &tint.Options{
		AddSource:  addSource,
		Level:      level,
		TimeFormat: time.ANSIC,
	})
}
