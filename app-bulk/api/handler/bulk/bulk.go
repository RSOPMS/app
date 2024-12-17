package handler

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

func (h *BulkHandler) HandleBulkInsert(w http.ResponseWriter, r *http.Request) error {
	log.Println("Gets to handler")

	var payload InputPayload
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		return err
	}

	for _, project := range payload.Projects {
		/*
			if project.Title == "" {
				return errors.New("project title is required")
			}*/
		// check if it exists? or change title to unique
		if err := pkg.CreateNewProject(h.Db, project.Title); err != nil {
			return err
		}
	}

	for _, issue := range payload.Issues {
		/*if issue.Title == "" || issue.Description == "" || issue.Project == "" ||
			issue.Status == "" || issue.Priority == "" || issue.Branch == "" {
			return errors.New("all fields are required in an issue")
		}*/

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

		newIssue, err := pkg.CreateNewIssue(h.Db, createdIssue)
		if err != nil {
			return err
		}
	}

	// Respond with success
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Bulk insert successful"))
	return nil
}
