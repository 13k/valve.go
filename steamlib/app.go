package steamlib

import (
	"fmt"
	"path/filepath"
	"strconv"
)

//nolint:gochecknoglobals
var (
	appManifestRelPathFmt      = filepath.Join("steamapps", "appmanifest_%d.acf")
	appRelInstallDirFromFolder = "common"
)

// AppManifest represents a parsed `appmanifest_*.acf` file inside a library folder.
type AppManifest struct {
	// ID
	AppID AppID
	// Name
	Name string
	// Install dir
	InstallDir string
	// Absolute install dir
	AbsInstallDir string
	// Size (in bytes)
	Size uint64
}

// NewAppManifest parses an AppManifest in a given `folder` for the given `appID`.
func NewAppManifest(folder *LibraryFolder, appID AppID) (*AppManifest, error) {
	filename := fmt.Sprintf(appManifestRelPathFmt, appID)
	path := filepath.Join(folder.Path, filename)

	return ParseAppManifest(path)
}

// ParseAppManifest parses an app manifest at the given `path`.
func ParseAppManifest(path string) (*AppManifest, error) {
	kvApp, err := kv1UnmarshalFile[appManifestKv](path)

	if err != nil {
		return nil, fmt.Errorf("failed to parse AppManifest: %w", err)
	}

	return appManifestFromKv(path, kvApp)
}

type appManifestKv struct {
	State appStateKv `mapstructure:"AppState"`
}

type appStateKv struct {
	AppID      string `mapstructure:"appid"`
	Name       string `mapstructure:"name"`
	InstallDir string `mapstructure:"installdir"`
	Size       string `mapstructure:"SizeOnDisk"`
}

func appManifestFromKv(path string, kv *appManifestKv) (*AppManifest, error) {
	appID, err := strconv.ParseUint(kv.State.AppID, 10, 64)

	if err != nil {
		return nil, fmt.Errorf("failed to parse AppManifest field `AppID` (%q): %w", kv.State.AppID, err)
	}

	size, err := strconv.ParseUint(kv.State.Size, 10, 64)

	if err != nil {
		return nil, fmt.Errorf("failed to parse AppManifest field `Size` (%q): %w", kv.State.Size, err)
	}

	folderDir := filepath.Dir(path)
	absInstallDir := filepath.Join(folderDir, appRelInstallDirFromFolder, kv.State.InstallDir)

	manifest := &AppManifest{
		AppID:         appID,
		Name:          kv.State.Name,
		InstallDir:    kv.State.InstallDir,
		AbsInstallDir: absInstallDir,
		Size:          size,
	}

	return manifest, nil
}
