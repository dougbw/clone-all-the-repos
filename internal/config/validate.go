package config

import (
	"clone-all-the-repos/internal/logger"

	"github.com/go-playground/validator/v10"
)

func Validate(config Root) {

	// validate root
	validate := validator.New()
	if err := validate.Struct(&config); err != nil {
		logger.Context = []string{
			"startup",
			"config",
			"validation",
		}
		logger.PrintErrf("Missing required attributes %v\n", err)
	}

	// validate azuredevops
	for _, azuredevops := range config.AzureDevOps {
		if err := validate.Struct(&azuredevops); err != nil {
			logger.Context = []string{
				"startup",
				"config",
				// "azuredevops",
				// azuredevops.Name,
			}
			logger.PrintErrf("invalid configuration at azuredevops.%s - %v\n", azuredevops.Name, err)
		}
	}

	// validate github
	for _, github := range config.GitHub {
		if err := validate.Struct(&github); err != nil {
			logger.Context = []string{
				"startup",
				"config",
				"validation",
				// "github",
				// github.Name,
			}
			logger.PrintErrf("invalid configuration at github.%s - %v\n", github.Name, err)
		}
	}

	// validate gitlab
	for _, gitlab := range config.GitLab {
		if err := validate.Struct(&gitlab); err != nil {
			logger.Context = []string{
				"startup",
				"config",
				"validation",
				// "gitlab",
				// gitlab.Name,
			}
			logger.PrintErrf("invalid configuration at gitlab.%s - %v\n", gitlab.Name, err)
		}
	}

	logger.Print("✅ configuration file is valid")

}
