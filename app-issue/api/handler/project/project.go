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
	return template.RenderLayout(w, "projectsPage", nil)
}

func (h *ProjectHandler) GetProjectPage(w http.ResponseWriter, r *http.Request) error {
	projectId := r.PathValue("projectId")

	project, err := pkg.ReadProject(h.Db, projectId)
	if err != nil {
		return err
	}

	return template.RenderLayout(w, "projectPage", project)
}

func (h *ProjectHandler) GetProjectsTable(w http.ResponseWriter, r *http.Request) error {
	projects, err := pkg.ReadProjects(h.Db)
	if err != nil {
		return err
	}

	return template.RenderProject(w, "projectsTable", projects)
}

func (h *ProjectHandler) GetIssuesTable(w http.ResponseWriter, r *http.Request) error {
	projectId := r.PathValue("projectId")

	issues, err := pkg.ReadIssues(h.Db, projectId)
	if err != nil {
		return err
	}

	return template.RenderIssue(w, "issuesTable", issues)
}
