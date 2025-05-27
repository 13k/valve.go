//go:build linux

package steamlib

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

const (
	steamFlatpakAppID = "com.valvesoftware.Steam"
)

//nolint:gochecknoglobals
var (
	steamDirSuffixes = []string{
		".steam/steam",
		".steam/root",
		".local/share/Steam",
	}
)

func findSteam() (*Steam, error) {
	home, err := userHome()

	if err != nil {
		return nil, err
	}

	snapDir, err := snapDir()

	if err != nil {
		return nil, err
	}

	prefixes := []string{
		// Standalone
		home,
		// Flatpak
		filepath.Join(home, ".var/app", steamFlatpakAppID),
		// Snap
		filepath.Join(snapDir, "steam/common"),
	}

	for _, prefix := range prefixes {
		for _, suffix := range steamDirSuffixes {
			path := filepath.Join(prefix, suffix)
			steam, err := trySteamDir(path)

			if err != nil {
				return nil, err
			}

			if steam != nil {
				return steam, nil
			}
		}
	}

	return nil, ErrSteamNotFound
}

func trySteamDir(path string) (*Steam, error) {
	fi, err := os.Stat(path)

	if err != nil {
		if errors.Is(err, fs.ErrNotExist) {
			return nil, nil //nolint:nilnil
		}

		return nil, fmt.Errorf("failed to access possible Steam path %q: %w", path, err)
	}

	if !fi.IsDir() {
		return nil, fmt.Errorf("%w: %q", ErrInvalidSteamInstallDir, path)
	}

	realPath, err := filepath.EvalSymlinks(path)

	if err != nil {
		return nil, fmt.Errorf("failed to normalize Steam path %q: %w", path, err)
	}

	steam := &Steam{Path: realPath}

	return steam, nil
}
