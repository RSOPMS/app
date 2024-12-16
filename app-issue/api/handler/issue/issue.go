package issue

import (
	"app-issue/pkg"
	"app-issue/template"
	"database/sql"
	"net/http"
	"strconv"
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

func (h *IssueHandler) CreateNewIssue(w http.ResponseWriter, r *http.Request) error {
	// Parse form values
	err := r.ParseForm()
	if err != nil {
		return err
	}

	// Extract values
	title := r.FormValue("title")
	description := r.FormValue("description")
	projectID := r.FormValue("project_id")
	statusID := r.FormValue("status_id")
	priorityID := r.FormValue("priority_id")
	branchID := r.FormValue("branch_id")

	newIssue, err := pkg.CreateNewIssue(h.Db, title, description, projectID, statusID, priorityID, branchID)
	if err != nil {
		return err
	}

	issue, err := pkg.ReadIssue(h.Db, strconv.Itoa(newIssue.Id))
	if err != nil {
		return err
	}

	return template.RenderIssue(w, "issueRow", issue)
}

// Get statuses from the database for the Create New Issue form
func (h *IssueHandler) GetStatusesForm(w http.ResponseWriter, r *http.Request) error {
	statuses, err := pkg.ReadStatuses(h.Db)
	if err != nil {
		return err
	}
	return template.RenderIssue(w, "statusesForm", statuses)
}

// Get priorities from the database for the Create New Issue form
func (h *IssueHandler) GetPrioritiesForm(w http.ResponseWriter, r *http.Request) error {
	priorities, err := pkg.ReadPriorities(h.Db)
	if err != nil {
		return err
	}
	return template.RenderIssue(w, "prioritiesForm", priorities)
}

// Get branches from the database for the Create New Issue form
func (h *IssueHandler) GetBranchesForm(w http.ResponseWriter, r *http.Request) error {
	branches, err := pkg.ReadBranches(h.Db)
	if err != nil {
		return err
	}

	return template.RenderIssue(w, "branchesForm", branches)
}
