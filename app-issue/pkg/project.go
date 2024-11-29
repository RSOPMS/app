package pkg

import "database/sql"

type Project struct {
	Id    int
	Title string
}

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

func ReadProjectTitle(db *sql.DB, project_id string) (*Project, error) {
	query := `
	SELECT id, title
	  FROM project
	  WHERE id = $1;
	`

	project := &Project{}
	err := db.QueryRow(query, project_id).Scan(&project.Id, &project.Title)
	if err != nil {
		return nil, err
	}

	return project, nil
}
