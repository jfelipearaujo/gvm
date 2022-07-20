package commands

import (
	"log"
)

type UninstallCommand struct {
	Version string `arg required short:"v" help:"An installed version of Go Lang"`
}

func (command *UninstallCommand) Run() error {
	log.Printf("Uninstalling go on %s...\n", command.Version)

	return nil
}
