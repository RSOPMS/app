package pkg

import "database/sql"

type Project struct {
	ID    int
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
		rows.Scan(&project.ID, &project.Title)
		projects = append(projects, project)
	}

	return projects, err
}
