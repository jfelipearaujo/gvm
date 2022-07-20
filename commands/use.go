package commands

import (
	"log"
)

type UseCommand struct {
	Version string `arg required short:"v" help:"A valid version of Go Lang"`
}

func (command *UseCommand) Run() error {
	log.Printf("Using go on %s...\n", command.Version)

	return nil
}
