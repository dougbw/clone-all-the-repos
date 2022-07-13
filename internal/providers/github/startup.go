package github

import (
	"clone-all-the-repos/internal/config"
	"clone-all-the-repos/internal/execwrapper"
	"clone-all-the-repos/internal/logger"
	"log"
	"os"
	"os/exec"
)

func Startup(config config.Root) {

	logger.Context = []string{
		"startup",
		"github",
	}

	var ghCliRequired bool
	var ghTokenRequired bool

	for _, githubConfig := range config.GitHub {
		if githubConfig.Cli {
			ghCliRequired = true
		} else {
			ghTokenRequired = true
		}
	}

	if ghCliRequired {

		// check gh is present
		_, err := exec.LookPath("gh")
		if err != nil {
			log.Fatal(err)
		}
		logger.Print("✅ gh detected")

		// check gh is logged in
		_, err = execwrapper.Exec("gh", "auth", "status")
		if err != nil {
			logger.PrintErr("⛔ gh is not logged in - please run 'gh auth login' to authenticate")
		}
		logger.Print("✅ gh is logged in")

	}

	// check env var is set
	if ghTokenRequired {
		var present bool
		_, present = os.LookupEnv("GITHUB_TOKEN")
		if !present {
			logger.Print("⛔ required environment variable not set: GITHUB_TOKEN")
			os.Exit(2)
		}
		logger.Print("✅ GITHUB_TOKEN present")

	}

}
