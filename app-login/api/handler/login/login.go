package login

import (
	"app-login/pkg"
	"app-login/template"
	"database/sql"
	"net/http"
	"os"
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
	return template.RenderLayout(w, "loginPage", nil)
}

func (h *LoginHandler) ProcessLogin(w http.ResponseWriter, r *http.Request) error {
	username := r.FormValue("e-mail")
	password := r.FormValue("password")

	cookie, err := pkg.ProcessLogin(h.Db, username, password)
	if err != nil {
		return err
	}

	http.SetCookie(w, cookie)
	w.Header().Set("HX-Redirect", os.Getenv("URL_PREFIX_ISSUE")+"/projects/")
	w.WriteHeader(http.StatusOK)
	return nil
}

func (h *LoginHandler) ProcessLogout(w http.ResponseWriter, r *http.Request) error {
	cookie := pkg.ProcessLogout()
	http.SetCookie(w, cookie)
	http.Redirect(w, r, os.Getenv("URL_PREFIX_LOGIN")+"/", http.StatusSeeOther)
	return nil
}
