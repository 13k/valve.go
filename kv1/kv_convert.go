package kv1

import (
	"strconv"

	"github.com/ccoveille/go-safecast"
)

// ToString returns the KeyValue's value formatted as a string.
//
// It returns an error if formatting fails or if the KeyValue holds no value (HasValue is false).
func (kv *KeyValue) ToString() (string, error) {
	if !kv.IsField() {
		return "", kvFormatTypeError(kv)
	}

	switch value := kv.value.(type) {
	case string:
		return value, nil
	case int32:
		return strconv.FormatInt(int64(value), 10), nil
	case int64:
		return strconv.FormatInt(value, 10), nil
	case uint64:
		return strconv.FormatUint(value, 10), nil
	case float32:
		return strconv.FormatFloat(float64(value), 'f', -1, 32), nil
	default:
		return "", kvFormatValueError(kv.value)
	}
}

// ToInt32 converts the KeyValue's value to int32.
//
// It returns an error if Type is not a numeric or string type or if the conversion fails.
func (kv *KeyValue) ToInt32() (int32, error) {
	return kvSafeConvert[int32](kv)
}

// ToColor is an alias to ToInt32.
func (kv *KeyValue) ToColor() (int32, error) {
	return kv.ToInt32()
}

// ToPointer is an alias to ToInt32.
func (kv *KeyValue) ToPointer() (int32, error) {
	return kv.ToInt32()
}

// ToInt64 converts the KeyValue's value to int64.
//
// It returns an error if Type is not a numeric or string type or if the conversion fails.
func (kv *KeyValue) ToInt64() (int64, error) {
	return kvSafeConvert[int64](kv)
}

// ToUint64 converts the KeyValue's value to uint64.
//
// It returns an error if Type is not a numeric or string type or if the conversion fails.
func (kv *KeyValue) ToUint64() (uint64, error) {
	return kvSafeConvert[uint64](kv)
}

// ToFloat32 converts the KeyValue's value to float32.
//
// It returns an error if Type is not a numeric or string type or if the conversion fails.
func (kv *KeyValue) ToFloat32() (float32, error) {
	return kvSafeConvert[float32](kv)
}

func safeConvert[Out safecast.Number, In KeyValueT](kv *KeyValue, input In) (Out, error) { //nolint:ireturn
	out, err := safecast.Convert[Out](input)

	if err != nil {
		return 0, kvConvertError[Out](kv, err)
	}

	return out, nil
}

func kvSafeConvert[Out safecast.Number](kv *KeyValue) (Out, error) { //nolint:ireturn
	if !kv.IsField() {
		return 0, kvConvertTypeError[Out](kv)
	}

	switch value := kv.value.(type) {
	case string:
		{
			return safeConvert[Out](kv, value)
		}
	case int32:
		{
			return safeConvert[Out](kv, value)
		}
	case int64:
		{
			return safeConvert[Out](kv, value)
		}
	case uint64:
		{
			return safeConvert[Out](kv, value)
		}
	case float32:
		{
			return safeConvert[Out](kv, value)
		}
	default:
		{
			return 0, kvConvertTypeError[uint64](kv)
		}
	}
}
