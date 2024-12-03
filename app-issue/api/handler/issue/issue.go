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

func (h *IssueHandler) GetIssuePage(w http.ResponseWriter, r *http.Request) error {
	issue_id := r.PathValue("issue_id")
	issue, err := pkg.ReadIssue(h.Db, issue_id)
	if err != nil {
		return err
	}

	return template.RenderLayout(w, "issue", issue)
}

func (h *IssueHandler) GetCommentTable(w http.ResponseWriter, r *http.Request) error {
	issue_id := r.PathValue("issue_id")
	comments, err := pkg.ReadIssueComments(h.Db, issue_id)
	if err != nil {
		return err
	}

	return template.RenderIssue(w, "table-comments", comments)
}
