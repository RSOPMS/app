package issue

import (
	"app-issue/pkg"
	"app-issue/template"
	"database/sql"
	_ "embed"
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

func (h *IssueHandler) GetIssueList(w http.ResponseWriter, r *http.Request) error {
	issues, err := pkg.ReadIssues(h.Db)
	if err != nil {
		return err
	}

	return template.RenderIssue(w, "list", issues)
}
