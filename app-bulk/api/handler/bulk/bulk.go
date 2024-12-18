package bulk

import (
	"app-bulk/pkg"
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
)

type BulkHandler struct {
	Db *sql.DB
}

func NewBulkHandler(db *sql.DB) *BulkHandler {
	return &BulkHandler{
		Db: db,
	}
}

func (h *BulkHandler) PostBulk(w http.ResponseWriter, r *http.Request) error {
	var payload InputPayload
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		return err
	}

	for _, project := range payload.Projects {
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
		if projectId == 0 {
			return errors.New("missing or invalid project ID")
		}

		statusId, err := pkg.GetStatusIdByName(h.Db, issue.Status)
		if err != nil {
			return err
		}
		if statusId == 0 {
			return errors.New("missing or invalid status ID")
		}

		priorityId, err := pkg.GetPriorityIdByName(h.Db, issue.Priority)
		if err != nil {
			return err
		}
		if priorityId == 0 {
			return errors.New("missing or invalid priority ID")
		}

		branchId, err := pkg.GetBranchIdByName(h.Db, issue.Branch)
		if err != nil {
			return err
		}
		if branchId == 0 {
			return errors.New("missing or invalid branch ID")
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

	return nil
}
