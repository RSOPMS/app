package bulk

import (
	"app-bulk/pkg"
	"database/sql"
	"encoding/json"
	"net/http"
)

type BulkHandler struct {
	Db *sql.DB
}

func NewBulkHandler(db *sql.DB) *BulkHandler {
	return &BulkHandler{
		Db: db,
	}
}

func (h *BulkHandler) PostBulk(w http.ResponseWriter, r *http.Request) error {
	var payload pkg.InputPayload
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		return err
	}

	err = pkg.AddPayloadToDB(h.Db, payload)
	if err != nil {
		return err
	}

	return nil
}
