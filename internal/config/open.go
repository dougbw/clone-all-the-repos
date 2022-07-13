package config

import (
	"clone-all-the-repos/internal/logger"
	"os"
	"path/filepath"
)

func Open(configPath string) (config Root) {

	logger.Context = []string{
		"startup",
		"config",
	}
	logger.Print("🔍 Loading configuration")

	absolute, _ := filepath.Abs(configPath)
	_, file := filepath.Split(absolute)

	// test path exists
	fileInfo, err := os.Stat(configPath)
	if err != nil {
		logger.PrintErrf("❌ configuration path not found: %s", absolute)
	}

	// test path is a file
	if fileInfo.IsDir() {
		logger.PrintErrf("❌ configuration path is not a file: %s", absolute)
	}

	logger.Printf("✅ configuration file exists: %s", file)

	// parse
	config, err = Parse(configPath)
	if err != nil {
		logger.PrintErrf("❌ failed to parse configuration file: %s", err)
	}

	// validate
	Validate(config)

	return

}
