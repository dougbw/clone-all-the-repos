package command

import (
	"clone-all-the-repos/internal/config"
	"clone-all-the-repos/internal/logger"
	"clone-all-the-repos/internal/providers/azuredevops"
	"clone-all-the-repos/internal/providers/github"
	"clone-all-the-repos/internal/providers/gitlab"
)

func DiscoveryCommand(startup Startup) (repos []config.GitRepo) {

	logger.Context = []string{
		"discovery",
	}
	logger.Print("🔍 Starting discovery...")

	configuration := startup.Config

	// azure devops
	for _, AzureDevopsConfig := range configuration.AzureDevOps {
		var AzdoRepos = azuredevops.GetRepos(AzureDevopsConfig)
		repos = append(repos, AzdoRepos...)
	}

	// github
	for _, GitHubConfig := range configuration.GitHub {
		var GitHubRepos = github.GetRepos(GitHubConfig)
		repos = append(repos, GitHubRepos...)
	}

	// gitlab
	for _, GitLabConfig := range configuration.GitLab {
		var GitLabRepos = gitlab.GetRepos(GitLabConfig)
		repos = append(repos, GitLabRepos...)
	}

	return repos

}
