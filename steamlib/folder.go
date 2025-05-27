package steamlib

import (
	"fmt"
	"slices"
	"strconv"
)

// LibraryFolder represents a library folder entry in a `libraryfolders.vdf` file.
type LibraryFolder struct {
	// Absolute path
	Path string
	// Label
	Label string
	// Total size (in bytes)
	TotalSize uint64
	// List of app IDs (sorted)
	Apps []AppID
}

// ContainsApp returns true if this LibraryFolder contains the given `appID`.
func (f *LibraryFolder) ContainsApp(appID AppID) bool {
	_, contains := slices.BinarySearch(f.Apps, appID)

	return contains
}

// ParseAppManifest parses the AppManifest for the given `appID`.
func (f *LibraryFolder) ParseAppManifest(appID AppID) (*AppManifest, error) {
	return NewAppManifest(f, appID)
}

type libraryFolderKv struct {
	Path      string            `mapstructure:"path"`
	Label     string            `mapstructure:"label"`
	TotalSize string            `mapstructure:"totalsize"`
	Apps      map[string]string `mapstructure:"apps"`
}

func libraryFolderFromKv(kv *libraryFolderKv) (*LibraryFolder, error) {
	totalSize, err := strconv.ParseUint(kv.TotalSize, 10, 64)

	if err != nil {
		return nil, fmt.Errorf("failed to parse LibraryFolder field `TotalSize` (%q): %w", kv.TotalSize, err)
	}

	apps := make([]AppID, 0, len(kv.Apps))

	for s := range kv.Apps {
		appID, err := strconv.ParseUint(s, 10, 64)

		if err != nil {
			return nil, fmt.Errorf("failed to parse LibraryFolder field `Apps` (%q): %w", s, err)
		}

		apps = append(apps, appID)
	}

	slices.Sort(apps)

	folder := &LibraryFolder{
		Path:      kv.Path,
		Label:     kv.Label,
		TotalSize: totalSize,
		Apps:      apps,
	}

	return folder, nil
}
