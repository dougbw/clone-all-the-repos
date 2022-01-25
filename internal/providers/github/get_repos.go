package github

import (
	"clone-all-the-repos/internal/config"
	"clone-all-the-repos/internal/logger"
	"fmt"
	"os"
	"path"
)

func GetGitHubRepos(GitHubConfig config.GitHubConfig) (repos []config.GitRepo) {

	loggingContext := []string{
		"provider:github",
		fmt.Sprintf("config:%s", GitHubConfig.Name),
		fmt.Sprintf("owner:%s", GitHubConfig.Owner),
	}

	// set discovery method
	var discoveryMethod string
	if GitHubConfig.Cli {
		discoveryMethod = "cli"
	} else {
		discoveryMethod = "api"
	}

	token, present := os.LookupEnv("GITHUB_TOKEN")
	if discoveryMethod == "api" && !present {
		logger.PrintLogMessage(loggingContext, "⛔ required environment variable not set: GITHUB_TOKEN")
		os.Exit(2)
	}

	loggingContext = []string{
		"provider:github",
		fmt.Sprintf("config:%s", GitHubConfig.Name),
		fmt.Sprintf("discovery:%s", discoveryMethod),
		fmt.Sprintf("owner:%s", GitHubConfig.Owner),
	}
	message := "🔍 Finding repos"
	logger.PrintLogMessage(loggingContext, message)

	// find list of repos
	var githubRepos []githubRepo
	if GitHubConfig.Cli {
		githubRepos = findReposCli(GitHubConfig.Owner)
	} else {
		githubRepos = findRepos(GitHubConfig.Owner, token)
	}

	// set clone path
	var destination string
	if GitHubConfig.Destination == "" {
		destination = path.Join("github", GitHubConfig.Owner)
	} else {
		destination = GitHubConfig.Destination
	}

	// apply filters
	if GitHubConfig.Include != "" {
		githubRepos = filterIncludeRepos(GitHubConfig.Include, githubRepos)
	}
	if GitHubConfig.Exclude != "" {
		githubRepos = filterExcludeRepos(GitHubConfig.Include, githubRepos)
	}

	// save results
	for _, githubRepo := range githubRepos {
		var gitRepo = config.GitRepo{
			Context:     loggingContext,
			Name:        githubRepo.Name,
			HttpsUrl:    githubRepo.Url,
			SshUrl:      githubRepo.SSHUrl,
			CloneMethod: GitHubConfig.CloneMethod,
			Destination: destination,
		}
		repos = append(repos, gitRepo)
	}

	loggingContext = []string{
		"provider:github",
		fmt.Sprintf("config:%s", GitHubConfig.Name),
		fmt.Sprintf("discovery:%s", discoveryMethod),
		fmt.Sprintf("owner:%s", GitHubConfig.Owner),
	}
	message = fmt.Sprintf("🔍 Found '%d' repos", len(repos))
	logger.PrintLogMessage(loggingContext, message)

	return

}
