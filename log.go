package logger

import (
	"io"
	"log/slog"

	slogmulti "github.com/samber/slog-multi"
	"github.com/whicu-modules/logger/config"
	"gopkg.in/natefinch/lumberjack.v2"
)

type nopCloser struct{}

func (nopCloser) Close() error { return nil }

func GetLogger(cfg config.Config) (log *slog.Logger, closer io.Closer, err error) {
	defer func() {
		if err != nil {
			err = wrapLoggerError(err)
		}
	}()

	handlers := make([]slog.Handler, 0, 2)
	closer = io.Closer(nopCloser{})

	if cfg.Path != "" {
		fileHandler, fileCloser, err := getLumberjackHandler(cfg.LumberjackConfig)
		if err != nil {
			return nil, nil, err
		}
		handlers = append(handlers, fileHandler)
		closer = fileCloser
	}

	level, err := getLevel(cfg.Level)
	if err != nil {
		return nil, nil, err
	}

	handlers = append(handlers, getHandler(cfg.Handler, level, cfg.AddSource))

	return slog.New(slogmulti.Fanout(handlers...)), closer, nil

}

func GetSubLogger(logger *slog.Logger, group string) *slog.Logger {
	return logger.WithGroup(group)
}

func getLumberjackHandler(cfg config.LumberjackConfig) (slog.Handler, io.Closer, error) {
	logFile := &lumberjack.Logger{
		Filename:  cfg.Path,
		MaxSize:   cfg.Size,
		LocalTime: true,
		Compress:  cfg.Compress,
	}

	level, err := getLevel(cfg.Level)
	if err != nil {
		return nil, nil, err
	}

	return slog.NewJSONHandler(logFile, &slog.HandlerOptions{
		Level: level,
	}), logFile, nil
}

func getHandler(handler slog.Handler, level slog.Level, addSource bool) slog.Handler {
	if handler != nil {
		return handler
	}

	return InitLogger(level, addSource)
}

func getLevel(level string) (slog.Level, error) {
	levels := map[string]slog.Level{
		"debug": slog.LevelDebug,
		"info":  slog.LevelInfo,
		"warn":  slog.LevelWarn,
		"error": slog.LevelError,
	}

	if parsed, ok := levels[level]; ok {
		return parsed, nil
	}

	return slog.LevelInfo, ErrInvalidLogLevel
}
