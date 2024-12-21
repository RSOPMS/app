package api

import (
	"log"
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

func FaultTolerantMiddleware(cb *CircuitBreaker, retryAttempts int, retryDelay time.Duration, bulkhead *Bulkhead, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := cb.Execute(func() error {
			return Retry(func() error {
				return bulkhead.Execute(func() error {
					next.ServeHTTP(w, r)
					return nil
				})
			}, retryAttempts, retryDelay)
		})

		if err != nil {
			// TODO: More sophisticated error handling
			log.Println(err)
		}
	})
}
