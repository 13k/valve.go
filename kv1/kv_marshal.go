package kv1

import (
	"bytes"
	"encoding"
)

var _ encoding.BinaryMarshaler = (*KeyValue)(nil)
var _ encoding.BinaryUnmarshaler = (*KeyValue)(nil)
var _ encoding.TextMarshaler = (*KeyValue)(nil)
var _ encoding.TextUnmarshaler = (*KeyValue)(nil)

// MarshalBinary implements the `encoding.BinaryMarshaler` interface.
func (kv *KeyValue) MarshalBinary() ([]byte, error) {
	b := &bytes.Buffer{}

	if err := NewBinaryEncoder(b).Encode(kv); err != nil {
		return nil, err
	}

	return b.Bytes(), nil
}

// UnmarshalBinary implements the `encoding.BinaryUnmarshaler` interface.
func (kv *KeyValue) UnmarshalBinary(data []byte) error {
	return NewBinaryDecoder(bytes.NewReader(data)).Decode(kv)
}

// MarshalText implements the `encoding.TextMarshaler` interface.
func (kv *KeyValue) MarshalText() ([]byte, error) {
	b := &bytes.Buffer{}

	if err := NewTextEncoder(b).Encode(kv); err != nil {
		return nil, err
	}

	return b.Bytes(), nil
}

// UnmarshalText implements the `encoding.TextUnmarshaler` interface.
func (kv *KeyValue) UnmarshalText(data []byte) error {
	return NewTextDecoder(bytes.NewReader(data)).Decode(kv)
}
