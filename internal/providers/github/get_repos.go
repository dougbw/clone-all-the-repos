package github

import (
	"clone-all-the-repos/internal/config"
	"clone-all-the-repos/internal/logger"
	"os"
	"path"
)

func GetRepos(GitHubConfig config.GitHubConfig) (repos []config.GitRepo) {

	token, _ := os.LookupEnv("GITHUB_TOKEN")

	// set discovery method
	var discoveryMethod string
	if GitHubConfig.Cli {
		discoveryMethod = "cli"
	} else {
		discoveryMethod = "api"
	}

	logger.Context = []string{
		"discovery",
		"github",
		GitHubConfig.Name,
		discoveryMethod,
		GitHubConfig.Owner,
	}

	logger.Print("🔍 Finding repos...")

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
			Context:     logger.Context[1:],
			Name:        githubRepo.Name,
			HttpsUrl:    githubRepo.Url,
			SshUrl:      githubRepo.SSHUrl,
			CloneMethod: GitHubConfig.CloneMethod,
			Destination: destination,
		}
		repos = append(repos, gitRepo)
	}

	logger.Printf("🔍 Found '%d' repos", len(repos))

	return

}
