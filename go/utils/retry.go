package utils

import (
	"context"
	"math/rand"
	"time"
)

type RetryOptions struct {
	MaxAttempts    int
	InitialDelay   time.Duration
	MaxDelay       time.Duration
	BackoffFactor  float64
	UseJitter      bool
}

func DefaultRetryOptions() RetryOptions {
	return RetryOptions{
		MaxAttempts:   3,
		InitialDelay:  1 * time.Second,
		MaxDelay:      30 * time.Second,
		BackoffFactor: 2.0,
		UseJitter:     true,
	}
}

func WithRetry(ctx context.Context, opts RetryOptions, operation func() error) error {
	var lastErr error
	delay := opts.InitialDelay

	for attempt := 1; attempt <= opts.MaxAttempts; attempt++ {
		if err := operation(); err == nil {
			return nil
		} else {
			lastErr = err
		}

		if attempt == opts.MaxAttempts {
			break
		}

		// Calculate sleep time with backoff
		sleepTime := delay
		if opts.UseJitter {
			// Add random jitter: [0, delay/2]
			jitter := time.Duration(rand.Float64() * float64(delay) / 2)
			sleepTime += jitter
		}

		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-time.After(sleepTime):
		}

		// Update delay for next iteration
		delay = time.Duration(float64(delay) * opts.BackoffFactor)
		if delay > opts.MaxDelay {
			delay = opts.MaxDelay
		}
	}

	return lastErr
}
