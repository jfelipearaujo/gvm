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
	Version string `arg:"" required:"" short:"v" help:"A valid version of Go Lang"`
}

func (command *UseCommand) Run() error {
	log.Printf("Setting go version to %s...\n", command.Version)

	homeDir, err := os.UserHomeDir()

	if err != nil {
		return err
	}

	versions := filepath.Join(homeDir, ".gvm", "versions")

	if _, err := os.Stat(versions); os.IsNotExist(err) {
		return fmt.Errorf("gvm directory does not exist. Please run command 'gvm install %v' and try again", command.Version)
	}

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

	previousGoRootBinPath, err := helpers.GetValueFromVariable("GOROOT")

	if err != nil {
		return err
	}

	previousGoRootBinPath = filepath.Join(previousGoRootBinPath, "bin", "x")

	previousGoRootBinPath = previousGoRootBinPath[:len(previousGoRootBinPath)-1]

	log.Printf("Previous GOROOT: %v\n", previousGoRootBinPath)

	log.Println("Updating environment variables...")

	err = helpers.SetGoRoot(filepath.Join(versions, command.Version, "go"))

	if err != nil {
		return err
	}

	err = helpers.SetGoCurrentVersion(command.Version)

	if err != nil {
		return err
	}

	newGoRootBinPath := filepath.Join(versions, command.Version, "go", "bin", "x")

	newGoRootBinPath = newGoRootBinPath[:len(newGoRootBinPath)-1]

	log.Printf("New GOROOT: %v\n", newGoRootBinPath)

	err = helpers.UpdatePath(previousGoRootBinPath, newGoRootBinPath)

	if err != nil {
		return err
	}

	log.Println("Environment variables updated successfully")

	log.Println("Done! Open a new prompt and type 'go version' to see the changes :)")

	return nil
}
