package api

import (
	"context"
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

// AuthMiddleware validates the JWT and sets the request context.
func AuthMiddleware(jwtHandler JwtHandler, onUnautorized http.HandlerFunc, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		onErr := func(err error) {
			log.Println(err)
			onUnautorized.ServeHTTP(w, r)
		}

		cookie, err := r.Cookie(jwtHandler.cookieName)
		if err != nil {
			onErr(err)
			return
		}

		err = cookie.Valid()
		if err != nil {
			onErr(err)
			return
		}

		token, err := jwtHandler.ParseJwt(cookie.Value)
		if err != nil {
			onErr(err)
			return
		}

		subject, err := token.Claims.GetSubject()
		if err != nil {
			onErr(err)
			return
		}

		req := r.WithContext(context.WithValue(r.Context(), "subject", subject))
		*r = *req

		next.ServeHTTP(w, r)
	})
}
