package main

import (
	"fmt"
	"log"
	"os"

	"github.com/NunoCG/dev-flow-forge/internal/github"
)

func main() {
	repoName := "example-repo"
	repoDescription := "This is a sample repository."

	token := os.Getenv("TF_VAR_GITHUB_TOKEN")
	if token == "" {
		log.Fatal("GitHub token not set. Please set TF_VAR_GITHUB_TOKEN environment variable.")
	}

	client := github.NewGitHubClient(token)
	err := client.CreateRepository(repoName, repoDescription)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Repository %s created successfully!\n", repoName)
}
