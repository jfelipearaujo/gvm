package commands

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/jfelipearaujo/gvm/helpers"
)

type ListCommand struct {
}

func (command *ListCommand) Run() error {
	goCurrentVersion := helpers.GetGoCurrentVersion()

	log.Printf("Current version: %s\n", goCurrentVersion)

	homeDir, err := os.UserHomeDir()

	if err != nil {
		return err
	}

	versions := filepath.Join(homeDir, ".gvm", "versions")

	dirs, err := ioutil.ReadDir(versions)

	if err != nil {
		return err
	}

	log.Println("Installed versions:")

	for _, dir := range dirs {
		if dir.IsDir() {
			fmt.Printf("- %s\n", dir.Name())
		}
	}

	return nil
}
