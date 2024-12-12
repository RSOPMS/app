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

func (h *IssueHandler) CreateNewIssue(w http.ResponseWriter, r *http.Request) error {
	// Parse form values
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Failed to parse form data", http.StatusBadRequest)
		return err
	}
	// Extract values
	title := r.FormValue("title")
	description := r.FormValue("description")
	projectID := r.FormValue("project_id")
	statusID := r.FormValue("status_id")
	priorityID := r.FormValue("priority_id")
	branchID := r.FormValue("branch_id")

	query := `
		INSERT INTO issues (title, description, project_id, status_id, priority_id, branch_id, created_at)
		VALUES ($1, $2, $3, $4, $5, $6, NOW())
		RETURNING id, title, description, project_id, status_id, priority_id, branch_id, created_at
	`

	var newIssue struct {
		ID          int
		Title       string
		Description string
		ProjectID   int
		StatusID    int
		PriorityID  int
		BranchID    int
		CreatedAt   string // Timestamp in a readable format
	}

	err = h.Db.QueryRow(query, title, description, projectID, statusID, priorityID, branchID).Scan(
		&newIssue.ID,
		&newIssue.Title,
		&newIssue.Description,
		&newIssue.ProjectID,
		&newIssue.StatusID,
		&newIssue.PriorityID,
		&newIssue.BranchID,
		&newIssue.CreatedAt,
	)
	if err != nil {
		http.Error(w, "Failed to create issue", http.StatusInternalServerError)
		return err
	}

	issues, err := pkg.ReadIssues(h.Db, projectID)
	if err != nil {
		return err
	}

	return template.RenderIssue(w, "issuesTable", issues)
}
