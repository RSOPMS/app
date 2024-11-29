package pkg

import "database/sql"

type Project struct {
	Title string
}

func ReadProjects(db *sql.DB) ([]*Project, error) {
	query := `
	SELECT title
	  FROM project;
	`

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}

	projects := []*Project{}
	for rows.Next() {
		project := &Project{}
		rows.Scan(&project.Title)
		projects = append(projects, project)
	}

	return projects, err
}
