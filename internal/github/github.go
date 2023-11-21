package github

import (
	"context"
	"fmt"

	"github.com/google/go-github/v35/github"
	"golang.org/x/oauth2"
)

type GitHubClient struct {
	client *github.Client
}

func NewGitHubClient(token string) *GitHubClient {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)

	return &GitHubClient{
		client: github.NewClient(tc),
	}
}

func (gc *GitHubClient) CreateRepository(repoName, repoDescription string) error {
	ctx := context.Background()
	repo := &github.Repository{
		Name:        github.String(repoName),
		Description: github.String(repoDescription),
		// Add more settings as needed
	}

	_, _, err := gc.client.Repositories.Create(ctx, "", repo)
	return err
}

func (gc *GitHubClient) DeleteRepository(repoName string) error {
	ctx := context.Background()
	_, err := gc.client.Repositories.Delete(ctx, "", repoName)
	if err != nil {
		return fmt.Errorf("failed to delete repository: %v", err)
	}
	return nil
}
