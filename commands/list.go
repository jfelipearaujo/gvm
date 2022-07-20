package commands

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

type ListCommand struct {
}

func (command *ListCommand) Run() error {
	log.Println("Listing installed versions...")

	goCurrentVersion := os.Getenv("GVM_CURRENT_GO_VERSION")

	homeDir, err := os.UserHomeDir()

	if err != nil {
		return err
	}

	versions := filepath.Join(homeDir, ".gvm", "versions")

	dirs, err := ioutil.ReadDir(versions)

	if err != nil {
		return err
	}

	for _, dir := range dirs {
		if dir.IsDir() {
			if goCurrentVersion == dir.Name() {
				fmt.Printf("- %s\tCURRENT VERSION\n", dir.Name())
			} else {
				fmt.Printf("- %s\n", dir.Name())
			}
		}
	}

	return nil
}
