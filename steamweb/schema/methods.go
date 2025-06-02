package schema

// Methods is a collection of `Method`s.
type Methods []*Method

// NewMethods creates a new `Method` collection.
//
// Returns `Methods.Validate` errors.
func NewMethods(methods ...*Method) (Methods, error) {
	c := Methods(methods)

	if err := c.Validate(); err != nil {
		return nil, err
	}

	return c, nil
}

// MustNewMethods is like `NewMethods` but panics on errors.
func MustNewMethods(methods ...*Method) Methods {
	c, err := NewMethods(methods...)

	if err != nil {
		panic(err)
	}

	return c
}

// Validate checks whether all methods in the collection are valid.
//
// Returns `Method.Validate` errors.
func (c Methods) Validate() error {
	for _, m := range c {
		if err := m.Validate(); err != nil {
			return err
		}
	}

	return nil
}

// Get returns the method with the given key.
//
// Returns error `*InterfaceMethodNotFoundError` if none was found.
func (c Methods) Get(key *MethodKey) (*Method, error) {
	for _, m := range c {
		if *key == *m.Key() {
			return m, nil
		}
	}

	return nil, &MethodNotFoundError{Key: key}
}

// GroupByName groups the methods by name.
func (c Methods) GroupByName() MethodGroupsByName {
	return NewMethodGroupsByName(c)
}

// MethodsIndex is an index of `Method`s keyed by `MethodKey`.
type MethodsIndex map[MethodKey]*Method

// NewMethodsIndex creates a `MethodIndex`.
func NewMethodsIndex(methods Methods) MethodsIndex {
	if methods == nil {
		return nil
	}

	result := make(MethodsIndex, len(methods))

	for _, sm := range methods {
		result[*sm.Key()] = sm
	}

	return result
}

// Versions collects the versions of all methods in the index.
func (i MethodsIndex) Versions() []int {
	if i == nil {
		return nil
	}

	result := make([]int, 0, len(i))

	for key := range i {
		result = append(result, key.Version)
	}

	return result
}

/*
MethodGroupsByName groups `Method`s with the same name.

It's a regular map and therefore provides no guarantees on consistency:

  - Keys are not guaranteed to be correctly associated to their respective methods
  - Methods are not guaranteed to be unique for each key
  - Methods are not guaranteed to have the same name

The group creator is responsible for ensuring consistency. Groups created with
`NewMethodsGroupByName` or `Methods.GroupByName` can be considered consistent.

Behavior of inconsistent groups is undefined.
*/
type MethodGroupsByName map[string]MethodsIndex

// NewMethodGroupsByName groups methods as described in `MethodGroupsByName`.
func NewMethodGroupsByName(methods Methods) MethodGroupsByName {
	if methods == nil {
		return nil
	}

	result := make(MethodGroupsByName)

	for _, sm := range methods {
		key := sm.Key()

		if result[key.Name] == nil {
			result[key.Name] = make(MethodsIndex)
		}

		result[key.Name][*key] = sm
	}

	return result
}
