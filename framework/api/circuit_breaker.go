package api

import (
	"errors"
	"sync"
	"time"
)

type CircuitBreakerState int

const (
	Closed CircuitBreakerState = iota
	Open
	HalfOpen
)

type CircuitBreaker struct {
	state            CircuitBreakerState
	failureCount     int
	successCount     int
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
