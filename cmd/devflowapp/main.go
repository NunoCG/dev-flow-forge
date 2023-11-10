package main

import (
	"log"
	"net/http"

	"github.com/NunoCG/dev-flow-forge/internal/app/devflowapp"
)

func main() {
	handler := devflowapp.NewHandler()
	server := &http.Server{
		Addr:    ":8080",
		Handler: handler,
	}

	log.Println("Server is starting...")
	err := server.ListenAndServe()
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
