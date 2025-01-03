package profile

import (
	"app-issue/pkg"
	"app-issue/template"
	"database/sql"
	"framework/api"
	"net/http"
)

type ProfileHandler struct {
	Db *sql.DB
}

func NewProfileHandler(db *sql.DB) *ProfileHandler {
	return &ProfileHandler{
		Db: db,
	}
}

func (h *ProfileHandler) GetProfilePage(w http.ResponseWriter, r *http.Request) error {
	email, ok := r.Context().Value(api.ContextSubjectKey).(string)
	if !ok {
		return nil
	}

	user, err := pkg.ReadUserProfile(h.Db, email)
	if err != nil {
		return err
	}

	return template.RenderLayout(w, "profilePage", user)
}
