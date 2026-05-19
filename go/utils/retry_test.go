package utils

import (
	"context"
	"errors"
	"testing"
	"time"
)

func TestWithRetry_Success(t *testing.T) {
	opts := DefaultRetryOptions()
	opts.MaxAttempts = 3
	opts.InitialDelay = 1 * time.Millisecond

	attempts := 0
	err := WithRetry(context.Background(), opts, func() error {
		attempts++
		if attempts < 2 {
			return errors.New("fail")
		}
		return nil
	})

	if err != nil {
		t.Errorf("Expected success, got error: %v", err)
	}
	if attempts != 2 {
		t.Errorf("Expected 2 attempts, got %d", attempts)
	}
}

func TestWithRetry_Failure(t *testing.T) {
	opts := DefaultRetryOptions()
	opts.MaxAttempts = 2
	opts.InitialDelay = 1 * time.Millisecond

	attempts := 0
	err := WithRetry(context.Background(), opts, func() error {
		attempts++
		return errors.New("permanent fail")
	})

	if err == nil {
		t.Error("Expected error, got nil")
	}
	if attempts != 2 {
		t.Errorf("Expected 2 attempts, got %d", attempts)
	}
}
