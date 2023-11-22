variable "TF_VAR_github_token" {
  description = "The github access token with the needed access grants"
  type        = string
  sensitive   = true
}

variable "github_actions_secrets" {
  description = "Map of GitHub Actions secrets"
  type        = map(string)
  default     = {}
}

variable "repos_to_create" {
  description = "List of repositories to create"
  type        = list(object({
    name        = string
    description = string
    visibility  = string
    // You can add more attributes as needed
  }))
  default     = []
}

variable "repos_to_delete" {
  description = "List of repositories to delete"
  type        = list(string)
  default     = []
}
