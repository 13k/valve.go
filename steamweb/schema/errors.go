package schema

import (
	"errors"
	"fmt"
)

var (
	// ErrEmptyInterfaces is returned when trying to manipulate an empty `Interfaces`.
	ErrEmptyInterfaces = errors.New("empty interfaces collection")
	// ErrMixedInterfaces is returned when trying to use group helpers on a `Interfaces` collection
	// containing mixed interface names.
	ErrMixedInterfaces = errors.New("mixed interfaces collection")
	// ErrMixedMethods is returned when trying to use group helpers on a `Methods` collection
	// containing mixed method names.
	ErrMixedMethods = errors.New("mixed methods collection")
)

//// Schema {{{

var _ error = (*InvalidSchemaError)(nil)

// InvalidSchemaError is returned when a `Schema` is invalid.
type InvalidSchemaError struct {
	err error
}

// Error implements interface `error`.
func (err *InvalidSchemaError) Error() string {
	return fmt.Sprintf("invalid Schema: %s", err.err.Error())
}

// Unwrap implements `errors.Unwrap`.
func (err *InvalidSchemaError) Unwrap() error {
	return err.err
}

//// Schema }}}
//// Interface {{{

// InvalidInterfaceError is returned when an `Interface` is invalid.
type InvalidInterfaceError struct {
	err error
}

// Error implements interface `error`.
func (err *InvalidInterfaceError) Error() string {
	return fmt.Sprintf("invalid Interface: %s", err.err.Error())
}

// Unwrap implements `errors.Unwrap`.
func (err *InvalidInterfaceError) Unwrap() error {
	return err.err
}

var _ error = (*InvalidInterfaceNameError)(nil)

// InvalidInterfaceNameError is returned when an interface has an invalid name.
type InvalidInterfaceNameError struct {
	Name string
	err  error
}

// Error implements interface `error`.
func (err *InvalidInterfaceNameError) Error() string {
	return fmt.Sprintf("invalid name %q", err.Name)
}

// Unwrap implements `errors.Unwrap`.
func (err *InvalidInterfaceNameError) Unwrap() error {
	return err.err
}

var _ error = (*InterfaceNotFoundError)(nil)

// InterfaceNotFoundError is returned when tried to access an interface with invalid name or invalid
// AppID.
type InterfaceNotFoundError struct {
	Key *InterfaceKey
}

// Error implements interface `error`.
func (err *InterfaceNotFoundError) Error() string {
	return fmt.Sprintf("interface not found %s", err.Key.String())
}

//// Interface }}}
//// Method {{{

var _ error = (*InvalidMethodError)(nil)

// InvalidMethodError is returned when a `Method` is invalid.
type InvalidMethodError struct {
	err error
}

// Error implements interface `error`.
func (err *InvalidMethodError) Error() string {
	return fmt.Sprintf("invalid Method: %s", err.err.Error())
}

// Unwrap implements `errors.Unwrap`.
func (err *InvalidMethodError) Unwrap() error {
	return err.err
}

var _ error = (*InvalidMethodNameError)(nil)

// InvalidMethodNameError is returned when a method has an invalid name.
type InvalidMethodNameError struct {
	Name    string
	Version int
}

func NewInvalidMethodNameError(sm *Method) *InvalidMethodNameError {
	return &InvalidMethodNameError{
		Name:    sm.Name,
		Version: sm.Version,
	}
}

// Error implements interface `error`.
func (err *InvalidMethodNameError) Error() string {
	return fmt.Sprintf("invalid name %q", err.Name)
}

var _ error = (*InvalidMethodVersionError)(nil)

// InvalidMethodVersionError is returned when a method has an invalid version.
type InvalidMethodVersionError struct {
	Name    string
	Version int
}

func NewInvalidMethodVersionError(sm *Method) *InvalidMethodVersionError {
	return &InvalidMethodVersionError{
		Name:    sm.Name,
		Version: sm.Version,
	}
}

// Error implements interface `error`.
func (err *InvalidMethodVersionError) Error() string {
	return fmt.Sprintf("invalid version %d", err.Version)
}

var _ error = (*InvalidMethodHTTPMethodError)(nil)

// InvalidMethodHTTPMethodError is returned when a method has an invalid version.
type InvalidMethodHTTPMethodError struct {
	Name       string
	Version    int
	HTTPMethod string
}

func NewInvalidMethodHTTPMethodError(sm *Method) *InvalidMethodHTTPMethodError {
	return &InvalidMethodHTTPMethodError{
		Name:       sm.Name,
		Version:    sm.Version,
		HTTPMethod: sm.HTTPMethod,
	}
}

// Error implements interface `error`.
func (err *InvalidMethodHTTPMethodError) Error() string {
	return fmt.Sprintf("invalid HTTP method %q", err.HTTPMethod)
}

var _ error = (*MethodNotFoundError)(nil)

// MethodNotFoundError is returned when tried to access an interface method with invalid
// name or invalid version.
type MethodNotFoundError struct {
	Key *MethodKey
}

// Error implements interface `error`.
func (err *MethodNotFoundError) Error() string {
	return fmt.Sprintf("interface method not found %s", err.Key.String())
}

//// Method }}}
//// MethodParam {{{

var _ error = (*RequiredParameterError)(nil)

// RequiredParameterError is returned when a required parameter is missing.
type RequiredParameterError struct {
	Param *MethodParam
}

// Error implements interface `error`.
func (err *RequiredParameterError) Error() string {
	return fmt.Sprintf("missing required parameter %q", err.Param.Name)
}

//// MethodParam }}}
