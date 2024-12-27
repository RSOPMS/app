package fault

import (
	"fmt"
	"net/http"
	"time"
)

type FaultTestHandler struct{}

func NewFaultTestHandler() *FaultTestHandler {
	return &FaultTestHandler{}
}

func (h *FaultTestHandler) GetTimeoutBad(w http.ResponseWriter, r *http.Request) error {
	time.Sleep(5 * time.Second)
	return nil
}

func (h *FaultTestHandler) GetRetryBad(w http.ResponseWriter, r *http.Request) error {
	return fmt.Errorf("Artificial error")
}
