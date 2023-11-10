package github

import (
	"context"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

type Client struct {
	ghClient *github.Client
}

func NewClient(ghToken string) *Client {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: ghToken},
	)
	tc := oauth2.NewClient(context.Background(), ts)

	return &Client{
		ghClient: github.NewClient(tc),
	}
}

func (c *Client) CreateRepo(name string, private bool) (*github.Repository, error) {
	repo, _, err := c.ghClient.Repositories.Create(context.Background(), "", &github.Repository{
		Name:    github.String(name),
		Private: github.Bool(private),
	})
	return repo, err
}

/*
 * This would be where to interact with the GitHub API.
 * Use a GitHub client library for Go to do this, which would take care of things like making HTTP requests
 * to the GitHub API, authentication, etc.
 */
