package login

import (
	"app-login/template"
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

func (h *LoginHandler) GetLoginPage(w http.ResponseWriter, r *http.Request) error {
	w.WriteHeader(http.StatusOK)

	return template.RenderLoginLayout(w, "loginPage", "test")
}

func (h *LoginHandler) PostUserLogin(w http.ResponseWriter, r *http.Request) error {
	w.WriteHeader(http.StatusOK)

	return nil
}
