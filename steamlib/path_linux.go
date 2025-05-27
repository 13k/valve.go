package steamlib

import (
	"os"
	"path/filepath"
)

const (
	snapDataEnvVar = "SNAP_USER_DATA"
)

func snapDir() (string, error) {
	home, err := userHome()

	if err != nil {
		return "", err
	}

	path := os.Getenv(snapDataEnvVar)

	if path == "" {
		path = filepath.Join(home, "snap")
	}

	return path, nil
}
