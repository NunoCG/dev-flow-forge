module "github_repo" {
  source                 = "./modules/github"
  repo_name              = var.repo_name
  repo_description       = var.repo_description
  TF_VAR_github_token    = var.TF_VAR_github_token
  # github_actions_secrets = {
  #   "secret_name_1" = "secret_value_1",
  #   "secret_name_2" = "secret_value_2",
  #   # Add more secrets as needed
  # }
}
