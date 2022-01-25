package main

import (
	"clone-all-the-repos/internal/config"
	"clone-all-the-repos/internal/git"
	"clone-all-the-repos/internal/logger"
	"clone-all-the-repos/internal/providers/azuredevops"
	"clone-all-the-repos/internal/providers/github"
	"clone-all-the-repos/internal/providers/gitlab"
	"fmt"
	"os"
	"path/filepath"
)

func main() {

	var configPath string
	if len(os.Args[1:]) == 1 {
		configPath = os.Args[1]
	} else {
		fmt.Println("usage:\nclone-all-the-repos <configfile>")
		os.Exit(1)
	}

	absolute, _ := filepath.Abs(configPath)
	dir, file := filepath.Split(absolute)
	baseDirectory := dir
	configuration := config.OpenYaml(configPath)

	context := []string{
		fmt.Sprintf("config:%s", file),
	}
	logger.PrintLogMessage(context, "🔍 Finding repos")

	var Repos []config.GitRepo

	// azure devops
	for _, AzureDevopsConfig := range configuration.AzureDevOps {
		var AzdoRepos = azuredevops.GetAzureDevOpsRepos(AzureDevopsConfig)
		Repos = append(Repos, AzdoRepos...)
	}

	// github
	for _, GitHubConfig := range configuration.GitHub {
		var GitHubRepos = github.GetGitHubRepos(GitHubConfig)
		Repos = append(Repos, GitHubRepos...)
	}

	// gitlab
	for _, GitLabConfig := range configuration.GitLab {
		var GitHubRepos = gitlab.GetGitLabRepos(GitLabConfig)
		Repos = append(Repos, GitHubRepos...)
	}

	// clone
	message := fmt.Sprintf("⌛ Cloning '%d' repos", len(Repos))
	logger.PrintLogMessage(context, message)

	for _, Repo := range Repos {
		git.CloneGitRepo(baseDirectory, Repo)
	}

	logger.PrintLogMessage(context, "Done")

}
