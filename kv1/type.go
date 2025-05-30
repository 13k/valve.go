package kv1

import "math"

//go:generate stringer -type Type -trimprefix Type

// Type represents a KeyValue's node type.
type Type byte

// KeyValue sentinel type.
const (
	// Empty/invalid.
	TypeInvalid Type = math.MaxUint8 // 0xff
)

// KeyValue types.
const (
	// Collection of KeyValues.
	TypeObject Type = iota // 0x00
	// Field backed by `string`.
	TypeString // 0x01
	// Field backed by `int32`.
	TypeInt32 // 0x02
	// Field backed by `float32`.
	TypeFloat32 // 0x03
	// Field backed by `int32`.
	TypePointer // 0x04
	// Field backed by `string`.
	TypeWString // 0x05
	// Field backed by `int32`.
	TypeColor // 0x06
	// Field backed by `uint64`.
	TypeUint64 // 0x07
	// Marker for end of binary stream.
	TypeEnd // 0x08
	_       // skip
	// Field backed by `int64`.
	TypeInt64 // 0x0a
)

// NewType converts a byte from binary format to a `Type`.
//
// Returns `TypeInvalid` if the given byte is invalid.
func NewType(b byte) Type {
	t := Type(b)

	switch t {
	case
		TypeInvalid,
		TypeObject,
		TypeString,
		TypeInt32,
		TypeFloat32,
		TypePointer,
		TypeWString,
		TypeColor,
		TypeUint64,
		TypeEnd,
		TypeInt64:
		return t
	default:
		return TypeInvalid
	}
}

// Byte returns the corresponding byte in binary format.
//
// Returns `TypeInvalid` byte if the underlying value is invalid.
func (t Type) Byte() byte {
	switch t {
	case
		TypeInvalid,
		TypeObject,
		TypeString,
		TypeInt32,
		TypeFloat32,
		TypePointer,
		TypeWString,
		TypeColor,
		TypeUint64,
		TypeEnd,
		TypeInt64:
		return byte(t)
	default:
		return byte(TypeInvalid)
	}
}
