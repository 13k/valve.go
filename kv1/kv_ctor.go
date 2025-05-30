package kv1

//nolint:gochecknoglobals
var (
	kvEmptyValue any
	kvEmptyKey   string
	kvNoParent   *KeyValue
)

// NewKeyValue creates a KeyValue.
//
// It returns an error if ty and value are not equivalent.
func NewKeyValue(ty Type, key string, value any, parent *KeyValue) (*KeyValue, error) {
	if err := kvValidate(ty, key, value); err != nil {
		return nil, err
	}

	kv := &KeyValue{
		ty:       ty,
		key:      key,
		value:    value,
		parent:   parent,
		children: nil,
	}

	if parent != nil {
		parent.AddChild(kv)
	}

	return kv, nil
}

// MustNewKeyValue is the same as NewKeyValue but panics on errors.
func MustNewKeyValue(ty Type, key string, value any, parent *KeyValue) *KeyValue {
	kv, err := NewKeyValue(ty, key, value, parent)

	if err != nil {
		panic(err)
	}

	return kv
}

// NewKeyValueEmpty creates an empty KeyValue with type `TypeInvalid`.
func NewKeyValueEmpty() *KeyValue {
	return MustNewKeyValue(TypeInvalid, kvEmptyKey, kvEmptyValue, kvNoParent)
}

// NewKeyValueEmptyChild creates an empty child KeyValue with type `TypeInvalid`.
func NewKeyValueEmptyChild(parent *KeyValue) *KeyValue {
	return MustNewKeyValue(TypeInvalid, kvEmptyKey, kvEmptyValue, parent)
}

// NewKeyValueEnd creates an end marker KeyValue with type `TypeEnd`.
func NewKeyValueEnd() *KeyValue {
	return MustNewKeyValue(TypeEnd, kvEmptyKey, kvEmptyValue, kvNoParent)
}

// NewKeyValueObject creates an object KeyValue with type `TypeObject`.
func NewKeyValueObject(key string, parent *KeyValue) *KeyValue {
	return MustNewKeyValue(TypeObject, key, kvEmptyValue, parent)
}

// NewKeyValueObjectRoot creates a root object KeyValue.
func NewKeyValueObjectRoot(key string) *KeyValue {
	return NewKeyValueObject(key, kvNoParent)
}

// NewKeyValueString creates a field KeyValue with type `TypeString`.
func NewKeyValueString(key string, value string, parent *KeyValue) *KeyValue {
	return MustNewKeyValue(TypeString, key, value, parent)
}

// NewKeyValueWString creates a field KeyValue with type `TypeWString`.
func NewKeyValueWString(key string, value string, parent *KeyValue) *KeyValue {
	return MustNewKeyValue(TypeWString, key, value, parent)
}

// NewKeyValueInt32 creates a field KeyValue with type `TypeInt32`.
func NewKeyValueInt32(key string, value int32, parent *KeyValue) *KeyValue {
	return MustNewKeyValue(TypeInt32, key, value, parent)
}

// NewKeyValueColor creates a field KeyValue with type `TypeColor`.
func NewKeyValueColor(key string, value int32, parent *KeyValue) *KeyValue {
	return MustNewKeyValue(TypeColor, key, value, parent)
}

// NewKeyValuePointer creates a field KeyValue with type `TypePointer`.
func NewKeyValuePointer(key string, value int32, parent *KeyValue) *KeyValue {
	return MustNewKeyValue(TypePointer, key, value, parent)
}

// NewKeyValueInt64 creates a field KeyValue with type `TypeInt64`.
func NewKeyValueInt64(key string, value int64, parent *KeyValue) *KeyValue {
	return MustNewKeyValue(TypeInt64, key, value, parent)
}

// NewKeyValueUint64 creates a field KeyValue with type `TypeUint64`.
func NewKeyValueUint64(key string, value uint64, parent *KeyValue) *KeyValue {
	return MustNewKeyValue(TypeUint64, key, value, parent)
}

// NewKeyValueFloat32 creates a field KeyValue with type `TypeFloat32`.
func NewKeyValueFloat32(key string, value float32, parent *KeyValue) *KeyValue {
	return MustNewKeyValue(TypeFloat32, key, value, parent)
}

func kvValidate(ty Type, key string, value any) error {
	if err := kvValidateKey(ty, key); err != nil {
		return err
	}

	if err := kvValidateValue(ty, value); err != nil {
		return err
	}

	return nil
}

func kvValidateKey(ty Type, key string) error {
	switch ty {
	case TypeInvalid, TypeEnd:
		{
			if key != "" {
				return kvKeyError(ty, key)
			}
		}
	case
		TypeObject,
		TypeString, TypeWString,
		TypeInt32, TypePointer, TypeColor,
		TypeInt64,
		TypeUint64,
		TypeFloat32:
		{
			if key == "" {
				return kvKeyError(ty, key)
			}
		}
	}

	return nil
}

func kvValidateValue(ty Type, value any) error { //nolint:cyclop
	switch ty {
	case TypeInvalid, TypeObject, TypeEnd:
		{
			if value != nil {
				return kvValueError(ty, value)
			}
		}
	case TypeString, TypeWString:
		{
			if _, ok := value.(string); !ok {
				return kvValueError(ty, value)
			}
		}
	case TypeInt32, TypePointer, TypeColor:
		{
			if _, ok := value.(int32); !ok {
				return kvValueError(ty, value)
			}
		}
	case TypeInt64:
		{
			if _, ok := value.(int64); !ok {
				return kvValueError(ty, value)
			}
		}
	case TypeUint64:
		{
			if _, ok := value.(uint64); !ok {
				return kvValueError(ty, value)
			}
		}
	case TypeFloat32:
		{
			if _, ok := value.(float32); !ok {
				return kvValueError(ty, value)
			}
		}
	}

	return nil
}
