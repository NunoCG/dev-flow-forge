module "github_repo" {
  source              = "./modules/github"
  repos_to_create     = var.repos_to_create
  TF_VAR_github_token = var.TF_VAR_github_token
}
