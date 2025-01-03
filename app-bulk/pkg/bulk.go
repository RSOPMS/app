package pkg

import (
	"encoding/json"
	"fmt"
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

// Send a message to a specific NATS subject
func PublishMessage(subject string, message []byte) error {
	if nc == nil {
		return fmt.Errorf("NATS connection is not initialized")
	}
	err := nc.Publish(subject, message)
	if err != nil {
		return err
	}
	return nil
}

func AddPayloadToDB(payload InputPayload) error {
	for _, project := range payload.Projects {
		// Send project info to app-ingress through NATS
		projectData, err := json.Marshal(project)
		if err != nil {
			return fmt.Errorf("error marshaling project data: %w", err)
		}

		err = PublishMessage("app.ingress.project", projectData)
		if err != nil {
			return fmt.Errorf("error sending project data to NATS: %w", err)
		}
	}

	for _, issue := range payload.Issues {
		// Send issue info to app-ingress through NATS
		issueData, err := json.Marshal(issue)
		if err != nil {
			return fmt.Errorf("error marshaling issue data: %w", err)
		}

		err = PublishMessage("app.ingress.issue", issueData)
		if err != nil {
			return fmt.Errorf("error sending issue data to NATS: %w", err)
		}
	}

	return nil
}
