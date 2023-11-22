#!/bin/bash

REPO_NAME=$1
USER_NAME=$2
GITHUB_TOKEN=$3

# Delete repository via GitHub API
curl -X DELETE -H "Authorization: token $GITHUB_TOKEN" "https://api.github.com/repos/$USER_NAME/$REPO_NAME"

# Remove resource from Terraform state
cd ../../../deployments/terraform
terraform state rm github_repository.$REPO_NAME
