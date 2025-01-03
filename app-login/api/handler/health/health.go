package health

import (
	"database/sql"
	"net/http"
)

type HealthHandler struct {
	Db *sql.DB
}

func NewHealthHandler(db *sql.DB) *HealthHandler {
	return &HealthHandler{
		Db: db,
	}
}

func (h *HealthHandler) GetHealthLive(w http.ResponseWriter, r *http.Request) error {
	if err := h.Db.Ping(); err != nil {
		return err
	}

	w.WriteHeader(http.StatusOK)

	return nil
}

func (h *HealthHandler) GetHealthReady(w http.ResponseWriter, r *http.Request) error {
	if err := h.Db.Ping(); err != nil {
		return err
	}

	w.WriteHeader(http.StatusOK)

	return nil
}
