variable "repo_name" {
  description = "The name of the GitHub repository"
  default     = "default_repo_name"
}

variable "repo_description" {
  description = "The description of the GitHub repository"
  default     = "Default repo description."
}

variable "repo_visibility" {
  description = "The visibility of the GitHub repository (public or private)"
  default     = "public"
}

variable "github_actions_secrets" {
  description = "Map of GitHub Actions secrets"
  type        = map(string)
  default     = {}
}