package pkg

import (
	"database/sql"
	"encoding/json"
	"log"
	"os"

	"github.com/nats-io/nats.go"
)

type ProjectInput struct {
	Title string `json:"title"`
}

type IssueInput struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Project     string `json:"project"`
	Status      string `json:"status"`
	Priority    string `json:"priority"`
	Branch      string `json:"branch"`
}

type InputPayload struct {
	Projects []ProjectInput `json:"projects"`
	Issues   []IssueInput   `json:"issues"`
}

var nc *nats.Conn

// Initialize NATS connection
func InitNATS() error {
	var err error
	natsUrl := os.Getenv("URL_NATS")
	nc, err = nats.Connect(natsUrl)
	if err != nil {
		return err
	}
	return nil
}

// Close the NATS connection
func CloseNATSConnection() {
	if nc != nil {
		nc.Close()
	}
}

func SubscribeToMessages(db *sql.DB) {
	// Ensure NATS is connected before subscribing
	if nc == nil {
		log.Fatalf("NATS connection is not initialized")
		return
	}

	// Subscribe to the project creation topic
	_, err := nc.Subscribe("app.ingress.project", func(msg *nats.Msg) {
		var project ProjectInput
		err := json.Unmarshal(msg.Data, &project)
		if err != nil {
			log.Printf("Error unmarshaling project data: %v", err)
			return
		}

		// Insert the project into the database
		AddProjectToDB(db, project)
	})

	if err != nil {
		log.Fatalf("Error subscribing to app.ingress.project: %v", err)
	}

	// Subscribe to the issue creation topic
	_, err = nc.Subscribe("app.ingress.issue", func(msg *nats.Msg) {
		var issue IssueInput
		err := json.Unmarshal(msg.Data, &issue)
		if err != nil {
			log.Printf("Error unmarshaling issue data: %v", err)
			return
		}

		// Insert the issue into the database
		AddIssueToDB(db, issue)
	})

	if err != nil {
		log.Fatalf("Error subscribing to app.ingress.issue: %v", err)
	}
}

func AddProjectToDB(db *sql.DB, project ProjectInput) error {
	err := CreateNewProject(db, project.Title)
	return err
}

func AddIssueToDB(db *sql.DB, issue IssueInput) error {
	projectId, err := GetProjectIdByTitle(db, issue.Project)
	if err != nil {
		return err
	}

	statusId, err := GetStatusIdByName(db, issue.Status)
	if err != nil {
		return err
	}

	priorityId, err := GetPriorityIdByName(db, issue.Priority)
	if err != nil {
		return err
	}

	branchId, err := GetBranchIdByName(db, issue.Branch)
	if err != nil {
		return err
	}

	err = CreateNewIssue(db, issue.Title, issue.Description, projectId, statusId, priorityId, branchId)
	return err
}

func CreateNewProject(db *sql.DB, title string) error {
	query := `
	INSERT INTO project (title)
	  VALUES ($1);
	`

	_, err := db.Exec(query, title)
	return err
}

func CreateNewIssue(db *sql.DB, title string, desc string, pId int, sId int, prId int, bId int) error {
	query := `
	INSERT INTO issue (title, description, project_id, status_id, priority_id, branch_id, created_at)
	  VALUES ($1, $2, $3, $4, $5, $6, NOW());
	`

	_, err := db.Exec(query, title, desc, pId, sId, prId, bId)

	return err
}

func GetProjectIdByTitle(db *sql.DB, title string) (int, error) {
	query := `
	SELECT id
	  FROM project
	 WHERE title = $1;
	`

	var id int

	err := db.QueryRow(query, title).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func GetPriorityIdByName(db *sql.DB, name string) (int, error) {
	query := `
	SELECT id
	  FROM priority
	 WHERE name = $1;
	`

	var id int

	err := db.QueryRow(query, name).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func GetStatusIdByName(db *sql.DB, name string) (int, error) {
	query := `
	SELECT id
	  FROM status
	 WHERE name = $1;
	`

	var id int

	err := db.QueryRow(query, name).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func GetBranchIdByName(db *sql.DB, name string) (int, error) {
	query := `
	SELECT id
	  FROM branch
	 WHERE name = $1;
	`

	var id int

	err := db.QueryRow(query, name).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}
