package pkg

import "database/sql"

type Comment struct {
	Id    		int
	IssueId		int
	Content     string
	CreatedAt 	string
}

func ReadComments(db *sql.DB) ([]*Comment, error) {
	query := `
	SELECT id, issue_id, content, created_at
	  FROM comment;
	`

	rows, err := db.Query(query)
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

func ReadIssueComments(db *sql.DB, issue_id string) ([]*Comment, error) {
	query := `
	SELECT id, issue_id, content, created_at
	  FROM comment
	  WHERE issue_id = $1;
	`

	rows, err := db.Query(query, issue_id)
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
