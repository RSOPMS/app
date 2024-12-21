package login

import (
	"app-login/pkg"
	"app-login/template"
	"database/sql"
	"log"
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
	w.WriteHeader(http.StatusOK)

	return template.RenderLoginLayout(w, "loginPage", nil)
}

func (h *LoginHandler) ProcessLogin(w http.ResponseWriter, r *http.Request) error {
	username := r.FormValue("e-mail")
	password := r.FormValue("password")

	log.Println("E-mail:", username)
	log.Println("Password:", password)

	cookie, err := pkg.ProcessLogin(h.Db, username, password)
	if err != nil {
		data := map[string]interface{}{
			"Error": "Invalid login credentials. Please try again.",
		}
		return template.RenderLoginLayout(w, "loginPage", data)
	}

	http.SetCookie(w, cookie)
	http.Redirect(w, r, os.Getenv("URL_PREFIX_ISSUE")+"/projects/", http.StatusSeeOther)
	return nil
}
