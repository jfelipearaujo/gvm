package helpers

import (
	"os"
	"path/filepath"
	"strings"
)

func updatePathUserEnvironmentVariable(oldValue string, newValue string) error {
	currentValue := os.Getenv("PATH")

	pathValues := strings.Split(currentValue, ";")

	for i, pathValue := range pathValues {
		if pathValue == newValue {
			// Requested value already exists into Path variable
			return nil
		} else if pathValue == oldValue {
			// Get other values from Path variable and ignore the existing one
			pathValues = append(pathValues[:i], pathValues[i+1:]...)
		}
	}

	currentValue = strings.Join(pathValues, ";")
	currentValue = currentValue + ";" + newValue

	err := os.Setenv("PATH", currentValue)

	if err != nil {
		return err
	}

	return nil
}

func GetValueFromVariable(name string) (string, error) {
	return os.Getenv(name)
}

func SetGoRoot(newGoRoot string) error {
	err := os.Setenv("GOROOT", newGoRoot)

	return err
}

func SetGoPath(newGoPath string) error {
	err := os.Setenv("GOPATH", newGoPath)

	return err
}

func UpdatePath(newValue string) error {
	goRoot := os.Getenv("GOROOT")

	err := updatePathUserEnvironmentVariable(filepath.Join(goRoot, "bin", string(os.PathSeparator), ""), newValue)

	return err
}
