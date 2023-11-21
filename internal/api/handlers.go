package api

import (
	"net/http"

	"github.com/NunoCG/devflowapp/pkg/github"
)

func SlackEventHandler(w http.ResponseWriter, r *http.Request) {
	// Extract repo_name and repo_description from Slack event
	repoName := "received_repo_name"
	repoDescription := "received_repo_description"

	// Trigger GitHub Actions workflow
	github.TriggerWorkflow(repoName, repoDescription)
}

// Other handler functions related to your API endpoints
