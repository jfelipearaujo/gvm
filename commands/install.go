package commands

import (
	"log"
)

type InstallCommand struct {
	Version string `arg required short:"v" help:"A valid version of Go Lang"`
}

func (command *InstallCommand) Run() error {
	log.Printf("Installing go on %s...\n", command.Version)

	return nil
}
