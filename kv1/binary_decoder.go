package kv1

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"io"
)

const (
	binDecStringDelim byte = 0x0
)

// BinaryDecoder reads and decodes binary-encoded KeyValue nodes from an input stream.
type BinaryDecoder struct {
	r *bufio.Reader
}

// NewBinaryDecoder returns a new binary decoder that reads from r.
//
// It wraps `r` with a `bufio.Reader`.
func NewBinaryDecoder(r io.Reader) *BinaryDecoder {
	return &BinaryDecoder{r: bufio.NewReader(r)}
}

// Decode reads the next binary-encoded KeyValue node from its input and stores it in the value
// pointed to by kv.
func (d *BinaryDecoder) Decode(kv *KeyValue) error { //nolint:cyclop
	if err := d.decodeType(kv); err != nil {
		return err
	}

	ty := kv.Type()

	if ty == TypeEnd {
		return nil
	}

	if err := d.decodeKey(kv); err != nil {
		return err
	}

	key := kv.Key()

	if ty == TypeInvalid || ty == TypeWString {
		return binDecUnsupportedError(ty, key)
	}

	var err error

	switch ty {
	case TypeObject:
		err = d.decodeObject(kv)
	case TypeString:
		err = d.decodeString(kv)
	case TypeInt32:
		err = d.decodeInt32(kv)
	case TypeColor:
		err = d.decodeColor(kv)
	case TypePointer:
		err = d.decodePointer(kv)
	case TypeInt64:
		err = d.decodeInt64(kv)
	case TypeUint64:
		err = d.decodeUint64(kv)
	case TypeFloat32:
		err = d.decodeFloat32(kv)
	case TypeEnd, TypeInvalid, TypeWString:
		panic("unreachable")
	}

	if err != nil {
		return err
	}

	return nil
}

func (d *BinaryDecoder) readObject() ([]*KeyValue, error) {
	var kvs []*KeyValue

	for {
		kv := NewKeyValueEmpty()

		if err := d.Decode(kv); err != nil {
			return nil, fmt.Errorf("failed to read object value: %w", err)
		}

		if kv.Type() == TypeEnd {
			break
		}

		kvs = append(kvs, kv)
	}

	return kvs, nil
}

func (d *BinaryDecoder) readType() (Type, error) {
	b, err := d.r.ReadByte()

	if err != nil {
		return TypeInvalid, fmt.Errorf("failed to read node type: %w", err)
	}

	return NewType(b), nil
}

func (d *BinaryDecoder) readKey() (string, error) {
	key, err := d.readString()

	if err != nil {
		return "", fmt.Errorf("failed to read node key: %w", err)
	}

	return key, nil
}

func (d *BinaryDecoder) readString() (string, error) {
	s, err := d.r.ReadString(binDecStringDelim)

	if err != nil {
		return "", fmt.Errorf("failed to read string value: %w", err)
	}

	return s[:len(s)-1], nil
}

func (d *BinaryDecoder) readInt32() (int32, error) {
	return readNumber[int32](d.r)
}

func (d *BinaryDecoder) readInt64() (int64, error) {
	return readNumber[int64](d.r)
}

func (d *BinaryDecoder) readUint64() (uint64, error) {
	return readNumber[uint64](d.r)
}

func (d *BinaryDecoder) readFloat32() (float32, error) {
	return readNumber[float32](d.r)
}

func readNumber[T int32 | int64 | uint64 | float32](r io.Reader) (T, error) { //nolint:ireturn
	var n T

	if err := binary.Read(r, binary.LittleEndian, &n); err != nil {
		return 0, fmt.Errorf("failed to read %T value: %w", n, err)
	}

	return n, nil
}

func (d *BinaryDecoder) decodeType(kv *KeyValue) error {
	ty, err := d.readType()

	if err != nil {
		return binDecReadTypeError(err)
	}

	kv.SetType(ty)

	return nil
}

func (d *BinaryDecoder) decodeKey(kv *KeyValue) error {
	key, err := d.readKey()

	if err != nil {
		return binDecReadKeyError(err, kv.Type())
	}

	kv.SetKey(key)

	return nil
}

func (d *BinaryDecoder) decodeObject(kv *KeyValue) error {
	children, err := d.readObject()

	if err != nil {
		return binDecError(err, kv.Type(), kv.Key())
	}

	kv.SetChildren(children)

	return nil
}

func (d *BinaryDecoder) decodeString(kv *KeyValue) error {
	value, err := d.readString()

	if err != nil {
		return binDecError(err, kv.Type(), kv.Key())
	}

	kv.SetString(value)

	return nil
}

func (d *BinaryDecoder) decodeInt32(kv *KeyValue) error {
	value, err := d.readInt32()

	if err != nil {
		return binDecError(err, kv.Type(), kv.Key())
	}

	kv.SetInt32(value)

	return nil
}

func (d *BinaryDecoder) decodeColor(kv *KeyValue) error {
	value, err := d.readInt32()

	if err != nil {
		return binDecError(err, kv.Type(), kv.Key())
	}

	kv.SetColor(value)

	return nil
}

func (d *BinaryDecoder) decodePointer(kv *KeyValue) error {
	value, err := d.readInt32()

	if err != nil {
		return binDecError(err, kv.Type(), kv.Key())
	}

	kv.SetPointer(value)

	return nil
}

func (d *BinaryDecoder) decodeInt64(kv *KeyValue) error {
	value, err := d.readInt64()

	if err != nil {
		return binDecError(err, kv.Type(), kv.Key())
	}

	kv.SetInt64(value)

	return nil
}

func (d *BinaryDecoder) decodeUint64(kv *KeyValue) error {
	value, err := d.readUint64()

	if err != nil {
		return binDecError(err, kv.Type(), kv.Key())
	}

	kv.SetUint64(value)

	return nil
}

func (d *BinaryDecoder) decodeFloat32(kv *KeyValue) error {
	value, err := d.readFloat32()

	if err != nil {
		return binDecError(err, kv.Type(), kv.Key())
	}

	kv.SetFloat32(value)

	return nil
}
