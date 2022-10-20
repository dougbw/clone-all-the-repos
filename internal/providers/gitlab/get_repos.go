package gitlab

import (
	"clone-all-the-repos/internal/config"
	"clone-all-the-repos/internal/logger"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"
)

// todo:
// * pagination
// * filtering

func GetRepos(GitLabConfig config.GitLabConfig) (Repos []config.GitRepo) {

	logger.Context = []string{
		"provider:gitlab",
		fmt.Sprintf("config:%s", GitLabConfig.Name),
		fmt.Sprintf("user:%s", GitLabConfig.User),
	}
	logger.Print("🔍 Finding projects")

	token, present := os.LookupEnv("GITLAB_TOKEN")
	if !present {
		logger.Print("⛔ required environment variable not set: GITLAB_TOKEN")
		os.Exit(2)
	}

	bearer := fmt.Sprintf("Bearer %s", token)

	// get all projects
	projects := getProjects(bearer)

	// apply filters
	beforeFilterCount := len(projects)
	if GitLabConfig.Include != "" {
		projects = filterIncludeProjects(GitLabConfig.Include, projects)
	}
	if GitLabConfig.Exclude != "" {
		projects = filterExcludeProjects(GitLabConfig.Exclude, projects)
	}
	afterFilterCount := len(projects)

	diff := beforeFilterCount - afterFilterCount
	fmt.Printf("filters removed '%d' projects", diff)

	// create the list of repos to return
	for _, project := range projects {

		// is the gitlab user required to clone using https?
		// how will this work for pulling repos owner by other users?
		var httpsUrl string
		if GitLabConfig.CloneMethod == "https" {
			authHttps := fmt.Sprintf("https://%s:%s@", GitLabConfig.User, token)
			httpsUrl = strings.Replace(project.HTTPURLToRepo, "https://", authHttps, -1)
		} else {
			httpsUrl = project.HTTPURLToRepo
		}

		// set destination path
		var destination string
		if GitLabConfig.Destination == "" {
			fullPath := path.Join("gitlab", project.PathWithNamespace)
			destination, _ = filepath.Split(fullPath)
		} else {
			destination = GitLabConfig.Destination
		}

		var GitRepo = config.GitRepo{
			Context:     logger.Context,
			Name:        project.Name,
			HttpsUrl:    httpsUrl,
			SshUrl:      project.SSHURLToRepo,
			CloneMethod: GitLabConfig.CloneMethod,
			Destination: destination,
		}
		Repos = append(Repos, GitRepo)
	}

	return
}
