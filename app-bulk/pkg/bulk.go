package pkg

import (
	"database/sql"
)

func CreateNewProject(db *sql.DB, title string) (*Project, error) {
	query := `
	INSERT INTO project (title)
	  VALUES ($1)
	 RETURNING id, title
	`
	newProject := &Project{}

	err := db.QueryRow(query, title).Scan(
		&newProject.Id,
		&newProject.Title,
	)

	if err != nil {
		return nil, err
	}

	return newProject, err
}

func CreateNewIssue(db *sql.DB, issue Issue) (*Issue, error) {
	query := `
	INSERT INTO issue (title, description, project_id, status_id, priority_id, branch_id, created_at)
	  VALUES ($1, $2, $3, $4, $5, $6, NOW())
	 RETURNING id, title, description, project_id, status_id, priority_id, branch_id, created_at
	`
	newIssue := &Issue{}

	err := db.QueryRow(query, issue.Title, issue.Description, issue.ProjectId, issue.StatusId, issue.PriorityId, issue.BranchId).Scan(
		&newIssue.Id,
		&newIssue.Title,
		&newIssue.Description,
		&newIssue.ProjectId,
		&newIssue.StatusId,
		&newIssue.PriorityId,
		&newIssue.BranchId,
		&newIssue.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return newIssue, err
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
