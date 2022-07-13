package command

import (
	"clone-all-the-repos/internal/config"
	"clone-all-the-repos/internal/logger"
	"clone-all-the-repos/internal/providers/azuredevops"
	"clone-all-the-repos/internal/providers/github"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

type Startup struct {
	ConfigPath struct {
		Absolute string `json:"absolutePath"`
		File     string `json:"file"`
		Dir      string `json:"dir"`
	} `json:"config"`
	Config config.Root
}

func StartupCommand() (startup Startup) {

	// parse cli args. If this needs to be extended then use https://github.com/spf13/cobra
	if len(os.Args[1:]) != 1 {
		fmt.Println("usage:\nclone-all-the-repos <configfile>")
		os.Exit(1)
	}

	logger.Context = []string{
		"startup",
	}

	// check if git is present
	_, err := exec.LookPath("git")
	if err != nil {
		log.Fatal(err)
		logger.PrintErr("❌ git not detected")
	}
	logger.Print("✅ git detected")

	// setup config paths
	configPath := os.Args[1]
	absolute, _ := filepath.Abs(configPath)
	dir, file := filepath.Split(absolute)
	startup.ConfigPath.Absolute = absolute
	startup.ConfigPath.Dir = dir
	startup.ConfigPath.File = file

	// load config file
	startup.Config = config.Open(absolute)

	// azuredevops startup checks
	azuredevops.Startup(startup.Config)

	// github startup checks
	github.Startup(startup.Config)

	// todo: gitlab startup checks
	// gitlab.Startup(startup.Config)

	return startup

}
