package kv1

// KeyValueT is a type constraint on the value that KeyValue fields can hold.
type KeyValueT interface {
	~string | // TypeString, TypeWString
		~int32 | // TypeInt32, TypeColor, TypePoint
		~int64 | // TypeInt64
		~uint64 | // TypeUint64
		~float32 // TypeFloat32
}

// KeyValue represents a node in a KeyValue tree.
//
// Nodes are divided into 3 kinds:
//   - special [`TypeEnd`, `TypeInvalid`]: marker for end-of-stream or emptiness.
//   - object [`TypeObject`]: collection of fields. It doesn't hold a value.
//   - field [all other types]: key-value pair. It holds a value and doesn't have children.
type KeyValue struct {
	ty       Type
	key      string
	value    any
	parent   *KeyValue
	children []*KeyValue
}

// Type returns the KeyValue's type.
func (kv *KeyValue) Type() Type { return kv.ty }

// IsField checks whether this KeyValue is a field with a value.
func (kv *KeyValue) IsField() bool {
	return !kv.IsObject() && !kv.IsEnd() && !kv.IsInvalid()
}

// IsObject returns true if this KeyValue's type is `TypeObject`.
func (kv *KeyValue) IsObject() bool { return kv.ty == TypeObject }

// IsString returns true if this KeyValue's type is `TypeString`.
func (kv *KeyValue) IsString() bool { return kv.ty == TypeString }

// IsWString returns true if this KeyValue's type is `TypeWString`.
func (kv *KeyValue) IsWString() bool { return kv.ty == TypeWString }

// IsInt32 returns true if this KeyValue's type is `TypeInt32`.
func (kv *KeyValue) IsInt32() bool { return kv.ty == TypeInt32 }

// IsColor returns true if this KeyValue's type is `TypeColor`.
func (kv *KeyValue) IsColor() bool { return kv.ty == TypeColor }

// IsPointer returns true if this KeyValue's type is `TypePointer`.
func (kv *KeyValue) IsPointer() bool { return kv.ty == TypePointer }

// IsInt64 returns true if this KeyValue's type is `TypeInt64`.
func (kv *KeyValue) IsInt64() bool { return kv.ty == TypeInt64 }

// IsUint64 returns true if this KeyValue's type is `TypeUint64`.
func (kv *KeyValue) IsUint64() bool { return kv.ty == TypeUint64 }

// IsFloat32 returns true if this KeyValue's type is `TypeFloat32`.
func (kv *KeyValue) IsFloat32() bool { return kv.ty == TypeFloat32 }

// IsEnd returns true if this KeyValue's type is `TypeEnd`.
func (kv *KeyValue) IsEnd() bool { return kv.ty == TypeEnd }

// IsInvalid returns true if this KeyValue's type is `TypeInvalid`.
func (kv *KeyValue) IsInvalid() bool { return kv.ty == TypeInvalid }

// SetType sets the type and returns itself.
//
// If ty differs from the current type, the value is unset.
func (kv *KeyValue) SetType(ty Type) *KeyValue {
	if kv.ty != ty {
		kv.ty = ty
		kv.value = kvEmptyValue
	}

	return kv
}

// Key returns the KeyValue's key.
func (kv *KeyValue) Key() string { return kv.key }

// SetKey sets the KeyValue's key and returns itself.
func (kv *KeyValue) SetKey(k string) *KeyValue {
	kv.key = k

	return kv
}

// Value returns the KeyValue's value.
//
// It returns `nil` if this KeyValue is not a field.
func (kv *KeyValue) Value() any {
	if kv.IsField() {
		return kv.value
	}

	return nil
}
