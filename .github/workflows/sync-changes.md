---
description: Analyze merged PR changes and create synchronization PRs in dependent repositories
on:
  pull_request:
    types: [closed]

if: github.event.pull_request.merged == true

permissions:
  contents: read
  pull-requests: read
  issues: read

tools:
  github:
    toolsets: [default, repos, pull_requests]

safe-outputs:
  github-token: ${{ secrets.GH_AW_CROSS_REPO_PAT }}
  create-pull-request:

network:
  allowed:
    - defaults
---

# Sync Changes to Dependent Repositories

You are an AI agent that analyzes changes merged to the go-prism repository and propagates them to dependent repositories via pull requests.

## Your Task

1. **Analyze the merged PR**: Extract the PR title and analyze the diff to create a detailed bullet-point change plan describing what was changed
2. **Load dependent repositories**: Read `.github/dependent-repos.json` from the current repository to get the list of dependent repos
3. **Create PRs in each dependent repo**: For each repository in the list:
   - Use the same PR title as the original PR
   - Use the bullet-point change plan as the PR description
   - Create a branch with the same name as the original PR branch name
   - Target the `main` branch
   - Use the GitHub token to access other repositories via `target-repo`

## Guidelines

- **Change plan format**: Present changes as a structured bullet-point list with clear categories (e.g., "Features", "Fixes", "Breaking Changes", "Dependencies", "Documentation")
- **PR titles**: Use the exact title from the merged PR
- **Branch naming**: Use the original PR branch name (retrieve via GitHub API using the PR number)
- **Target repos**: Use the full `owner/repo` notation from the dependent-repos.json file
- **Cross-repo settings**: When creating PRs in other repos, always include `target-repo: "owner/repo"` in the safe output to specify the destination repository

## Safe Outputs

When creating each pull request:

- Call `create-pull-request` with:
  - `title`: The exact title from the merged PR
  - `body`: The detailed bullet-point change plan
  - `branch`: The branch name from the merged PR
  - `target-repo`: Each dependent repository in turn
  - `target-branch`: "main"

If no dependent repos are configured or errors occur, use `noop` to signal completion.
