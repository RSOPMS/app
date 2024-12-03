package pkg

import "database/sql"

type Issue struct {
	Id          int
	Title       string
	Description string
	ProjectId   int
}

type Comment struct {
	Id        int
	IssueId   int
	Content   string
	CreatedAt string
}

// Reads all issues
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

// Reads an issue by Id
func ReadIssue(db *sql.DB, id string) (*Issue, error) {
	query := `
	SELECT id, title, description, project_id
	  FROM issue
	 WHERE id = $1;
	`

	issue := &Issue{}
	err := db.QueryRow(query, id).Scan(&issue.Id, &issue.Title, &issue.Description, &issue.ProjectId)
	if err != nil {
		return nil, err
	}

	return issue, nil
}

// Reads all comments of an issue by Id
func ReadIssueComments(db *sql.DB, issueId string) ([]*Comment, error) {
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
