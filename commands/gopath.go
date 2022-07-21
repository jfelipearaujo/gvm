package commands

import (
	"log"
	"os"
	"path/filepath"

	"github.com/jfelipearaujo/gvm/helpers"
)

type GoPathCommand struct {
	Folder string `arg short:"f" help:"A valid directory for the GOPATH"`
}

func (command *GoPathCommand) Run() error {
	log.Println("Setting up GoPath...")

	if command.Folder == "" {
		homeDir, err := os.UserHomeDir()

		if err != nil {
			return err
		}

		command.Folder = filepath.Join(homeDir, "go")
	}

	// If the main folder doesn't exist, create it
	if _, err := os.Stat(command.Folder); os.IsNotExist(err) {
		err = os.MkdirAll(command.Folder, os.ModePerm)

		if err != nil {
			return err
		}
	}

	folders := []string{
		"bin",
		"pkg",
		"src",
	}

	for _, folder := range folders {
		subFolder := filepath.Join(command.Folder, folder)

		// If the sub folder doesn't exist, create it
		if _, err := os.Stat(subFolder); os.IsNotExist(err) {
			err = os.MkdirAll(subFolder, os.ModePerm)

			if err != nil {
				return err
			}
		}
	}

	err := helpers.SetGoPath(command.Folder)

	if err != nil {
		return err
	}

	log.Println("GoPath setup completed")

	return nil
}
