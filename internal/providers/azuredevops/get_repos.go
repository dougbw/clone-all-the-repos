package azuredevops

import (
	"clone-all-the-repos/internal/config"
	"clone-all-the-repos/internal/logger"
	"os"
	"path"
)

func GetRepos(AzureDevOpsConfig config.AzureDevOpsConfig) (Repos []config.GitRepo) {

	// set discovery method
	var discoveryMethod string
	if AzureDevOpsConfig.Cli {
		discoveryMethod = "cli"
	} else {
		discoveryMethod = "api"
	}

	logger.Context = []string{
		"discovery",
		"azuredevops",
		AzureDevOpsConfig.Name,
		discoveryMethod,
		AzureDevOpsConfig.Organization,
	}
	logger.Print("🔍 Finding projects...")

	token, _ := os.LookupEnv("AZDO_PERSONAL_ACCESS_TOKEN")

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

	logger.Printf("🔍 Found '%d' projects", len(projects))

	// find repos
	for _, project := range projects {

		logger.Context = []string{
			"discovery",
			"azuredevops",
			AzureDevOpsConfig.Name,
			discoveryMethod,
			AzureDevOpsConfig.Organization,
			project.Name,
		}

		logger.Print("🔍 Finding repos...")

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

		logger.Printf("🔍 Found '%d' repos", len(repos))

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
				Context:     logger.Context[1:],
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
