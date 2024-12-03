package pkg

import "database/sql"

type Project struct {
	Id    int
	Title string
}

// Reads all projects
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

// Reads a project by Id
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

// Reads all issues of a project by Id
func ReadProjectIssues(db *sql.DB, projectId string) ([]*Issue, error) {
	query := `
	SELECT id, title, description, project_id
	  FROM issue
	 WHERE project_id = $1;
	`

	rows, err := db.Query(query, projectId)
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
