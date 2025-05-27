package steamlib

import (
	"fmt"
	"path/filepath"
)

//nolint:gochecknoglobals
var (
	libraryRelPath = filepath.Join("steamapps", "libraryfolders.vdf")
)

// Library represents a parsed `libraryfolders.vdf` file.
type Library struct {
	// List of library folders
	Folders []*LibraryFolder

	apps map[AppID]*LibraryFolder
}

// NewLibrary parses a Library found in a given `steamDir` directory.
func NewLibrary(steamDir string) (*Library, error) {
	return ParseLibrary(filepath.Join(steamDir, libraryRelPath))
}

// ParseLibrary parses a `libraryfolders.vdf` file at the given `path`.
func ParseLibrary(path string) (*Library, error) {
	kvLib, err := kv1UnmarshalFile[libraryKv](path)

	if err != nil {
		return nil, fmt.Errorf("failed to parse Library: %w", err)
	}

	return libraryFromKv(kvLib)
}

// FindApp finds and parses the app manifest for the given `appID`, returning the folder in which
// the app was found and the parsed manifest.
//
// It returns `(nil, nil, nil)` if the given `appID` was not found.
func (l *Library) FindApp(appID AppID) (*LibraryFolder, *AppManifest, error) {
	folder, ok := l.apps[appID]

	if !ok {
		return nil, nil, nil
	}

	app, err := folder.ParseAppManifest(appID)

	if err != nil {
		return nil, nil, fmt.Errorf("failed to find app %d: %w", appID, err)
	}

	return folder, app, nil
}

type libraryKv struct {
	Folders map[string]libraryFolderKv `mapstructure:"libraryfolders"`
}

func libraryFromKv(kv *libraryKv) (*Library, error) {
	kvFolders, err := kv1Array(kv.Folders)

	if err != nil {
		return nil, fmt.Errorf("failed to normalize Library: %w", err)
	}

	folders := make([]*LibraryFolder, len(kvFolders))
	apps := make(map[AppID]*LibraryFolder)

	for i, kvFolder := range kvFolders {
		folder, err := libraryFolderFromKv(&kvFolder)

		if err != nil {
			return nil, err
		}

		folders[i] = folder

		for _, appID := range folder.Apps {
			apps[appID] = folder
		}
	}

	lib := &Library{
		Folders: folders,
		apps:    apps,
	}

	return lib, nil
}
