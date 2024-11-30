package issue

import (
	"app-issue/pkg"
	"app-issue/template"
	"database/sql"
	"net/http"
)

type IssueHandler struct {
	Db *sql.DB
}

func NewIssueHandler(db *sql.DB) *IssueHandler {
	return &IssueHandler{
		Db: db,
	}
}

func (h *IssueHandler) GetIssue(w http.ResponseWriter, r *http.Request) error {
	return template.RenderLayout(w, "issue", nil)
}

func (h *IssueHandler) GetProjectIssues(w http.ResponseWriter, r *http.Request) error {
	project_id := r.PathValue("id")
	project, err := pkg.ReadProject(h.Db, project_id)

	if err != nil {
		return err
	}

	return template.RenderLayout(w, "project_issues", map[string]interface{}{
		"ProjectId":    project_id,
		"ProjectTitle": project.Title,
	})
}

func (h *IssueHandler) GetProjectIssueList(w http.ResponseWriter, r *http.Request) error {
	project_id := r.PathValue("id")

	issues, err := pkg.ReadProjectIssues(h.Db, project_id)
	if err != nil {
		return err
	}

	return template.RenderIssue(w, "table", issues)
}

func (h *IssueHandler) GetIssueComments(w http.ResponseWriter, r *http.Request) error {
	issue_id := r.PathValue("i_id")

	issue, err := pkg.ReadIssue(h.Db, issue_id)
	if err != nil {
		return err
	}

	return template.RenderLayout(w, "issue_page", issue)
}

func (h *IssueHandler) GetIssueCommentList(w http.ResponseWriter, r *http.Request) error {
	issue_id := r.PathValue("i_id")

	comments, err := pkg.ReadIssueComments(h.Db, issue_id)
	if err != nil {
		return err
	}

	return template.RenderComment(w, "table", comments)
}
