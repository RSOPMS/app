package bulk

import (
	"app-bulk/pkg"
	"encoding/json"
	"net/http"
)

type BulkHandler struct {
}

func NewBulkHandler() *BulkHandler {
	return &BulkHandler{}
}

func (h *BulkHandler) PostBulk(w http.ResponseWriter, r *http.Request) error {
	var payload pkg.InputPayload
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		return err
	}

	err = pkg.AddPayloadToDB(payload)
	if err != nil {
		return err
	}

	return nil
}
