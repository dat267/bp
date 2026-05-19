package logger

import (
	"log/slog"
	"testing"
)

func TestNewLogger(t *testing.T) {
	logger := NewLogger("debug")
	if logger == nil {
		t.Fatal("Expected logger to be initialized, got nil")
	}

	// Verify level is correctly handled (manual check of output usually, but here we just check init)
	if logger.Handler().Enabled(nil, slog.LevelDebug) == false {
		t.Error("Expected debug level to be enabled")
	}
}
