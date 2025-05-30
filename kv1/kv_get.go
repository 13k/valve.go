package kv1

// String returns the KeyValue's value as `string`.
//
// It returns an error if the KeyValue' type is not `TypeString`.
func (kv *KeyValue) String() (string, error) {
	return kvGet[string](kv, TypeString)
}

// WString returns the KeyValue's value as `string`.
//
// It returns an error if the KeyValue' type is not `TypeWString`.
func (kv *KeyValue) WString() (string, error) {
	return kvGet[string](kv, TypeWString)
}

// Int32 returns the KeyValue's value as `int32`.
//
// It returns an error if the KeyValue's type is not `TypeInt32`.
func (kv *KeyValue) Int32() (int32, error) {
	return kvGet[int32](kv, TypeInt32)
}

// Color returns the KeyValue's value as `int32`.
//
// It returns an error if the KeyValue's type is not `TypeColor`.
func (kv *KeyValue) Color() (int32, error) {
	return kvGet[int32](kv, TypeColor)
}

// Pointer returns the KeyValue's value as `int32`.
//
// It returns an error if the KeyValue's type is not `TypePointer`.
func (kv *KeyValue) Pointer() (int32, error) {
	return kvGet[int32](kv, TypePointer)
}

// Int64 returns the KeyValue's value as `int64`.
//
// It returns an error if the KeyValue's type is not `TypeInt64`.
func (kv *KeyValue) Int64() (int64, error) {
	return kvGet[int64](kv, TypeInt64)
}

// Uint64 returns the KeyValue's value as `uint64`.
//
// It returns an error if the KeyValue's type is not `TypeUint64`.
func (kv *KeyValue) Uint64() (uint64, error) {
	return kvGet[uint64](kv, TypeUint64)
}

// Float32 returns the KeyValue's value as `float32`.
//
// It returns an error if the KeyValue's type is not `TypeFloat32`.
func (kv *KeyValue) Float32() (float32, error) {
	return kvGet[float32](kv, TypeFloat32)
}

func kvGet[T KeyValueT](kv *KeyValue, ty Type) (T, error) { //nolint:ireturn
	var zvalue T

	if kv.Type() != ty {
		return zvalue, kvGetError(kv, ty)
	}

	value, ok := kv.value.(T)

	if !ok {
		return zvalue, kvGetError(kv, ty)
	}

	return value, nil
}
