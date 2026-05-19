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

func TestWithRetry_ContextCancel(t *testing.T) {
	opts := DefaultRetryOptions()
	opts.MaxAttempts = 5
	opts.InitialDelay = 1 * time.Hour // Long delay

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()

	err := WithRetry(ctx, opts, func() error {
		return errors.New("fail")
	})

	if err != context.DeadlineExceeded {
		t.Errorf("Expected DeadlineExceeded, got %v", err)
	}
}
