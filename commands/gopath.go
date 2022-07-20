package commands

import (
	"log"
)

type GoPathCommand struct {
	Folder string `arg required short:"f" help:"A valid directory for the GOPATH"`
}

func (command *GoPathCommand) Run() error {
	log.Printf("Uninstalling go on %s...\n", command.Folder)

	return nil
}
