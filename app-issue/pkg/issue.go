package pkg

import "database/sql"

type Issue struct {
	Id    		int
	Title       string
	Description string
	ProjectId 	int
}

func ReadIssues(db *sql.DB) ([]*Issue, error) {
	query := `
	SELECT id, title, description, project_id
	  FROM issue;
	`

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}

	issues := []*Issue{}
	for rows.Next() {
		issue := &Issue{}
		rows.Scan(&issue.Id, &issue.Title, &issue.Description, &issue.ProjectId)
		issues = append(issues, issue)
	}

	return issues, err
}

func ReadProjectIssues(db *sql.DB, project_id string) ([]*Issue, error) {
	query := `
	SELECT id, title, description, project_id
	  FROM issue
	  WHERE project_id = $1;
	`

	rows, err := db.Query(query, project_id)
	if err != nil {
		return nil, err
	}

	issues := []*Issue{}
	for rows.Next() {
		issue := &Issue{}
		rows.Scan(&issue.Id, &issue.Title, &issue.Description, &issue.ProjectId)
		issues = append(issues, issue)
	}

	return issues, err
}
