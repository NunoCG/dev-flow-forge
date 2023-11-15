variable "repo_name" {
  description = "The name of the GitHub repository"
  type        = string
}

variable "repo_description" {
  description = "The description of the GitHub repository"
  type        = string
}

variable "TF_VAR_github_token" {
  description = "The github access token with the needed access grants"
  type        = string
  sensitive   = true
}
