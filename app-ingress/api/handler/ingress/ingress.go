package ingress

import (
	"app-ingress/pkg"
	"database/sql"
	"encoding/json"
	"net/http"
)

type IngressHandler struct {
	Db *sql.DB
}

func NewIngressHandler(db *sql.DB) *IngressHandler {
	return &IngressHandler{
		Db: db,
	}
}

// TODO THIS IS UNNEEDED
func (h *IngressHandler) PostIngress(w http.ResponseWriter, r *http.Request) error {
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
