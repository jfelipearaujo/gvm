package commands

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/jfelipearaujo/gvm/helpers"
)

type UninstallCommand struct {
	Version string `arg:"" required:"" short:"v" help:"An installed version of Go Lang"`
}

func (command *UninstallCommand) Run() error {
	log.Printf("Uninstalling go version %s...\n", command.Version)

	goCurrentVersion, err := helpers.GetValueFromVariable("GVM_CURRENT_GO_VERSION")

	if err != nil {
		return err
	}
	if goCurrentVersion == command.Version {
		return fmt.Errorf("you can't uninstall your current version")
	}

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
		return fmt.Errorf("version %v does not exist", command.Version)
	}

	destination := filepath.Join(homeDir, ".gvm", "versions", command.Version)

	err = os.RemoveAll(destination)

	if err != nil {
		return err
	}

	log.Println("Version uninstalled successfully")

	return nil
}
