package register

import (
	"app-login/pkg"
	"app-login/template"
	"database/sql"
	"errors"
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

func (h *RegisterHandler) PostRegisterNew(w http.ResponseWriter, r *http.Request) error {
	err := r.ParseForm()
	if err != nil {
		return err
	}

	name := r.FormValue("name")
	surname := r.FormValue("surname")
	email := r.FormValue("email")
	password := r.FormValue("password")
	passwordRepeat := r.FormValue("passwordRepeat")

	if password != passwordRepeat {
		return errors.New("passwords do not match")
	}

	roleId, err := strconv.Atoi(r.FormValue("roleId"))
	if err != nil {
		return err
	}

	err = pkg.ProcessRegister(h.Db, name, surname, roleId, email, password)
	if err != nil {
		return err
	}

	http.Redirect(w, r, os.Getenv("URL_PREFIX_LOGIN")+"/", http.StatusMovedPermanently)
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
