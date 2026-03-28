# logger

Logging module for Go projects built on top of `log/slog`.

Features:

- colored console logging via `tint`
- optional JSON file logging via `lumberjack` rotation
- handler fanout (console + file at the same time)
- integration helpers for `uber/fx`
- no-op logger for tests and disabled logging scenarios

## Installation

```bash
go get github.com/whicu-modules/logger
```

## Configuration

Use `config.Config` from `github.com/whicu-modules/logger/config`.

```go
type Config struct {
	Level     string       // debug | info | warn | error
	AddSource bool
	Handler   slog.Handler // optional custom console handler
	config.LumberjackConfig
}

type LumberjackConfig struct {
	Level    string // debug | info | warn | error
	Path     string // if empty: file logging is disabled
	Size     int    // max size in MB before rotation
	Compress bool
}
```

## Quick Start

```go
package main

import (
	"log/slog"

	"github.com/whicu-modules/logger"
	"github.com/whicu-modules/logger/config"
)

func main() {
	cfg := config.Config{
		Level:     "info",
		AddSource: false,
		LumberjackConfig: config.LumberjackConfig{
			Path:     "./app.log",
			Size:     128,
			Level:    "info",
			Compress: true,
		},
	}

	log, closer, err := logger.GetLogger(cfg)
	if err != nil {
		panic(err)
	}
	defer closer.Close()

	log.Info("service started", slog.String("service", "api"))
}
```

## Public API

### `GetLogger(cfg config.Config) (*slog.Logger, io.Closer, error)`

- creates logger with console handler
- if `cfg.Path` is not empty, also adds JSON file handler with rotation
- returns `io.Closer` (close it on shutdown)

### `GetSubLogger(logger *slog.Logger, group string) *slog.Logger`

Returns grouped logger (`logger.WithGroup(group)`).

### `InitLogger(level slog.Level, addSource bool) slog.Handler`

Creates default console handler (tint).

### NOP logger

- `NewNOPHandler() NOPHandler`
- `NewNOPSlog() *slog.Logger`

Useful for tests or when logging should be fully disabled.

### Fx helpers

- `NewLoggerModule(moduleName string) fx.Option`
- `NewLogger() fx.Option`
- `NewSubLoggerModule(moduleName, group string) fx.Option`
- `NewSubLogger(group string) fx.Option`

These helpers register `GetLogger` / `GetSubLogger` providers in Fx.

## Error behavior

- invalid level returns `ErrInvalidLogLevel`
- all internal errors are wrapped with `logger:` prefix

Supported level values:

- `debug`
- `info`
- `warn`
- `error`