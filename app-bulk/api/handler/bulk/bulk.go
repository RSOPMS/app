package bulk

import (
	"app-bulk/pkg"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
)

// Input structs for JSON decoding
type ProjectInput struct {
	Title string `json:"title"`
}

type IssueInput struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Project     string `json:"project"`
	Status      string `json:"status"`
	Priority    string `json:"priority"`
	Branch      string `json:"branch"`
}

type InputPayload struct {
	Projects []ProjectInput `json:"projects"`
	Issues   []IssueInput   `json:"issues"`
}

type BulkHandler struct {
	Db *sql.DB
}

func NewBulkHandler(db *sql.DB) *BulkHandler {
	return &BulkHandler{
		Db: db,
	}
}

func (h *BulkHandler) PostBulk(w http.ResponseWriter, r *http.Request) error {
	log.Println("Gets to the handler")
	var payload InputPayload
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		return err
	}

	for _, project := range payload.Projects {
		// check if it exists? or change title to unique
		_, err := pkg.CreateNewProject(h.Db, project.Title)
		if err != nil {
			return err
		}
	}

	for _, issue := range payload.Issues {

		projectId, err := pkg.GetProjectIdByTitle(h.Db, issue.Project)
		if err != nil {
			return err
		}

		statusId, err := pkg.GetStatusIdByName(h.Db, issue.Status)
		if err != nil {
			return err
		}

		priorityId, err := pkg.GetPriorityIdByName(h.Db, issue.Priority)
		if err != nil {
			return err
		}

		branchId, err := pkg.GetBranchIdByName(h.Db, issue.Branch)
		if err != nil {
			return err
		}

		// probably needs some validation
		createdIssue := pkg.Issue{
			Title:       issue.Title,
			Description: issue.Description,
			ProjectId:   projectId,
			StatusId:    statusId,
			PriorityId:  priorityId,
			BranchId:    branchId,
		}

		_, err = pkg.CreateNewIssue(h.Db, createdIssue)
		if err != nil {
			return err
		}
	}

	// Respond with success
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Bulk insert successful"))
	return nil
}
