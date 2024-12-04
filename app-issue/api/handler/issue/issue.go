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
	issueId := r.PathValue("issueId")

	issue, err := pkg.ReadIssue(h.Db, issueId)
	if err != nil {
		return err
	}

	return template.RenderLayout(w, "issuePage", issue)
}

func (h *IssueHandler) GetCommentsTable(w http.ResponseWriter, r *http.Request) error {
	issueId := r.PathValue("issueId")

	comments, err := pkg.ReadComments(h.Db, issueId)
	if err != nil {
		return err
	}

	return template.RenderIssue(w, "commentsTable", comments)
}
