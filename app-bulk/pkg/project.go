package pkg

import "database/sql"

func ReadProjects(db *sql.DB) ([]*Project, error) {
	query := `
	SELECT id, title
	  FROM project;
	`

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}

	projects := []*Project{}
	for rows.Next() {
		project := &Project{}
		rows.Scan(&project.Id, &project.Title)
		projects = append(projects, project)
	}

	return projects, err
}

func ReadProject(db *sql.DB, id string) (*Project, error) {
	query := `
	SELECT id, title
	  FROM project
	 WHERE id = $1;
    `

	project := &Project{}
	err := db.QueryRow(query, id).Scan(&project.Id, &project.Title)
	if err != nil {
		return nil, err
	}

	return project, nil
}

func ReadIssues(db *sql.DB, projectId string) ([]*Issue, error) {
	query := `
	SELECT issue.id,
	       issue.title,
	       issue.description,
	       issue.project_id,
	       status.name,
	       priority.name
	  FROM issue
	  JOIN status   ON issue.status_id = status.id
	  JOIN priority ON issue.priority_id = priority.id
	 WHERE issue.project_id = $1
	 ORDER BY issue.created_at DESC;
    `

	rows, err := db.Query(query, projectId)
	if err != nil {
		return nil, err
	}

	issues := []*Issue{}
	for rows.Next() {
		issue := &Issue{}
		rows.Scan(&issue.Id, &issue.Title, &issue.Description, &issue.ProjectId, &issue.StatusName, &issue.PriorityName)
		issues = append(issues, issue)
	}

	return issues, err
}
