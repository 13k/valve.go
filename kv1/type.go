package kv1

import "math"

//go:generate stringer -type Type -trimprefix Type

// Type represents a KeyValue's node type.
type Type uint8

// KeyValue valid types.
const (
	// Object (has child KeyValues, holds no value).
	TypeObject Type = iota // 0x00
	// String (value backed by `string`).
	TypeString // 0x01
	// Int32 (value backed by `int32`).
	TypeInt32 // 0x02
	// Float32 (value backed by `float32`).
	TypeFloat32 // 0x03
	// Pointer (value backed by `int32`).
	TypePointer // 0x04
	// WString (value backed by `string`).
	TypeWString // 0x05
	// Color (value backed by `int32`).
	TypeColor // 0x06
	// Uint64 (value backed by `uint64`).
	TypeUint64 // 0x07
	// End (signals end of binary stream, holds no value).
	TypeEnd // 0x08
	_       // skip
	// Int64 (value backed by `int64`).
	TypeInt64 // 0x0a
)

// KeyValue sentinel type.
const (
	// Invalid (empty).
	TypeInvalid Type = math.MaxUint8 // 0xff
)

// TypeFromByte converts a byte from binary format to a Type.
//
// Returns TypeInvalid if the given byte is not a valid Type.
func TypeFromByte(b byte) Type {
	t := Type(b)

	switch t {
	case
		TypeObject,
		TypeString,
		TypeInt32,
		TypeFloat32,
		TypePointer,
		TypeWString,
		TypeColor,
		TypeUint64,
		TypeEnd,
		TypeInt64,
		TypeInvalid:
		return t
	}

	return TypeInvalid
}

// Byte returns the corresponding byte in binary format.
func (t Type) Byte() byte {
	switch t {
	case
		TypeObject,
		TypeString,
		TypeInt32,
		TypeFloat32,
		TypePointer,
		TypeWString,
		TypeColor,
		TypeUint64,
		TypeEnd,
		TypeInt64,
		TypeInvalid:
		return byte(t)
	}

	return byte(TypeInvalid)
}
