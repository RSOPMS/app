package api

import (
	"errors"
	"sync"
	"time"
)

// -------- Retry --------

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

// ------ Timeout -------

// TimeoutHandler handles timeout for requests.
type TimeoutHandler struct {
	timeout time.Duration
}

// TimeoutHandlerOptionFunc sets a TimeoutHandler field.
type TimeoutHandlerOptionFunc func(t *TimeoutHandler)

// WithTimeout sets the timeout duration.
func WithTimeout(timeout time.Duration) TimeoutHandlerOptionFunc {
	return func(t *TimeoutHandler) {
		t.timeout = timeout
	}
}

// NewTimeoutHandler creates a new timeout handler.
func NewTimeoutHandler(options ...TimeoutHandlerOptionFunc) *TimeoutHandler {
	t := &TimeoutHandler{
		timeout: time.Second, // default timeout
	}

	for _, optionFunc := range options {
		optionFunc(t)
	}

	return t
}

// The ones below are not used in the app, but are provided for reference.
// ------ Circuit Breaker -------

type CircuitBreakerState int

const (
	Closed CircuitBreakerState = iota
	Open
	HalfOpen
)

type CircuitBreaker struct {
	state            CircuitBreakerState
	failureCount     int
	failureThreshold int
	recoveryTimeout  time.Duration
	mutex            sync.Mutex
	lastFailureTime  time.Time
}

func NewCircuitBreaker(failureThreshold int, recoveryTimeout time.Duration) *CircuitBreaker {
	return &CircuitBreaker{
		state:            Closed,
		failureThreshold: failureThreshold,
		recoveryTimeout:  recoveryTimeout,
	}
}

func (cb *CircuitBreaker) Execute(request func() error) error {
	cb.mutex.Lock()
	defer cb.mutex.Unlock()

	switch cb.state {
	case Open:
		if time.Since(cb.lastFailureTime) > cb.recoveryTimeout {
			cb.state = HalfOpen
		} else {
			return errors.New("circuit breaker is open")
		}
	case HalfOpen:
		if err := request(); err != nil {
			cb.state = Open
			cb.lastFailureTime = time.Now()
			return err
		}
		cb.state = Closed
		return nil
	}

	if err := request(); err != nil {
		cb.failureCount++
		if cb.failureCount >= cb.failureThreshold {
			cb.state = Open
			cb.lastFailureTime = time.Now()
		}
		return err
	}

	cb.failureCount = 0
	return nil
}

// ------ Bulkhead -------

type Bulkhead struct {
	maxConcurrent int
	current       int
	mutex         sync.Mutex
}

func NewBulkhead(maxConcurrent int) *Bulkhead {
	return &Bulkhead{
		maxConcurrent: maxConcurrent,
	}
}

func (b *Bulkhead) Execute(request func() error) error {
	b.mutex.Lock()
	if b.current >= b.maxConcurrent {
		b.mutex.Unlock()
		return errors.New("bulkhead limit reached")
	}
	b.current++
	b.mutex.Unlock()

	defer func() {
		b.mutex.Lock()
		b.current--
		b.mutex.Unlock()
	}()

	return request()
}
