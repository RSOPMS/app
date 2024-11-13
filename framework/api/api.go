// Package api provides helper functions for error handling and response
// generation.
package api

import (
	"log"
	"net/http"
)

// Handler is a handler type that can return an error.
type Handler func(http.ResponseWriter, *http.Request) error

// CreateHandler handles errors and returns a HandlerFunc.
func CreateHandler(handler Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := handler(w, r); err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
		}
	}
}
