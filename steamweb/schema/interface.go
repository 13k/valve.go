package schema

import (
	"fmt"
	"regexp"

	"github.com/13k/valve.go/steamlib"
)

var (
	siNameRegexp      = regexp.MustCompile(`^[A-Z]\w*$`)
	siNameAppIDRegexp = regexp.MustCompile(`^(\w+)_(\d+)$`)
)

// Interface holds the specification of an API interface.
type Interface struct {
	Name    string  `json:"name"`
	Methods Methods `json:"methods"`

	key *InterfaceKey
}

// NewInterface creates a new `Interface`.
//
// Returns `Interface.Validate` errors.
func NewInterface(name string, methods Methods) (*Interface, error) {
	si := &Interface{
		Name:    name,
		Methods: methods,
	}

	if err := si.Validate(); err != nil {
		return nil, err
	}

	return si, nil
}

// MustNewInterface is like `NewInterface` but panics on errors.
func MustNewInterface(name string, methods Methods) *Interface {
	si, err := NewInterface(name, methods)

	if err != nil {
		panic(err)
	}

	return si
}

func (si *Interface) parseKey() error {
	var err error

	si.key, err = parseInterfaceKey(si.Name)

	if err != nil {
		return err
	}

	return nil
}

// Validate checks if fields are valid.
//
// It should be used after creating or updating an `Interface` (unmarshalling or direct
// instantiation without `NewInterface`).
//
// Returns `*InvalidInterfaceError` on error, wrapping one of the errors:
//   - `*InvalidInterfaceNameError` if the interface has an invalid name.
//   - `*InvalidMethodError` if any of the interface methods are invalid.
func (si *Interface) Validate() error {
	if err := si.parseKey(); err != nil {
		return &InvalidInterfaceError{err: err}
	}

	if err := si.Methods.Validate(); err != nil {
		return &InvalidInterfaceError{err: err}
	}

	return nil
}

// Key returns the key identifying the interface.
//
// It panics if key was not parsed (`Interface` was created without `NewInterface` and `Validate`
// was not called).
func (si *Interface) Key() *InterfaceKey {
	if si.key == nil {
		panic(fmt.Sprintf("steamweb/schema.Interface (name=%q): key is nil", si.Name))
	}

	return si.key
}

// InterfaceKey is the key that uniquely identifies a `Interface`.
//
// Interface names are formed by a base name and an optional AppID part, in the format
// `<BaseName>_<AppID>`. Interfaces can be non app-specific and omit the "_<AppID>" part, in these
// cases Key returns a key with AppID = 0.
//
// For example in "IDOTAMatch_570", the base name is "IDOTAMatch" and AppID is 570. In "ISteamApps",
// the base name is "ISteamApps" and AppID is 0.
type InterfaceKey struct {
	Name  string
	AppID steamlib.AppID
}

func parseInterfaceKey(name string) (*InterfaceKey, error) {
	if !siNameRegexp.MatchString(name) {
		return nil, &InvalidInterfaceNameError{Name: name}
	}

	var (
		appID steamlib.AppID
		err   error
	)

	baseName := name

	if matches := siNameAppIDRegexp.FindStringSubmatch(name); matches != nil {
		appID, err = steamlib.ParseAppID(matches[2])

		if err != nil {
			return nil, &InvalidInterfaceNameError{Name: name, err: err}
		}

		baseName = matches[1]
	}

	key := &InterfaceKey{
		Name:  baseName,
		AppID: appID,
	}

	return key, nil
}

// String formats the key.
func (k *InterfaceKey) String() string {
	return fmt.Sprintf("(%q, %d)", k.Name, k.AppID)
}
