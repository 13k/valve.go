package steamlib

import (
	"encoding"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"strconv"
)

var (
	_ fmt.Stringer               = (*AppID)(nil)
	_ encoding.BinaryMarshaler   = (*AppID)(nil)
	_ encoding.BinaryUnmarshaler = (*AppID)(nil)
	_ encoding.TextMarshaler     = (*AppID)(nil)
	_ encoding.TextUnmarshaler   = (*AppID)(nil)
	_ json.Marshaler             = (*AppID)(nil)
	_ json.Unmarshaler           = (*AppID)(nil)
)

// AppID represents an Steam application ID.
type AppID uint32

// ParseAppID parses a string into an AppID.
//
// Returns `*AppIDParseError` on error.
func ParseAppID(s string) (AppID, error) {
	appID, err := strconv.ParseUint(s, 10, 32)

	if err != nil {
		return 0, &AppIDParseError{value: s, err: err}
	}

	return AppID(appID), nil
}

// Uint32 returns the id as uint32.
func (id *AppID) Uint32() uint32 {
	return uint32(*id)
}

// Uint32 returns the id as uint64.
func (id *AppID) Uint64() uint64 {
	return uint64(*id)
}

// String implements interface `fmt.Stringer`.
//
// It returns the id formatted as string using base 10.
func (id *AppID) String() string {
	return strconv.FormatUint(id.Uint64(), 10)
}

// MarshalBinary implements interface `encoding.BinaryMarshaler`.
//
// It assumes little endianness.
//
// Returns `*AppIDMarshalBinaryError` on error.
func (id *AppID) MarshalBinary() ([]byte, error) {
	data, err := binary.Append(nil, binary.LittleEndian, id)

	if err != nil {
		return nil, &AppIDMarshalBinaryError{value: *id, err: err}
	}

	return data, nil
}

// UnmarshalBinary implements interface `encoding.BinaryUnmarshaler`.
//
// It assumes little endianness.
//
// Returns `*AppIDUnmarshalBinaryError` on error.
func (id *AppID) UnmarshalBinary(data []byte) error {
	_, err := binary.Decode(data, binary.LittleEndian, id)

	if err != nil {
		return &AppIDUnmarshalBinaryError{value: data, err: err}
	}

	return nil
}

// MarshalText implements interface `encoding.TextMarshaler`.
//
// It uses base 10.
//
// Returns `*AppIDMarshalTextError` on error.
func (id *AppID) MarshalText() ([]byte, error) {
	return []byte(id.String()), nil
}

// UnmarshalText implements interface `encoding.TextUnmarshaler`.
//
// It assumes base 10.
//
// Returns `*AppIDUnmarshalTextError` on error.
func (id *AppID) UnmarshalText(data []byte) error {
	n, err := ParseAppID(string(data))

	if err != nil {
		return &AppIDUnmarshalTextError{value: data, err: err}
	}

	*id = n

	return nil
}

// MarshalJSON implements interface `encoding/json.Marshaler`
//
// It assumes base 10.
//
// Returns `*AppIDMarshalJSONError` on error.
func (id *AppID) MarshalJSON() ([]byte, error) {
	data, err := id.MarshalText()

	if err != nil {
		return nil, &AppIDMarshalJSONError{value: *id, err: err}
	}

	return data, nil
}

// UnmarshalJSON implements interface `encoding/json.Unmarshaler`
//
// It assumes base 10.
//
// Returns `*AppIDUnmarshalJSONError` on error.
func (id *AppID) UnmarshalJSON(data []byte) error {
	if err := id.UnmarshalText(data); err != nil {
		return &AppIDUnmarshalJSONError{value: data, err: err}
	}

	return nil
}
