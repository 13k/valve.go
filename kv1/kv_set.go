package kv1

// SetString sets the KeyValue's type to TypeString and sets its value to the given value.
func (kv *KeyValue) SetString(value string) *KeyValue {
	return kvSet(kv, TypeString, value)
}

// SetInt32 sets the KeyValue's type to TypeInt32 and sets its value to the given value.
func (kv *KeyValue) SetInt32(value int32) *KeyValue {
	return kvSet(kv, TypeInt32, value)
}

// SetColor sets the KeyValue's type to TypeColor and sets its value to the given value.
func (kv *KeyValue) SetColor(value int32) *KeyValue {
	return kvSet(kv, TypeColor, value)
}

// SetPointer sets the KeyValue's type to TypePointer and sets its value to the given value.
func (kv *KeyValue) SetPointer(value int32) *KeyValue {
	return kvSet(kv, TypePointer, value)
}

// SetInt64 sets the KeyValue's type to TypeInt64 and sets its value to the given value.
func (kv *KeyValue) SetInt64(value int64) *KeyValue {
	return kvSet(kv, TypeInt64, value)
}

// SetUint64 sets the KeyValue's type to TypeUint64 and sets its value to the given value.
func (kv *KeyValue) SetUint64(value uint64) *KeyValue {
	return kvSet(kv, TypeUint64, value)
}

// SetFloat32 sets the KeyValue's type to TypeFloat32 and sets its value to the given value.
func (kv *KeyValue) SetFloat32(value float32) *KeyValue {
	return kvSet(kv, TypeFloat32, value)
}

func kvSet[T KeyValueT](kv *KeyValue, ty Type, value T) *KeyValue {
	kv.ty = ty
	kv.value = value

	return kv
}
