package devflowapp

import (
	"net/http"
)

type Server struct {
	mux *http.ServeMux
}

func NewHandler() http.Handler {
	s := &Server{
		mux: http.NewServeMux(),
	}
	s.routes()
	return s.mux
}

func (s *Server) routes() {
	s.mux.HandleFunc("/createRepo", s.handleCreateRepo())
}

// server.go --> Defines the HTTP server, routes, middleware, etc.
