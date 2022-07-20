package commands

import (
	"log"
)

type ListCommand struct {
}

func (command *ListCommand) Run() error {
	log.Println("Listing versions...")

	return nil
}
