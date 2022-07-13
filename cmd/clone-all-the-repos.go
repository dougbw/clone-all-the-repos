package main

import (
	"clone-all-the-repos/internal/command"
)

func main() {

	// startup
	startup := command.StartupCommand()

	// discovery
	discovery := command.DiscoveryCommand(startup)

	// clone
	command.CloneCommand(startup, discovery)

}
