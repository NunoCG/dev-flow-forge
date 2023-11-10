package config

import (
	"github.com/joho/godotenv"

	"log"
	"os"
)

type Config struct {
	GitHubToken string
}

func LoadConfig() Config {
	err := godotenv.Load() // Load environment variables from the .env file
	if err != nil {
		// TODO: Handle errors properly, such as failing gracefully or logging the error
		log.Fatalf("Error loading .env file: %v", err) // Log the error and exit the program
	}

	return Config{
		GitHubToken: os.Getenv("GITHUB_TOKEN"), // Load GitHub token from the "GITHUB_TOKEN" environment variable
	}
}

// This is where to define the application's configuration.
