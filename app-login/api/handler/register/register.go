package register

import (
	"app-login/pkg"
	"app-login/template"
	"database/sql"
	"net/http"
	"os"
	"strconv"
)

type RegisterHandler struct {
	Db *sql.DB
}

func NewRegisterHandler(db *sql.DB) *RegisterHandler {
	return &RegisterHandler{
		Db: db,
	}
}

func (h *RegisterHandler) GetRegisterPage(w http.ResponseWriter, r *http.Request) error {
	return template.RenderLayout(w, "registerPage", nil)
}

func (h *RegisterHandler) ProcessRegister(w http.ResponseWriter, r *http.Request) error {
	err := r.ParseForm()
	if err != nil {
		return err
	}

	name := r.FormValue("name")
	surname := r.FormValue("surname")
	email := r.FormValue("email")
	password := r.FormValue("password")
	rep_password := r.FormValue("rep_password")

	if password != rep_password {
		return nil
	}

	roleId, err := strconv.Atoi(r.FormValue("roleId"))
	if err != nil {
		return err
	}

	err = pkg.ProcessRegister(h.Db, name, surname, roleId, email, password)
	if err != nil {
		return err
	}

	w.Header().Set("HX-Redirect", os.Getenv("URL_PREFIX_LOGIN")+"/")
	w.WriteHeader(http.StatusOK)
	return nil
}

// Get roles from the database for the Create New Issue form
func (h *RegisterHandler) GetRolesForm(w http.ResponseWriter, r *http.Request) error {
	roles, err := pkg.ReadRoles(h.Db)
	if err != nil {
		return err
	}

	return template.RenderLayout(w, "rolesForm", roles)
}
