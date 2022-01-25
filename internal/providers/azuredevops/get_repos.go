package azuredevops

import (
	"clone-all-the-repos/internal/config"
	"clone-all-the-repos/internal/logger"
	"fmt"
	"os"
	"path"
)

func GetAzureDevOpsRepos(AzureDevOpsConfig config.AzureDevOpsConfig) (Repos []config.GitRepo) {

	loggingContext := []string{
		"provider:azuredevops",
		fmt.Sprintf("config:%s", AzureDevOpsConfig.Name),
		fmt.Sprintf("org:%s", AzureDevOpsConfig.Organization),
	}

	// set discovery method
	var discoveryMethod string
	if AzureDevOpsConfig.Cli {
		discoveryMethod = "cli"
	} else {
		discoveryMethod = "api"
	}

	// check for environment variable
	var token string
	var present bool
	token, present = os.LookupEnv("AZDO_PERSONAL_ACCESS_TOKEN")
	if discoveryMethod == "api" && !present {
		logger.PrintLogMessage(loggingContext, "⛔ required environment variable not set: AZDO_PERSONAL_ACCESS_TOKEN")
		os.Exit(2)
	}

	// checks for az cli
	// * is present
	// * is logged in

	loggingContext = []string{
		"provider:azuredevops",
		fmt.Sprintf("config:%s", AzureDevOpsConfig.Name),
		fmt.Sprintf("discovery:%s", discoveryMethod),
		fmt.Sprintf("org:%s", AzureDevOpsConfig.Organization),
	}
	message := "🔍 Finding projects..."
	logger.PrintLogMessage(loggingContext, message)

	// find projects
	var projects []project
	if AzureDevOpsConfig.Cli {
		projects = listProjectsCli(AzureDevOpsConfig.Organization)
	} else {
		projects = listProjects(token, AzureDevOpsConfig.Organization)
	}

	// filter projects
	if AzureDevOpsConfig.Include != "" {
		projects = filterIncludeProjects(AzureDevOpsConfig.Include, projects)
	}
	if AzureDevOpsConfig.Exclude != "" {
		projects = filterExcludeProjects(AzureDevOpsConfig.Exclude, projects)
	}

	loggingContext = []string{
		"provider:azuredevops",
		fmt.Sprintf("config:%s", AzureDevOpsConfig.Name),
		fmt.Sprintf("discovery:%s", discoveryMethod),
		fmt.Sprintf("org:%s", AzureDevOpsConfig.Organization),
	}
	message = fmt.Sprintf("🔍 Found '%d' projects", len(projects))
	logger.PrintLogMessage(loggingContext, message)

	// find repos
	for _, project := range projects {

		// list repos in project
		var repos []repo
		if AzureDevOpsConfig.Cli {
			repos = listReposCli(AzureDevOpsConfig.Organization, project.Name)
		} else {
			repos = listRepos(token, AzureDevOpsConfig.Organization, project.Name)
		}

		// apply filters
		if AzureDevOpsConfig.Include != "" {
			repos = filterIncludeRepos(AzureDevOpsConfig.Include, repos)
		}
		if AzureDevOpsConfig.Exclude != "" {
			repos = filterExcludeRepos(AzureDevOpsConfig.Exclude, repos)
		}

		loggingContext = []string{
			"provider:azuredevops",
			fmt.Sprintf("config:%s", AzureDevOpsConfig.Name),
			fmt.Sprintf("discovery:%s", discoveryMethod),
			fmt.Sprintf("org:%s", AzureDevOpsConfig.Organization),
			fmt.Sprintf("project:%s", project.Name),
		}
		message = fmt.Sprintf("🔍 Found '%d' repos", len(repos))
		logger.PrintLogMessage(loggingContext, message)

		// set clone path
		var destination string
		if AzureDevOpsConfig.Destination == "" {
			destination = path.Join("azuredevops", AzureDevOpsConfig.Organization, project.Name)
		} else {
			destination = AzureDevOpsConfig.Destination
		}

		for _, repo := range repos {

			// do not attempt to clone disabled repos
			if repo.IsDisabled {
				continue
			}

			var GitRepo = config.GitRepo{
				Context:     loggingContext,
				Name:        repo.Name,
				HttpsUrl:    repo.RemoteURL,
				SshUrl:      repo.SSHURL,
				CloneMethod: AzureDevOpsConfig.CloneMethod,
				Destination: destination,
			}
			Repos = append(Repos, GitRepo)

		}
	}

	return

}
