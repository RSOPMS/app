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

func (h *ProjectHandler) GetProjectsPage(w http.ResponseWriter, r *http.Request) error {
	return template.RenderLayout(w, "projects", nil)
}

func (h *ProjectHandler) GetProjectPage(w http.ResponseWriter, r *http.Request) error {
	project_id := r.PathValue("project_id")
	project, err := pkg.ReadProject(h.Db, project_id)
	if err != nil {
		return err
	}

	return template.RenderLayout(w, "project", project)
}

func (h *ProjectHandler) GetProjectTable(w http.ResponseWriter, r *http.Request) error {
	projects, err := pkg.ReadProjects(h.Db)
	if err != nil {
		return err
	}

	return template.RenderProject(w, "table", projects)
}

func (h *ProjectHandler) GetIssueTable(w http.ResponseWriter, r *http.Request) error {
	project_id := r.PathValue("project_id")
	issues, err := pkg.ReadProjectIssues(h.Db, project_id)
	if err != nil {
		return err
	}

	return template.RenderIssue(w, "table", issues)
}
