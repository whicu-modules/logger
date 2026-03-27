package logger

import (
	"context"
	"log/slog"
	"testing"
)

func TestInitLogger_RespectsLevel(t *testing.T) {
	h := InitLogger(slog.LevelWarn, false)
	if h == nil {
		t.Fatal("expected non-nil handler")
	}

	ctx := context.Background()
	if h.Enabled(ctx, slog.LevelInfo) {
		t.Fatal("expected info level to be disabled")
	}
	if !h.Enabled(ctx, slog.LevelWarn) {
		t.Fatal("expected warn level to be enabled")
	}
}
