# logger

Simple logging module for Go projects.

It uses slog for console logs and can also write JSON logs to a file with rotation.

## What it does

- Creates a ready-to-use logger with configurable level.
- Supports console output with optional source info.
- Supports file output via lumberjack with size-based rotation.

## Quick example

```go
cfg := logger.Config{
	Level:     "info",
	AddSource: false,
	LumberjackConfig: logger.LumberjackConfig{
		Path:  "./app.log",
		Size:  128,
		Level: "info",
	},
}

log, err := logger.GetLogger(cfg)
if err != nil {
	panic(err)
}

log.Info("service started")
```