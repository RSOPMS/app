package project

import (
	"app-issue/pkg"
	"app-issue/template"
	"database/sql"
	"net/http"
)

type ProjectHandler struct {
	Db *sql.DB
}

func NewProjectHandler(db *sql.DB) *ProjectHandler {
	return &ProjectHandler{
		Db: db,
	}
}

func (h *ProjectHandler) GetProject(w http.ResponseWriter, r *http.Request) error {
	return template.RenderLayout(w, "project", nil)
}

func (h *ProjectHandler) GetProjectList(w http.ResponseWriter, r *http.Request) error {
	projects, err := pkg.ReadProjects(h.Db)
	if err != nil {
		return err
	}

	return template.RenderProject(w, "list", projects)
}
