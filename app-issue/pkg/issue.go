package pkg

import (
	"database/sql"
)

func ReadIssue(db *sql.DB, id string) (*Issue, error) {
	query := `
	SELECT issue.id,
	       issue.title,
	       issue.description,
	       issue.project_id,
	       status.name,
	       priority.name,
	       branch.name,
	       branch.url,
	       issue.created_at
	  FROM issue
	  JOIN status   ON issue.status_id = status.id
	  JOIN priority ON issue.priority_id = priority.id
	  JOIN branch   ON issue.branch_id = branch.id
	 WHERE issue.id = $1;
	`

	issue := &Issue{}
	err := db.QueryRow(query, id).Scan(
		&issue.Id,
		&issue.Title,
		&issue.Description,
		&issue.ProjectId,
		&issue.StatusName,
		&issue.PriorityName,
		&issue.BranchName,
		&issue.BranchUrl,
		&issue.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return issue, nil
}

func ReadComments(db *sql.DB, issueId string) ([]*Comment, error) {
	query := `
	SELECT id, issue_id, content, created_at
	  FROM comment
	 WHERE issue_id = $1
	 ORDER BY created_at DESC;
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

func ReadStatuses(db *sql.DB) ([]*Status, error) {
	query := `
	SELECT id, name
	  FROM status
	 ORDER BY display_order;
	`
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	statuses := []*Status{}
	for rows.Next() {
		status := &Status{}
		rows.Scan(&status.Id, &status.Name)
		statuses = append(statuses, status)
	}
	return statuses, err
}

func ReadPriorities(db *sql.DB) ([]*Priority, error) {
	query := `
	SELECT id, name
	  FROM priority
	 ORDER BY display_order;
	`
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	priorities := []*Priority{}
	for rows.Next() {
		priority := &Priority{}
		rows.Scan(&priority.Id, &priority.Name)
		priorities = append(priorities, priority)
	}
	return priorities, err
}

func ReadBranches(db *sql.DB) ([]*Branch, error) {
	query := `
	SELECT id, name
	  FROM branch;
	`
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	branches := []*Branch{}
	for rows.Next() {
		branch := &Branch{}
		rows.Scan(&branch.Id, &branch.Name)
		branches = append(branches, branch)
	}
	return branches, err
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

func CreateNewComment(db *sql.DB, comment Comment) (*Comment, error) {
	query := `
	INSERT INTO comment (issue_id, content, created_at)
	  VALUES ($1, $2, NOW())
	 RETURNING id, issue_id, content, created_at
	`
	newComment := &Comment{}

	err := db.QueryRow(query, comment.IssueId, comment.Content).Scan(
		&newComment.Id,
		&newComment.IssueId,
		&newComment.Content,
		&newComment.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return newComment, err
}
