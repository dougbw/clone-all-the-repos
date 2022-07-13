package azuredevops

import (
	"clone-all-the-repos/internal/config"
	"clone-all-the-repos/internal/logger"
	"log"
	"os"
	"os/exec"
)

func Startup(config config.Root) {
	// perform initial checks such as:
	// - is required env var set
	// - is az cli binary present
	// - is az cli logged in

	logger.Context = []string{
		"startup",
		"azuredevops",
	}

	var azCliRequired bool
	var azTokenRequired bool

	for _, azuredevops := range config.AzureDevOps {
		if azuredevops.Cli {
			azCliRequired = true
		} else {
			azTokenRequired = true
		}
	}

	if azCliRequired {

		// check az is present
		_, err := exec.LookPath("az")
		if err != nil {
			log.Fatal(err)
		}
		logger.Print("✅ az cli present")

		// check az is logged in - todo: work out how to actually check this properly
		// output, err := execwrapper.Exec("az", "account", "show")
		// if err != nil {
		// 	logger.PrintErr(context, err.Error())
		// }
		// logger.Print( "✅ logged into az cli")

	}

	// check if env var is set
	if azTokenRequired {
		var present bool
		_, present = os.LookupEnv("AZDO_PERSONAL_ACCESS_TOKEN")
		if !present {
			logger.Print("⛔ required environment variable not set: AZDO_PERSONAL_ACCESS_TOKEN")
			os.Exit(2)
		}

		logger.Print("✅ AZDO_PERSONAL_ACCESS_TOKEN is set")

	}

	/*
		todo:
		- check token is valid
		- test token has correct scope/permissions
	*/

}
