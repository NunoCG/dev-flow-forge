package devflowapp

import (
	"net/http"
)

func (s *Server) handleCreateRepo() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Parse the request, call the GitHub API to create a repo, etc.
		w.Write([]byte("Repo created successfully"))
	}
}
