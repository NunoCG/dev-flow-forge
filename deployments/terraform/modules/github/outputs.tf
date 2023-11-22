output "created_repository_urls" {
  value = [for repo in local.created_repos : repo.html_url]
}
