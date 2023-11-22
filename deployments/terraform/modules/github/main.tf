locals {
  created_repos = [for repo in var.repos_to_create : github_repository.repo_config[repo.name]]
}

resource "github_repository" "repo_config" {
  for_each = { for repo in var.repos_to_create : repo.name => repo }

  name        = each.value.name
  description = each.value.description
  visibility  = each.value.visibility
}

# resource "null_resource" "delete_repos" {
#   for_each = { for repo_name in var.repos_to_delete : repo_name => repo_name }

#   triggers = {
#     repo_name = each.value
#   }

#   provisioner "local-exec" {
#     command = "curl -X DELETE -H 'Authorization: token ${var.TF_VAR_github_token}' https://api.github.com/repos/${each.value}"
#   }
# }
