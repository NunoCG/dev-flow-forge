package api

import (
	"fmt"
	"os/exec"
)

func TriggerWorkflow(repoName, repoDescription string) {
	// Command to run Terraform with updated variables
	cmd := exec.Command("terraform", "apply", "-auto-approve", "-input=false", fmt.Sprintf("-var=repo_name=%s", repoName), fmt.Sprintf("-var=repo_description=%s", repoDescription))
	
	// Run the command and handle errors
	err := cmd.Run()
	if err != nil {
		// Handle command execution error
	}
}

