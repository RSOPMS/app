package pkg

import (
	"database/sql"
)

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

func ReadStatuses(db *sql.DB) ([]*Status, error) {
	query := `
	SELECT id, name, display_order
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
		rows.Scan(&status.Id, &status.Name, &status.DisplayOrder)
		statuses = append(statuses, status)
	}
	return statuses, err
}

func ReadPriorities(db *sql.DB) ([]*Priority, error) {
	query := `
	SELECT id, name, display_order
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
		rows.Scan(&priority.Id, &priority.Name, &priority.DisplayOrder)
		priorities = append(priorities, priority)
	}
	return priorities, err
}

func ReadBranches(db *sql.DB) ([]*Branch, error) {
	query := `
	SELECT id, name, url
	  FROM branch;
	`
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	branches := []*Branch{}
	for rows.Next() {
		branch := &Branch{}
		rows.Scan(&branch.Id, &branch.Name, &branch.Url)
		branches = append(branches, branch)
	}
	return branches, err
}

func CreateNewIssue(db *sql.DB, title string, description string, projectId string, statusId string, priorityId string, branchId string) (*Issue, error) {
	query := `
	INSERT INTO issue (title, description, project_id, status_id, priority_id, branch_id, created_at)
	  VALUES ($1, $2, $3, $4, $5, $6, NOW())
	 RETURNING id, title, description, project_id, status_id, priority_id, branch_id, created_at
	`
	newIssue := &Issue{}

	err := db.QueryRow(query, title, description, projectId, statusId, priorityId, branchId).Scan(
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
