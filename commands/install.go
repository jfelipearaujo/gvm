package commands

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/jfelipearaujo/gvm/helpers"
)

type InstallCommand struct {
	Version string `arg:"" required:"" short:"v" help:"A valid version of Go Lang"`
}

func (command *InstallCommand) Run() error {
	log.Println("Installing golang...")

	osArch := strings.ToLower(os.Getenv("PROCESSOR_ARCHITECTURE"))

	if osArch == "x86" {
		osArch = "386"
	}

	zipFileDir, err := helpers.DownloadGoLang(command.Version, "windows", osArch)

	if err != nil {
		return err
	}

	homeDir, err := os.UserHomeDir()

	if err != nil {
		return err
	}

	destination := filepath.Join(homeDir, ".gvm", "versions", command.Version)

	_, err = os.Stat(destination)

	if os.IsNotExist(err) {
		err = os.MkdirAll(destination, os.ModePerm)

		if err != nil {
			return err
		}
	} else {
		log.Println("The requested version it's already installed, so will be removed and replaced...")

		err = os.RemoveAll(destination)

		if err != nil {
			return err
		}

		log.Println("Previous files were removed successfully")

		err = os.MkdirAll(destination, os.ModePerm)

		if err != nil {
			return err
		}
	}

	err = helpers.UnzipSource(zipFileDir, destination)

	if err != nil {
		return err
	}

	err = os.Remove(zipFileDir)

	if err != nil {
		return err
	}

	log.Println("Updating environment variables...")

	err = helpers.SetGoRoot(filepath.Join(destination, "go"))

	if err != nil {
		return err
	}

	err = helpers.SetGoCurrentVersion(command.Version)

	if err != nil {
		return err
	}

	err = helpers.UpdatePath(filepath.Join(destination, "go", "bin"))

	if err != nil {
		return err
	}

	log.Println("Environment variables updated successfully")

	log.Println("Installation completed successfully")

	return nil
}
