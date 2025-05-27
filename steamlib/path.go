package steamlib

import (
	"fmt"
	"os"
)

func userHome() (string, error) {
	path, err := os.UserHomeDir()

	if err != nil {
		return "", fmt.Errorf("failed to determine user home: %w", err)
	}

	return path, nil
}
