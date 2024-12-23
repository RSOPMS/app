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

// AuthHandler handles authentication.
type AuthHandler struct {
	jwtHandler    JwtHandler
	onUnautorized http.HandlerFunc
}

// NewAuthHandler creates a new AuthMiddleware that validates the JWT and
// sets the erquest context.
func NewAuthHandler(jwtHandler JwtHandler, onUnautorized http.HandlerFunc) *AuthHandler {
	return &AuthHandler{
		jwtHandler:    jwtHandler,
		onUnautorized: onUnautorized,
	}
}

// AuthMiddleware validates the JWT and sets the request context.
func (h *AuthHandler) AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		onErr := func(err error) {
			log.Println(err)
			h.onUnautorized.ServeHTTP(w, r)
		}

		cookie, err := r.Cookie(h.jwtHandler.cookieName)
		if err != nil {
			onErr(err)
			return
		}

		err = cookie.Valid()
		if err != nil {
			onErr(err)
			return
		}

		token, err := h.jwtHandler.ParseJwt(cookie.Value)
		if err != nil {
			onErr(err)
			return
		}

		subject, err := token.Claims.GetSubject()
		if err != nil {
			onErr(err)
			return
		}

		ctx := context.WithValue(r.Context(), ContextSubjectKey, subject)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
