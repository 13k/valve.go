package schema

import (
	"fmt"
	"net/http"
	"net/url"
	"regexp"
)

const (
	methodMinVersion int = 1
)

var (
	smNameRegexp       = regexp.MustCompile(`^[[:alpha:]]\w*$`)
	smValidHTTPMethods = map[string]bool{
		http.MethodGet:  true,
		http.MethodPut:  true,
		http.MethodPost: true,
	}
)

// Method holds the specification of an API interface method.
type Method struct {
	Name         string       `json:"name"`
	Version      int          `json:"version"`
	HTTPMethod   string       `json:"httpmethod"`
	Params       MethodParams `json:"parameters"`
	Undocumented bool         `json:"undocumented"`

	key        *MethodKey
	strVersion string
}

// NewMethod creates a new interface `Method`.
//
// Returns `Method.Validate` errors.
func NewMethod(
	name string,
	version int,
	httpMethod string,
	params MethodParams,
	undocumented bool,
) (*Method, error) {
	sm := &Method{
		Name:         name,
		Version:      version,
		HTTPMethod:   httpMethod,
		Params:       params,
		Undocumented: undocumented,
		strVersion:   fmt.Sprintf("v%d", version),
	}

	if err := sm.Validate(); err != nil {
		return nil, err
	}

	return sm, nil
}

func MustNewMethod(
	name string,
	version int,
	httpMethod string,
	params MethodParams,
	undocumented bool,
) *Method {
	sm, err := NewMethod(name, version, httpMethod, params, undocumented)

	if err != nil {
		panic(err)
	}

	return sm
}

// Validate checks if fields are valid.
//
// It should be used after creating or updating a Method (unmarshalling or direct instantiation
// without `NewMethod`).
//
// Returns `*InvalidMethodError` on error, wrapping one of the following errors:
//   - `*InvalidMethodNameError` if the method has an invalid name.
//   - `*InvalidMethodVersionError` if the method has an invalid version.
//   - `*InvalidMethodHTTPMethodError` if the method has an invalid http method.
func (sm *Method) Validate() error {
	if sm.key == nil {
		sm.key = &MethodKey{
			Name:    sm.Name,
			Version: sm.Version,
		}
	}

	if !smNameRegexp.MatchString(sm.Name) {
		return &InvalidMethodError{err: NewInvalidMethodNameError(sm)}
	}

	if sm.Version < methodMinVersion {
		return &InvalidMethodError{err: NewInvalidMethodVersionError(sm)}
	}

	if _, ok := smValidHTTPMethods[sm.HTTPMethod]; !ok {
		return &InvalidMethodError{err: NewInvalidMethodHTTPMethodError(sm)}
	}

	return nil
}

// Key returns the key identifying the method.
func (m *Method) Key() *MethodKey {
	return m.key
}

// FormatVersion returns the formatted version string (in the format "v<version>") to be used as a
// request path parameter.
func (m *Method) FormatVersion() string {
	return m.strVersion
}

// ValidateParams validates the given parameters against the params collection.
//
// Returns errors described in `MethodParams.ValidateParams`.
func (m *Method) ValidateParams(params url.Values) error {
	return m.Params.ValidateParams(params)
}

// MethodKey is the key that uniquely identifies a `Method`.
type MethodKey struct {
	Name    string
	Version int
}

// String formats the key.
func (k *MethodKey) String() string {
	return fmt.Sprintf("(%q, %d)", k.Name, k.Version)
}
