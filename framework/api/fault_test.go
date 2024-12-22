package api

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"
	"time"
)

func TestCircuitBreaker(t *testing.T) {
	cb := NewCircuitBreaker(3, 2*time.Second)

	request := func() error {
		return errors.New("request failed")
	}

	for i := 0; i < 3; i++ {
		if err := cb.Execute(request); err == nil {
			t.Fatal("expected error, got nil")
		}
	}

	if err := cb.Execute(request); err == nil {
		t.Fatal("expected circuit breaker to be open")
	}

	time.Sleep(3 * time.Second)

	if err := cb.Execute(func() error { return nil }); err != nil {
		t.Fatal("expected circuit breaker to be half-open")
	}
}

func TestRetry(t *testing.T) {
	attempts := 0
	request := func() error {
		attempts++
		if attempts < 3 {
			return errors.New("request failed")
		}
		return nil
	}

	err := Retry(request, 5, 100*time.Millisecond)
	if err != nil {
		t.Fatal("expected request to succeed after retries")
	}

	if attempts != 3 {
		t.Fatalf("expected 3 attempts, got %d", attempts)
	}
}

func TestBulkhead(t *testing.T) {
	bulkhead := NewBulkhead(2)
	var wg sync.WaitGroup
	wg.Add(3)

	executeRequest := func() {
		defer wg.Done()
		err := bulkhead.Execute(func() error {
			time.Sleep(100 * time.Millisecond)
			return nil
		})
		if err != nil {
			t.Log(err)
		}
	}

	go executeRequest()
	go executeRequest()
	go executeRequest()

	wg.Wait()
}

func TestFaultTolerantMiddleware(t *testing.T) {
	cb := NewCircuitBreaker(3, 2*time.Second)
	bulkhead := NewBulkhead(2)
	retryAttempts := 3
	retryDelay := 100 * time.Millisecond

	handler := FaultTolerantMiddleware(cb, retryAttempts, retryDelay, bulkhead, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	req, _ := http.NewRequest("GET", "/", nil)
	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}
