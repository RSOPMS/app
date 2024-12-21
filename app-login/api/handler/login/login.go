package login

import (
	"app-login/template"
	"database/sql"
	"log"
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

	return template.RenderLoginLayout(w, "loginPage", nil)
}

func (h *LoginHandler) ProcessLogin(w http.ResponseWriter, r *http.Request) error {
	w.WriteHeader(http.StatusOK)

	username := r.FormValue("username")
	password := r.FormValue("password")

	log.Println("Username:", username)
	log.Println("Password:", password)
	return nil
}
