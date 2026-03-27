package logger

import (
	"context"
	"errors"
	"log/slog"
	"os"
	"path/filepath"
	"testing"

	"github.com/whicu-modules/logger/config"
)

func Test_getLevel(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    slog.Level
		wantErr bool
	}{
		{name: "debug", input: "debug", want: slog.LevelDebug},
		{name: "info", input: "info", want: slog.LevelInfo},
		{name: "warn", input: "warn", want: slog.LevelWarn},
		{name: "error", input: "error", want: slog.LevelError},
		{name: "invalid", input: "trace", want: slog.LevelInfo, wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getLevel(tt.input)
			if got != tt.want {
				t.Fatalf("unexpected level: got %v, want %v", got, tt.want)
			}

			if tt.wantErr {
				if !errors.Is(err, ErrInvalidLogLevel) {
					t.Fatalf("expected ErrInvalidLogLevel, got %v", err)
				}
				return
			}

			if err != nil {
				t.Fatalf("expected nil error, got %v", err)
			}
		})
	}
}

func TestGetLogger_InvalidMainLevel(t *testing.T) {
	log, closer, err := GetLogger(config.Config{Level: "trace"})
	if log != nil {
		t.Fatal("expected nil logger")
	}
	if closer != nil {
		t.Fatal("expected nil closer")
	}
	if !errors.Is(err, ErrInvalidLogLevel) {
		t.Fatalf("expected ErrInvalidLogLevel, got %v", err)
	}
}

func TestGetLogger_InvalidFileLevel(t *testing.T) {
	log, closer, err := GetLogger(config.Config{
		Level: "info",
		LumberjackConfig: config.LumberjackConfig{
			Path:  filepath.Join(t.TempDir(), "test.log"),
			Level: "trace",
		},
	})
	if log != nil {
		t.Fatal("expected nil logger")
	}
	if closer != nil {
		t.Fatal("expected nil closer")
	}
	if !errors.Is(err, ErrInvalidLogLevel) {
		t.Fatalf("expected ErrInvalidLogLevel, got %v", err)
	}
}

func TestGetLogger_SuccessWithoutFile(t *testing.T) {
	log, closer, err := GetLogger(config.Config{Level: "debug"})
	t.Cleanup(func() {
		closer.Close()
	})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if log == nil {
		t.Fatal("expected non-nil logger")
	}
	if !log.Enabled(context.Background(), slog.LevelDebug) {
		t.Fatal("expected logger to enable debug level")
	}
}

func TestGetLogger_SuccessWithFile(t *testing.T) {
	path := filepath.Join(t.TempDir(), "app.log")

	log, closer, err := GetLogger(config.Config{
		Level:   "debug",
		Handler: NewNOPHandler(),
		LumberjackConfig: config.LumberjackConfig{
			Path:  path,
			Size:  1,
			Level: "info",
		},
	})
	t.Cleanup(func() {
		closer.Close()
	})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if log == nil {
		t.Fatal("expected non-nil logger")
	}

	log.Info("write to file")

	info, statErr := os.Stat(path)
	if statErr != nil {
		t.Fatalf("expected log file to be created: %v", statErr)
	}
	if info.Size() == 0 {
		t.Fatal("expected log file to be non-empty")
	}
}
