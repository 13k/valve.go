package steamlib

import "fmt"

var _ error = (*AppIDParseError)(nil)

type AppIDParseError struct {
	value string
	err   error
}

func (err *AppIDParseError) Error() string {
	return fmt.Sprintf("failed to parse AppID from string %q: %s", err.value, err.err.Error())
}

func (err *AppIDParseError) Unwrap() error {
	return err.err
}

var _ error = (*AppIDMarshalBinaryError)(nil)

type AppIDMarshalBinaryError struct {
	value AppID
	err   error
}

func (err *AppIDMarshalBinaryError) Error() string {
	return fmt.Sprintf("failed to marshal AppID %v to binary: %s", err.value, err.err.Error())
}

func (err *AppIDMarshalBinaryError) Unwrap() error {
	return err.err
}

var _ error = (*AppIDUnmarshalBinaryError)(nil)

type AppIDUnmarshalBinaryError struct {
	value []byte
	err   error
}

func (err *AppIDUnmarshalBinaryError) Error() string {
	return fmt.Sprintf("failed to unmarshal AppID from binary data %q: %s", err.value, err.err.Error())
}

func (err *AppIDUnmarshalBinaryError) Unwrap() error {
	return err.err
}

var _ error = (*AppIDMarshalTextError)(nil)

type AppIDMarshalTextError struct {
	value AppID
	err   error
}

func (err *AppIDMarshalTextError) Error() string {
	return fmt.Sprintf("failed to marshal AppID %v to text: %s", err.value, err.err.Error())
}

func (err *AppIDMarshalTextError) Unwrap() error {
	return err.err
}

var _ error = (*AppIDUnmarshalTextError)(nil)

type AppIDUnmarshalTextError struct {
	value []byte
	err   error
}

func (err *AppIDUnmarshalTextError) Error() string {
	return fmt.Sprintf("failed to unmarshal AppID from text data %q: %s", err.value, err.err.Error())
}

func (err *AppIDUnmarshalTextError) Unwrap() error {
	return err.err
}

var _ error = (*AppIDMarshalJSONError)(nil)

type AppIDMarshalJSONError struct {
	value AppID
	err   error
}

func (err *AppIDMarshalJSONError) Error() string {
	return fmt.Sprintf("failed to marshal AppID %v to JSON: %s", err.value, err.err.Error())
}

func (err *AppIDMarshalJSONError) Unwrap() error {
	return err.err
}

var _ error = (*AppIDUnmarshalJSONError)(nil)

type AppIDUnmarshalJSONError struct {
	value []byte
	err   error
}

func (err *AppIDUnmarshalJSONError) Error() string {
	return fmt.Sprintf("failed to unmarshal AppID from JSON data %q: %s", err.value, err.err.Error())
}

func (err *AppIDUnmarshalJSONError) Unwrap() error {
	return err.err
}
