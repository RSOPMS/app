package login

import (
	"database/sql"
	"net/http"
)

type LoginHandler struct {
	Db *sql.DB
}

func NewLoginHandler(db *sql.DB) *LoginHandler {
	return &LoginHandler{
		Db: db,
	}
}

func (h *LoginHandler) GetLoginLive(w http.ResponseWriter, r *http.Request) error {
	w.WriteHeader(http.StatusOK)

	return nil
}
