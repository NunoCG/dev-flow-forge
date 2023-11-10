variable "repo_name" {
  description = "The name of the GitHub repository"
  default     = "default_repo_name"
  type        = string
}

variable "repo_description" {
  description = "The description of the GitHub repository"
  default     = "Default repo description."
  type        = string
}

variable "repo_visibility" {
  description = "The visibility of the GitHub repository (public or private)"
  default     = "public"
  type        = string
}

variable "github_token" {
  description = "GitHub personal access token with full access"
  type        = string
  default     = lookup(env, "GITHUB_TOKEN", "")
  sensitive   = true
}

variable "github_actions_secrets" {
  description = "Map of GitHub Actions secrets"
  type        = map(string)
  default     = {}
}
