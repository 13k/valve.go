package kv1

import (
	"bytes"
	"encoding"
	"strconv"
)

const kvEmptyValue int32 = 0

var _ encoding.BinaryMarshaler = (*KeyValue)(nil)
var _ encoding.BinaryUnmarshaler = (*KeyValue)(nil)
var _ encoding.TextMarshaler = (*KeyValue)(nil)
var _ encoding.TextUnmarshaler = (*KeyValue)(nil)

// KeyValue represents a node in a KeyValue tree.
type KeyValue struct {
	ty       Type
	key      string
	value    any
	parent   *KeyValue
	children []*KeyValue
}

type KeyValueT interface {
	~string | ~int32 | ~int64 | ~uint64 | ~float32
}

// NewKeyValue creates a KeyValue node.
//
// It returns an error if `ty` and `value` are not equivalent.
func NewKeyValue[T KeyValueT](ty Type, key string, value T, parent *KeyValue) (*KeyValue, error) {
	if err := kvValidate(ty, value); err != nil {
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

// MustNewKeyValue is the same as `NewKeyValue` but panics on errors.
func MustNewKeyValue[T KeyValueT](ty Type, key string, value T, parent *KeyValue) *KeyValue {
	kv, err := NewKeyValue(ty, key, value, parent)

	if err != nil {
		panic(err)
	}

	return kv
}

// NewKeyValueEmpty creates an empty KeyValue node (TypeInvalid).
func NewKeyValueEmpty() *KeyValue {
	return MustNewKeyValue(TypeInvalid, "", kvEmptyValue, nil)
}

// NewKeyValueEnd creates an end marker KeyValue node (TypeEnd).
func NewKeyValueEnd() *KeyValue {
	return MustNewKeyValue(TypeEnd, "", kvEmptyValue, nil)
}

// NewKeyValueRoot creates a root KeyValue node (TypeObject).
func NewKeyValueRoot(key string) *KeyValue {
	return NewKeyValueObject(key, nil)
}

// NewKeyValueObject creates a KeyValue node with TypeObject type.
func NewKeyValueObject(key string, parent *KeyValue) *KeyValue {
	return MustNewKeyValue(TypeObject, key, kvEmptyValue, parent)
}

// NewKeyValueString creates a KeyValue node with TypeString type.
func NewKeyValueString(key string, value string, parent *KeyValue) *KeyValue {
	return MustNewKeyValue(TypeString, key, value, parent)
}

// NewKeyValueWString creates a KeyValue node with TypeWString type.
func NewKeyValueWString(key string, value string, parent *KeyValue) *KeyValue {
	return MustNewKeyValue(TypeWString, key, value, parent)
}

// NewKeyValueInt32 creates a KeyValue node with TypeInt32 type.
func NewKeyValueInt32(key string, value int32, parent *KeyValue) *KeyValue {
	return MustNewKeyValue(TypeInt32, key, value, parent)
}

// NewKeyValueColor creates a KeyValue node with TypeColor type.
func NewKeyValueColor(key string, value int32, parent *KeyValue) *KeyValue {
	return MustNewKeyValue(TypeColor, key, value, parent)
}

// NewKeyValuePointer creates a KeyValue node with TypePointer type.
func NewKeyValuePointer(key string, value int32, parent *KeyValue) *KeyValue {
	return MustNewKeyValue(TypePointer, key, value, parent)
}

// NewKeyValueInt64 creates a KeyValue node with TypeInt64 type.
func NewKeyValueInt64(key string, value int64, parent *KeyValue) *KeyValue {
	return MustNewKeyValue(TypeInt64, key, value, parent)
}

// NewKeyValueUint64 creates a KeyValue node with TypeUint64 type.
func NewKeyValueUint64(key string, value uint64, parent *KeyValue) *KeyValue {
	return MustNewKeyValue(TypeUint64, key, value, parent)
}

// NewKeyValueFloat32 creates a KeyValue node with TypeFloat32 type.
func NewKeyValueFloat32(key string, value float32, parent *KeyValue) *KeyValue {
	return MustNewKeyValue(TypeFloat32, key, value, parent)
}

// Type returns the node's type.
func (kv *KeyValue) Type() Type { return kv.ty }

// SetType sets the node's type and returns the node.
//
// If `ty` differs from the node's current type, the node's value is unset.
func (kv *KeyValue) SetType(ty Type) *KeyValue {
	if kv.ty != ty {
		kv.ty = ty
		kv.value = kvEmptyValue
	}

	return kv
}

// Key returns the node's key.
func (kv *KeyValue) Key() string { return kv.key }

// SetKey sets the node's key and returns the node.
func (kv *KeyValue) SetKey(k string) *KeyValue {
	kv.key = k

	return kv
}

// Value returns the node's value.
func (kv *KeyValue) Value() any { return kv.value }

// Format returns the node's value formatted as a string.
//
// It returns an error if formatting fails or if the node holds no value
// (TypeObject, TypeInvalid, TypeEnd).
func (kv *KeyValue) Format() (string, error) {
	if kv.ty == TypeInvalid || kv.ty == TypeObject || kv.ty == TypeEnd {
		return "", kvFormatTypeError(kv.ty)
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

// Map returns the node as map[string]any.
//
// It returns an error if the node's type is not TypeObject.
func (kv *KeyValue) Map() (map[string]any, error) {
	if kv.ty != TypeObject {
		return nil, kvMapTypeError(kv.ty)
	}

	result := make(map[string]any)

	if err := kvMap(kv, result); err != nil {
		return nil, err
	}

	return result, nil
}

// String returns the node's value as string.
//
// It returns an error if the nodes' type is not `TypeString`.
func (kv *KeyValue) String() (string, error) {
	return kvGet[string](kv, TypeString)
}

// Int32 returns the node's value as int32.
//
// It returns an error if the nodes' type is not `TypeInt32`.
func (kv *KeyValue) Int32() (int32, error) {
	return kvGet[int32](kv, TypeInt32)
}

// Color returns the node's value as color.
//
// It returns an error if the nodes' type is not `TypeColor`.
func (kv *KeyValue) Color() (int32, error) {
	return kvGet[int32](kv, TypeColor)
}

// Pointer returns the node's value as pointer.
//
// It returns an error if the nodes' type is not `TypePointer`.
func (kv *KeyValue) Pointer() (int32, error) {
	return kvGet[int32](kv, TypePointer)
}

// Int64 returns the node's value as int64.
//
// It returns an error if the nodes' type is not `TypeInt64`.
func (kv *KeyValue) Int64() (int64, error) {
	return kvGet[int64](kv, TypeInt64)
}

// Uint64 returns the node's value as uint64.
//
// It returns an error if the nodes' type is not `TypeUint64`.
func (kv *KeyValue) Uint64() (uint64, error) {
	return kvGet[uint64](kv, TypeUint64)
}

// Float32 returns the node's value as float32.
//
// It returns an error if the nodes' type is not `TypeFloat32`.
func (kv *KeyValue) Float32() (float32, error) {
	return kvGet[float32](kv, TypeFloat32)
}

// SetString sets the node's type to `TypeString` and sets its value to the given `value`.
func (kv *KeyValue) SetString(value string) *KeyValue {
	return kvSet(kv, TypeString, value)
}

// SetInt32 sets the node's type to `TypeInt32` and sets its value to the given `value`.
func (kv *KeyValue) SetInt32(value int32) *KeyValue {
	return kvSet(kv, TypeInt32, value)
}

// SetColor sets the node's type to `TypeColor` and sets its value to the given `value`.
func (kv *KeyValue) SetColor(value int32) *KeyValue {
	return kvSet(kv, TypeColor, value)
}

// SetPointer sets the node's type to `TypePointer` and sets its value to the given `value`.
func (kv *KeyValue) SetPointer(value int32) *KeyValue {
	return kvSet(kv, TypePointer, value)
}

// SetInt64 sets the node's type to `TypeInt64` and sets its value to the given `value`.
func (kv *KeyValue) SetInt64(value int64) *KeyValue {
	return kvSet(kv, TypeInt64, value)
}

// SetUint64 sets the node's type to `TypeUint64` and sets its value to the given `value`.
func (kv *KeyValue) SetUint64(value uint64) *KeyValue {
	return kvSet(kv, TypeUint64, value)
}

// SetFloat32 sets the node's type to `TypeFloat32` and sets its value to the given `value`.
func (kv *KeyValue) SetFloat32(value float32) *KeyValue {
	return kvSet(kv, TypeFloat32, value)
}

// Parent returns the parent node.
func (kv *KeyValue) Parent() *KeyValue { return kv.parent }

// SetParent sets the node's parent node and returns the node.
func (kv *KeyValue) SetParent(p *KeyValue) *KeyValue {
	kv.parent = p

	return kv
}

// Children returns the node's children.
func (kv *KeyValue) Children() []*KeyValue { return kv.children }

// ResetChildren empties the node's children and returns the node.
func (kv *KeyValue) ResetChildren() *KeyValue {
	kv.children = nil

	return kv
}

// SetChildren sets the node's type to `TypeObject`, sets the node's children and returns the node.
func (kv *KeyValue) SetChildren(children []*KeyValue) *KeyValue {
	kv.ty = TypeObject
	kv.children = children

	for _, c := range children {
		c.SetParent(kv)
	}

	return kv
}

// FindChild finds a child node with the given key.
//
// It returns `nil` if no child node was found.
func (kv *KeyValue) FindChild(key string) *KeyValue {
	for _, c := range kv.children {
		if c.key == key {
			return c
		}
	}

	return nil
}

// NewChild sets the node type to `TypeObject`, creates an empty child node and returns the child.
func (kv *KeyValue) NewChild() *KeyValue {
	return MustNewKeyValue(TypeInvalid, "", kvEmptyValue, kv)
}

// AddChild sets the node type to `TypeObject` and creates and adds a child node and returns the
// node.
func (kv *KeyValue) AddChild(c *KeyValue) *KeyValue {
	kv.ty = TypeObject
	kv.children = append(kv.children, c)

	c.SetParent(kv)

	return kv
}

// AddObject sets the node type to `TypeObject`, adds a TypeObject child node and returns the node.
func (kv *KeyValue) AddObject(key string) *KeyValue {
	NewKeyValueObject(key, kv)

	return kv
}

// AddString sets the node type to `TypeObject`, adds a TypeString child node and returns the node.
func (kv *KeyValue) AddString(key string, value string) *KeyValue {
	NewKeyValueString(key, value, kv)

	return kv
}

// AddString sets the node type to `TypeObject`, adds a TypeWString child node and returns the node.
func (kv *KeyValue) AddWString(key string, value string) *KeyValue {
	NewKeyValueWString(key, value, kv)

	return kv
}

// AddInt32 sets the node type to `TypeObject`, adds a TypeInt32 child node and returns the node.
func (kv *KeyValue) AddInt32(key string, value int32) *KeyValue {
	NewKeyValueInt32(key, value, kv)

	return kv
}

// AddColor sets the node type to `TypeObject`, adds a TypeColor child node and returns the node.
func (kv *KeyValue) AddColor(key string, value int32) *KeyValue {
	NewKeyValueColor(key, value, kv)

	return kv
}

// AddPointer sets the node type to `TypeObject`, adds a TypePointer child node and returns the node.
func (kv *KeyValue) AddPointer(key string, value int32) *KeyValue {
	NewKeyValuePointer(key, value, kv)

	return kv
}

// AddInt64 sets the node type to `TypeObject`, adds a TypeInt64 child node and returns the node.
func (kv *KeyValue) AddInt64(key string, value int64) *KeyValue {
	NewKeyValueInt64(key, value, kv)

	return kv
}

// AddUint64 sets the node type to `TypeObject`, adds a TypeUint64 child node and returns the node.
func (kv *KeyValue) AddUint64(key string, value uint64) *KeyValue {
	NewKeyValueUint64(key, value, kv)

	return kv
}

// AddFloat32 sets the node type to `TypeObject`, adds a TypeFloat32 child node and returns the node.
func (kv *KeyValue) AddFloat32(key string, value float32) *KeyValue {
	NewKeyValueFloat32(key, value, kv)

	return kv
}

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

func kvValidate(ty Type, value any) error { //nolint:cyclop
	switch ty {
	case TypeInvalid, TypeObject, TypeEnd:
		return nil
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

func kvGet[T KeyValueT](kv *KeyValue, ty Type) (T, error) { //nolint:ireturn
	zvalue := *new(T)

	if kv.ty != ty {
		return zvalue, kvGetError(kv, ty)
	}

	value, ok := kv.value.(T)

	if !ok {
		return zvalue, kvGetError(kv, ty)
	}

	return value, nil
}

func kvSet[T KeyValueT](kv *KeyValue, ty Type, value T) *KeyValue {
	kv.ty = ty
	kv.value = value

	return kv
}

func kvMap(kv *KeyValue, m map[string]any) error {
	switch kv.ty {
	case TypeInvalid:
		return kvMapTypeError(kv.ty)
	case TypeEnd:
		return nil
	case TypeObject:
		{
			children := make(map[string]any, len(kv.children))
			m[kv.key] = children

			for _, child := range kv.children {
				if err := kvMap(child, children); err != nil {
					return err
				}
			}
		}
	case TypeString, TypeWString, TypeInt32, TypeColor, TypePointer, TypeInt64, TypeUint64, TypeFloat32:
		{
			m[kv.key] = kv.value
		}
	}

	return nil
}
