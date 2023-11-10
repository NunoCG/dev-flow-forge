variable "repo_name" {
  description = "The name of the GitHub repository"
  type        = string
}

variable "repo_description" {
  description = "The description of the GitHub repository"
  type        = string
}

variable "github_token" {
  description = "GitHub personal access token with full access"
  type        = string
  default     = lookup(env, "GITHUB_TOKEN", "")
  sensitive   = true
}
