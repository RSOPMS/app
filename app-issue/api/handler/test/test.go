package test

import (
	"log"
	"net/http"
	"time"
)

type FaultTestHandler struct{}

func NewFaultTestHandler() *FaultTestHandler {
	return &FaultTestHandler{}
}

func (h *FaultTestHandler) GetFaultTest(w http.ResponseWriter, r *http.Request) error {
	log.Println("GetFaultTest handler called")

	// Simulate a delay of 5 seconds
	time.Sleep(5 * time.Second)

	// Respond with HTTP 200 OK
	w.WriteHeader(http.StatusOK)
	return nil
}
