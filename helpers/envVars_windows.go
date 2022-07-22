package helpers

import (
	"os"
	"strings"

	"golang.org/x/sys/windows/registry"
)

func setUserEnvironmentVariable(name string, value string) error {
	regplace := registry.CURRENT_USER

	key, err := registry.OpenKey(regplace, "Environment", registry.ALL_ACCESS)

	if err != nil {
		return err
	}

	defer key.Close()

	err = key.SetStringValue(name, value)

	return err
}

func updatePathUserEnvironmentVariable(oldValue string, newValue string) error {
	regplace := registry.CURRENT_USER

	key, err := registry.OpenKey(regplace, "Environment", registry.ALL_ACCESS)

	if err != nil {
		return err
	}

	defer key.Close()

	currentValue, _, err := key.GetStringValue("PATH")

	if err != nil {
		return err
	}

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

	err = key.SetStringValue("PATH", currentValue)

	if err != nil {
		return err
	}

	return nil
}

func GetValueFromVariable(name string) (string, error) {
	regplace := registry.CURRENT_USER

	key, err := registry.OpenKey(regplace, "Environment", registry.ALL_ACCESS)

	if err != nil {
		return "", err
	}

	defer key.Close()

	currentValue, _, err := key.GetStringValue(name)

	if err != nil {
		return "", err
	}

	return currentValue, nil
}

func SetGoRoot(newGoRoot string) error {
	err := setUserEnvironmentVariable("GOROOT", newGoRoot)

	return err
}

func SetGoPath(newGoPath string) error {
	err := setUserEnvironmentVariable("GOPATH", newGoPath)

	return err
}

func GetGoCurrentVersion() string {
	return os.Getenv("GVM_CURRENT_GO_VERSION")
}

func SetGoCurrentVersion(newGoVersion string) error {
	err := setUserEnvironmentVariable("GVM_CURRENT_GO_VERSION", newGoVersion)

	return err
}

func UpdatePath(oldValue string, newValue string) error {
	err := updatePathUserEnvironmentVariable(oldValue, newValue)

	return err
}
