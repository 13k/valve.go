package kv1

import (
	"errors"
	"fmt"
)

var (
	ErrUnsupportedType = errors.New("type is not supported")
	ErrIntegerOverflow = errors.New("integer overflow")
	ErrFloatOverflow   = errors.New("float overflow")
	ErrFloatInvalid    = errors.New("invalid float value")
)

///////////////////////////////// general

func writeError(err error) error {
	if err == nil {
		return nil
	}

	return fmt.Errorf("failed to write to stream: %w", err)
}

///////////////////////////////// KeyValue

var _ error = (*KeyValueError)(nil)

type KeyValueError struct {
	msg string
}

func (err *KeyValueError) Error() string {
	return "KeyValue error: " + err.msg
}

func kvError(msg string) error {
	return &KeyValueError{msg: msg}
}

func kvKeyError(ty Type, key string) error {
	return kvError(fmt.Sprintf("invalid key %q for type %s", key, ty))
}

func kvValueError(ty Type, value any) error {
	return kvError(fmt.Sprintf("cannot set value of type %s with value %v (%T)", ty, value, value))
}

func kvGetError(kv *KeyValue, ty Type) error {
	return kvError(fmt.Sprintf("cannot get value of type %s as %s", kv.Type(), ty))
}

func kvFormatTypeError(kv *KeyValue) error {
	return kvError(fmt.Sprintf("cannot format type %s", kv.Type()))
}

func kvFormatValueError(value any) error {
	return kvError(fmt.Sprintf("cannot format value %v (%T)", value, value))
}

func kvMapTypeError(kv *KeyValue) error {
	return kvError(fmt.Sprintf("cannot convert type %s to map", kv.Type()))
}

func kvConvertTypeError[T any](kv *KeyValue) error {
	return kvError(fmt.Sprintf("cannot convert type %s to %T", kv.Type(), *new(T)))
}

var _ error = (*KeyValueConvertError)(nil)

type KeyValueConvertError struct {
	err   error
	kv    *KeyValue
	value any
}

func (err *KeyValueConvertError) Error() string {
	return fmt.Sprintf(
		"failed to convert value %v (%T) to %T: %s",
		err.kv.Value(),
		err.kv.Value(),
		err.value,
		err.err,
	)
}

func (err *KeyValueConvertError) Unwrap() error {
	return err.err
}

func kvConvertError[T any](kv *KeyValue, err error) error {
	if err == nil {
		return nil
	}

	return &KeyValueConvertError{
		err:   err,
		kv:    kv,
		value: *new(T),
	}
}

///////////////////////////////// Binary encoding/decoding

var _ error = (*BinaryEncodeError)(nil)

type BinaryEncodeError struct {
	err error
	kv  *KeyValue
}

func (err *BinaryEncodeError) Error() string {
	return fmt.Sprintf(
		"binary encode error: failed to encode key %q with type %s: %s",
		err.kv.Key(),
		err.kv.Type(),
		err.err,
	)
}

func (err *BinaryEncodeError) Unwrap() error {
	return err.err
}

func binEncError(err error, kv *KeyValue) error {
	if err == nil {
		return nil
	}

	return &BinaryEncodeError{
		kv:  kv,
		err: err,
	}
}

func binEncUnsupportedError(kv *KeyValue) error {
	return binEncError(ErrUnsupportedType, kv)
}

var _ error = (*BinaryDecodeError)(nil)

type BinaryDecodeError struct {
	err error
	ty  Type
	key string
}

func (err *BinaryDecodeError) Error() string {
	if err.key != "" {
		return fmt.Sprintf(
			"binary decode error: failed to decode key %q with type %s: %s",
			err.key,
			err.ty,
			err.err,
		)
	}

	if err.ty != TypeInvalid {
		return fmt.Sprintf(
			"binary decode error: failed to decode value with type %s: %s",
			err.ty,
			err.err,
		)
	}

	return fmt.Sprintf(
		"binary decode error: %s",
		err.err,
	)
}

func (err *BinaryDecodeError) Unwrap() error {
	return err.err
}

func binDecError(err error, ty Type, key string) error {
	if err == nil {
		return nil
	}

	return &BinaryDecodeError{
		err: err,
		ty:  ty,
		key: key,
	}
}

func binDecReadTypeError(err error) error {
	return binDecError(err, TypeInvalid, "")
}

func binDecReadKeyError(err error, ty Type) error {
	return binDecError(err, ty, "")
}

func binDecUnsupportedError(ty Type, key string) error {
	return binDecError(ErrUnsupportedType, ty, key)
}

///////////////////////////////// Text encoding/decoding

var _ error = (*TextEncodeError)(nil)

type TextEncodeError struct {
	err error
	kv  *KeyValue
}

func (err *TextEncodeError) Error() string {
	return fmt.Sprintf(
		"text encode error: failed to encode key %q with type %s: %s",
		err.kv.Key(),
		err.kv.Type(),
		err.err,
	)
}

func (err *TextEncodeError) Unwrap() error {
	return err.err
}

func textEncError(err error, kv *KeyValue) error {
	if err == nil {
		return nil
	}

	return &TextEncodeError{
		kv:  kv,
		err: err,
	}
}

func textEncUnsupportedError(kv *KeyValue) error {
	return textEncError(ErrUnsupportedType, kv)
}

var _ error = (*TextDecodeError)(nil)

type TextDecodeError struct {
	err error
}

func (err *TextDecodeError) Error() string {
	return fmt.Sprintf("text decode error: %s", err.err)
}

func (err *TextDecodeError) Unwrap() error {
	return err.err
}

func textDecError(err error) error {
	if err == nil {
		return nil
	}

	return &TextDecodeError{err: err}
}
