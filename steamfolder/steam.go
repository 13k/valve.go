package steamfolder

import (
	"errors"
)

var (
	ErrSteamNotFound          = errors.New("Steam not found")
	ErrInvalidSteamInstallDir = errors.New("Steam install path exists but is not a directory")
)

// Steam represents a Steam installation.
type Steam struct {
	// Install directory
	Path string
}

// NewSteam creates a new Steam instance with the given `path`.
func NewSteam(path string) *Steam {
	return &Steam{Path: path}
}

// FindSteam tries to find an Steam installation directory.
//
// It returns error `ErrSteamNotFound` if it did not find a suitable directory.
func FindSteam() (*Steam, error) {
	return findSteam()
}

// Library parses the Steam library.
func (s *Steam) Library() (*Library, error) {
	return NewLibrary(s.Path)
}
