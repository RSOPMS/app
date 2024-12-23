package api

import (
	"time"
)

// RetryHandler handles retry attempts on errors.
type RetryHandler struct {
	attempts  int
	delay     time.Duration
	maxJitter time.Duration
}

// RetryHandlerOptionFunc sets a RetryHandler field.
type RetryHandlerOptionFunc func(r *RetryHandler)

// WithAttempts sets retry delay.
func WithAttempts(attempts int) RetryHandlerOptionFunc {
	return func(r *RetryHandler) {
		r.attempts = attempts
	}
}

// WithDelay sets retry delay.
func WithDelay(delay time.Duration) RetryHandlerOptionFunc {
	return func(r *RetryHandler) {
		r.delay = delay
	}
}

// WithMaxJitter sets maximum retry jitter.
func WithMaxJitter(maxJitter time.Duration) RetryHandlerOptionFunc {
	return func(r *RetryHandler) {
		r.maxJitter = maxJitter
	}
}

// NewRetryHandler creates a new retry handler.
func NewRetryHandler(options ...RetryHandlerOptionFunc) *RetryHandler {
	r := &RetryHandler{
		attempts:  5,
		delay:     time.Second,
		maxJitter: 100 * time.Millisecond,
	}

	for _, optionFunc := range options {
		optionFunc(r)
	}

	return r
}
