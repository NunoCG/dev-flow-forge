terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 4.0"
    }
    github = {
      source  = "integrations/github"
      version = "~> 5.0"
    }
  }
}

provider "aws" {
  region  = "ap-south-1"
  profile = "terraformprofile"
}

provider "github" {
  token = var.github_token
}

resource "github_repository" "repo_config" {
  name        = var.repo_name
  description = var.repo_description
  visibility  = var.repo_visibility
}

resource "github_actions_secret" "repo_secrets" {
  repository      = github_repository.repo_config.name
  for_each        = var.github_actions_secrets
  secret_name     = each.key
  plaintext_value = each.value
}
