package commands

import (
	"log"
	"os"
	"path/filepath"

	"github.com/jfelipearaujo/gvm/helpers"
)

type GoPathCommand struct {
	Folder string `arg:"" optional:"" type:"path" short:"f" help:"A valid directory for the GOPATH"`
}

func (command *GoPathCommand) Run() error {

	if command.Folder == "" {
		log.Println("Using the default directory")

		homeDir, err := os.UserHomeDir()

		if err != nil {
			return err
		}

		command.Folder = filepath.Join(homeDir, "go")
	}

	log.Printf("Setting up GoPath for directory: %s", command.Folder)

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
