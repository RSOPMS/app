package pkg

import (
	"database/sql"
)

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

func AddPayloadToDB(db *sql.DB, payload InputPayload) error {
	for _, project := range payload.Projects {
		err := CreateNewProject(db, project.Title)
		if err != nil {
			return err
		}
	}

	for _, issue := range payload.Issues {
		projectId, err := GetProjectIdByTitle(db, issue.Project)
		if err != nil {
			return err
		}

		statusId, err := GetStatusIdByName(db, issue.Status)
		if err != nil {
			return err
		}

		priorityId, err := GetPriorityIdByName(db, issue.Priority)
		if err != nil {
			return err
		}

		branchId, err := GetBranchIdByName(db, issue.Branch)
		if err != nil {
			return err
		}

		err = CreateNewIssue(db, issue.Title, issue.Description, projectId, statusId, priorityId, branchId)
		if err != nil {
			return err
		}
	}

	return nil
}

func CreateNewProject(db *sql.DB, title string) error {
	query := `
	INSERT INTO project (title)
	  VALUES ($1);
	`

	_, err := db.Exec(query, title)
	return err
}

func CreateNewIssue(db *sql.DB, title string, desc string, pId int, sId int, prId int, bId int) error {
	query := `
	INSERT INTO issue (title, description, project_id, status_id, priority_id, branch_id, created_at)
	  VALUES ($1, $2, $3, $4, $5, $6, NOW());
	`

	_, err := db.Exec(query, title, desc, pId, sId, prId, bId)

	return err
}

func GetProjectIdByTitle(db *sql.DB, title string) (int, error) {
	query := `
	SELECT id
	  FROM project
	 WHERE title = $1;
	`

	var id int

	err := db.QueryRow(query, title).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func GetPriorityIdByName(db *sql.DB, name string) (int, error) {
	query := `
	SELECT id
	  FROM priority
	 WHERE name = $1;
	`

	var id int

	err := db.QueryRow(query, name).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func GetStatusIdByName(db *sql.DB, name string) (int, error) {
	query := `
	SELECT id
	  FROM status
	 WHERE name = $1;
	`

	var id int

	err := db.QueryRow(query, name).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func GetBranchIdByName(db *sql.DB, name string) (int, error) {
	query := `
	SELECT id
	  FROM branch
	 WHERE name = $1;
	`

	var id int

	err := db.QueryRow(query, name).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}
