package helpers

import (
	"fmt"
	"io"
	"io/fs"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

const GoBaseAddress string = "https://golang.org/dl/"

func DownloadGoLang(version string, osName string, arch string) (string, error) {
	log.Println("Downloading file...")

	fileName := fmt.Sprintf("go%s.%s-%s.zip", version, osName, arch)

	url := fmt.Sprintf("%s%s", GoBaseAddress, fileName)

	homeDir, err := os.UserHomeDir()

	if err != nil {
		return "", err
	}

	err = os.MkdirAll(filepath.Join(homeDir, ".gvm", "downloads"), fs.ModePerm)

	if err != nil {
		return "", err
	}

	zipFileDir := filepath.Join(homeDir, ".gvm", "downloads", fileName)

	_, err = os.Stat(zipFileDir)

	if !os.IsNotExist(err) {
		err = os.Remove(zipFileDir)

		if err != nil {
			return "", err
		}
	}

	file, err := os.Create(zipFileDir)

	if err != nil {
		return "", err
	}

	defer file.Close()

	response, err := http.Get(url)

	if err != nil {
		return "", err
	}

	defer response.Body.Close()

	_, err = io.Copy(file, response.Body)

	if err != nil {
		return "", err
	}

	if response.StatusCode != 200 {
		defer os.Remove(zipFileDir)

		return "", fmt.Errorf("download failed with status code %v", response.StatusCode)
	}

	log.Println("Download completed successfully")

	return zipFileDir, nil
}
