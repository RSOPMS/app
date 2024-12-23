package api

import (
	"log"
	"math/rand"
	"net/http"
	"time"
)

// Middleware is a http.Handler function wrapper.
type Middleware func(http.Handler) http.Handler

// CreateMiddlewareStack creates a middleware stack. It is used to reduce
// middleware nesting.
func CreateMiddlewareStack(middleware ...Middleware) Middleware {
	return func(next http.Handler) http.Handler {
		for i := len(middleware) - 1; i >= 0; i-- {
			m := middleware[i]
			next = m(next)
		}
		return next
	}
}

// LoggingMiddleware logs the request method, path and response time.
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Println(r.Method, r.URL.Path, r.RemoteAddr, time.Since(start))
	})
}

// Retry wraps the Handler type and provides retry functionality.
func (h *RetryHandler) Retry(next Handler) Handler {
	return func(w http.ResponseWriter, r *http.Request) error {
		var err error
		delay := h.delay

		for i := 0; i < h.attempts; i++ {
			if err = next(w, r); err == nil {
				return nil
			}

			time.Sleep(delay)
			delay *= 2
			delay += time.Duration(rand.Float32()) * h.maxJitter
		}

		return err
	}
}
