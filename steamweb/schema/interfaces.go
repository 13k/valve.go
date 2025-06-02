package schema

import (
	"maps"

	"github.com/13k/valve.go/steamlib"
)

// Interfaces is a collection of `Interface`s.
type Interfaces []*Interface

// NewInterfaces creates a new `Interface` collection.
//
// Returns `Interfaces.Validate` errors.
func NewInterfaces(interfaces ...*Interface) (Interfaces, error) {
	c := Interfaces(interfaces)

	if err := c.Validate(); err != nil {
		return nil, err
	}

	return c, nil
}

// MustNewInterfaces is like `NewInterfaces` but panics if it returned an error.
func MustNewInterfaces(interfaces ...*Interface) Interfaces {
	c, err := NewInterfaces(interfaces...)

	if err != nil {
		panic(err)
	}

	return c
}

// Validate checks whether all interfaces in the collection are valid.
//
// Returns `Interface.Validate` errors.
func (c Interfaces) Validate() error {
	for _, si := range c {
		if err := si.Validate(); err != nil {
			return err
		}
	}

	return nil
}

// Get returns the interface with given key.
//
// Returns error `*InterfaceNotFoundError` if none was found.
func (c Interfaces) Get(key *InterfaceKey) (*Interface, error) {
	for _, si := range c {
		if *key == *si.Key() {
			return si, nil
		}
	}

	return nil, &InterfaceNotFoundError{Key: key}
}

// GroupByBaseName groups the interfaces by base name.
//
// The definition of base name is described in `InterfaceKey`.
func (c Interfaces) GroupByBaseName() InterfaceGroupsByName {
	return NewInterfaceGroupsByName(c)
}

// InterfacesIndex is an index of `Interface`s keyed by `InterfaceKey`.
type InterfacesIndex map[InterfaceKey]*Interface

// NewInterfacesIndex creates an `InterfacesIndex`.
func NewInterfacesIndex(interfaces Interfaces) InterfacesIndex {
	if interfaces == nil {
		return nil
	}

	result := make(InterfacesIndex, len(interfaces))

	for _, si := range interfaces {
		result[*si.Key()] = si
	}

	return result
}

// Name returns the base name of the first `Interface` in the index.
//
// The definition of base name is described in `InterfaceKey`.
//
// This should be used with an `InterfaceIndex` within an `InterfaceGroupsByName`.
//
// Returns an empty string if the index is empty.
func (i InterfacesIndex) Name() string {
	for key := range i {
		return key.Name
	}

	return ""
}

// AppIDs collects the AppID of all interfaces in the index.
//
// Interfaces with no AppID (0) are omitted.
func (i InterfacesIndex) AppIDs() []steamlib.AppID {
	if i == nil {
		return nil
	}

	result := make([]steamlib.AppID, 0, len(i))

	for key := range i {
		if key.AppID != 0 {
			result = append(result, key.AppID)
		}
	}

	return result
}

// GroupMethodsByName groups all methods in the index by method name.
func (i InterfacesIndex) GroupMethodsByName() MethodGroupsByName {
	if i == nil {
		return nil
	}

	var result MethodGroupsByName

	for _, si := range i {
		methodGroups := si.Methods.GroupByName()

		if result == nil {
			result = methodGroups
			continue
		}

		for methodName, methodsIndex := range methodGroups {
			if result[methodName] == nil {
				result[methodName] = methodsIndex
				continue
			}

			maps.Copy(result[methodName], methodsIndex)
		}
	}

	return result
}

/*
InterfaceGroupsByName groups `Interface`s with the same base name.

The definition of base name is described in `InterfaceKey`.

It's a regular map and therefore provides no guarantees on consistency:

  - Keys are not guaranteed to be correctly associated to their respective interfaces
  - Interfaces are not guaranteed to be unique for each key
  - Interfaces are not guaranteed to have the same base name

The group creator is responsible for ensuring consistency. Groups created with
`NewInterfaceGroupsByName` or `Interfaces.GroupByName` can be considered consistent.

Behavior of inconsistent groups is undefined.
*/
type InterfaceGroupsByName map[string]InterfacesIndex

// NewInterfaceGroupsByName groups interfaces as described in `InterfaceGroupsByName`.
func NewInterfaceGroupsByName(interfaces Interfaces) InterfaceGroupsByName {
	if interfaces == nil {
		return nil
	}

	result := make(InterfaceGroupsByName)

	for _, si := range interfaces {
		key := si.Key()

		if result[key.Name] == nil {
			result[key.Name] = make(InterfacesIndex)
		}

		result[key.Name][*key] = si
	}

	return result
}
