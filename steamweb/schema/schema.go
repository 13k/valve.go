// Package schema represents the tree structure of an API schema.
//
// For more information, refer to the main "steamweb" package documentation.
package schema

// Schema is the root of an API schema specification.
type Schema struct {
	Interfaces Interfaces `json:"interfaces"`
}

// NewSchema creates a new `Schema`.
//
// Returns `Schema.Validate` errors.
func NewSchema(interfaces Interfaces) (*Schema, error) {
	ss := &Schema{Interfaces: interfaces}

	if err := ss.Validate(); err != nil {
		return nil, err
	}

	return ss, nil
}

// Validate checks if fields are valid.
//
// It should be used after creating or updating a Schema (unmarshalling or direct instantiation
// without `NewSchema`).
//
// Returns `*InvalidSchemaError` on error.
func (s *Schema) Validate() error {
	if err := s.Interfaces.Validate(); err != nil {
		return &InvalidSchemaError{err: err}
	}

	return nil
}
