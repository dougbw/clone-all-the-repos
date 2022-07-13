package command

import (
	"clone-all-the-repos/internal/config"
	"clone-all-the-repos/internal/git"
	"clone-all-the-repos/internal/logger"
)

func CloneCommand(startup Startup, repos []config.GitRepo) {

	logger.Context = []string{
		"clone",
	}

	if len(repos) == 0 {
		logger.Print("no repos to clone")
		return
	}

	logger.Printf("⌛ Cloning '%d' repos", len(repos))

	baseDirectory := startup.ConfigPath.Dir

	for _, Repo := range repos {
		git.CloneGitRepo(baseDirectory, Repo)
	}

	logger.Context = []string{
		"clone",
	}
	logger.Print("Done")

}
