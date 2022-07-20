package commands

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/jfelipearaujo/gvm/helpers"
)

type UseCommand struct {
	Version string `arg required short:"v" help:"A valid version of Go Lang"`
}

func (command *UseCommand) Run() error {
	log.Printf("Setting go version to %s...\n", command.Version)

	homeDir, err := os.UserHomeDir()

	if err != nil {
		return err
	}

	versions := filepath.Join(homeDir, ".gvm", "versions")

	dirs, err := ioutil.ReadDir(versions)

	if err != nil {
		return err
	}

	requestVersionExists := false

	for _, dir := range dirs {
		if dir.IsDir() {
			if command.Version == dir.Name() {
				requestVersionExists = true
				break
			}
		}
	}

	if !requestVersionExists {
		return fmt.Errorf("version %v does not exist, please run the install command and try again", command.Version)
	}

	log.Println("Updating environment variables...")

	err = helpers.SetGoRoot(filepath.Join(versions, command.Version, "go", string(os.PathSeparator), ""))

	if err != nil {
		return err
	}

	err = helpers.SetGoCurrentVersion(command.Version)

	if err != nil {
		return err
	}

	err = helpers.UpdatePath(filepath.Join(versions, command.Version, "go", "bin", string(os.PathSeparator), ""))

	if err != nil {
		return err
	}

	log.Println("Environment variables updated successfully")

	log.Println("Done! Open a new prompt and type 'go version' to see the changes :)")

	return nil
}
