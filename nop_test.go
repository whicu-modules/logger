package logger

import (
	"context"
	"log/slog"
	"testing"
	"time"
)

func TestNOPHandler_NoOpBehavior(t *testing.T) {
	h := NewNOPHandler()
	ctx := context.Background()

	if h.Enabled(ctx, slog.LevelError) {
		t.Fatal("expected NOP handler to disable all levels")
	}

	rec := slog.NewRecord(time.Now(), slog.LevelInfo, "test", 0)
	if err := h.Handle(ctx, rec); err != nil {
		t.Fatalf("expected nil error from Handle, got %v", err)
	}

	if _, ok := h.WithAttrs(nil).(NOPHandler); !ok {
		t.Fatal("expected WithAttrs to return NOPHandler")
	}
	if _, ok := h.WithGroup("group").(NOPHandler); !ok {
		t.Fatal("expected WithGroup to return NOPHandler")
	}
}

func TestNewNOPSlog_DisablesAllLevels(t *testing.T) {
	log := NewNOPSlog()
	if log == nil {
		t.Fatal("expected non-nil logger")
	}

	if log.Enabled(context.Background(), slog.LevelDebug) {
		t.Fatal("expected debug level to be disabled")
	}
}
