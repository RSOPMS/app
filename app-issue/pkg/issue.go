package pkg

import "database/sql"

func ReadIssue(db *sql.DB, id string) (*Issue, error) {
	query := `
	SELECT issue.id, issue.title, issue.description, issue.project_id, status.name, priority.name
	  FROM issue
	  JOIN status ON issue.status_id = status.id
	  JOIN priority ON issue.priority_id = priority.id
	 WHERE issue.id = $1;
	`

	issue := &Issue{}
	err := db.QueryRow(query, id).Scan(&issue.Id, &issue.Title, &issue.Description, &issue.ProjectId, &issue.StatusName, &issue.PriorityName)
	if err != nil {
		return nil, err
	}

	return issue, nil
}

func ReadComments(db *sql.DB, issueId string) ([]*Comment, error) {
	query := `
	SELECT id, issue_id, content, created_at
	  FROM comment
	 WHERE issue_id = $1;
	`

	rows, err := db.Query(query, issueId)
	if err != nil {
		return nil, err
	}

	comments := []*Comment{}
	for rows.Next() {
		comment := &Comment{}
		rows.Scan(&comment.Id, &comment.IssueId, &comment.Content, &comment.CreatedAt)
		comments = append(comments, comment)
	}

	return comments, err
}
