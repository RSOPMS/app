package pkg

import "database/sql"

type Issue struct {
	Title       string
	Description string
}

func ReadIssues(db *sql.DB) ([]*Issue, error) {
	query := `
	SELECT title, description
	  FROM issue;
	`

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}

	issues := []*Issue{}
	for rows.Next() {
		issue := &Issue{}
		rows.Scan(&issue.Title, &issue.Description)
		issues = append(issues, issue)
	}

	return issues, err
}
