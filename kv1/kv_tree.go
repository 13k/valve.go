package kv1

// Parent returns the parent node.
func (kv *KeyValue) Parent() *KeyValue { return kv.parent }

// SetParent sets the parent node and returns itself.
func (kv *KeyValue) SetParent(p *KeyValue) *KeyValue {
	kv.parent = p

	return kv
}

// Children returns child nodes.
func (kv *KeyValue) Children() []*KeyValue { return kv.children }

// ResetChildren empties child nodes and returns itself.
func (kv *KeyValue) ResetChildren() *KeyValue {
	kv.children = nil

	return kv
}

// SetChildren sets the type to `TypeObject`, sets the children and returns itself.
func (kv *KeyValue) SetChildren(children []*KeyValue) *KeyValue {
	kv.ty = TypeObject
	kv.children = children

	for _, c := range children {
		c.SetParent(kv)
	}

	return kv
}

// FindChild finds a child node with the given `key`.
//
// It returns `nil` if none was found.
func (kv *KeyValue) FindChild(key string) *KeyValue {
	for _, c := range kv.children {
		if c.key == key {
			return c
		}
	}

	return nil
}

// NewChild sets type to `TypeObject`, adds an empty child node and returns the child.
func (kv *KeyValue) NewChild() *KeyValue {
	return NewKeyValueEmptyChild(kv)
}

// AddChild sets type to `TypeObject` and adds a child node and returns the node.
func (kv *KeyValue) AddChild(c *KeyValue) *KeyValue {
	kv.ty = TypeObject
	kv.children = append(kv.children, c)

	c.SetParent(kv)

	return kv
}

// AddObject sets type to `TypeObject`, adds a `TypeObject` child node and returns itself.
func (kv *KeyValue) AddObject(key string) *KeyValue {
	NewKeyValueObject(key, kv)

	return kv
}

// AddString sets type to `TypeObject`, adds a `TypeString` child node and returns itself.
func (kv *KeyValue) AddString(key string, value string) *KeyValue {
	NewKeyValueString(key, value, kv)

	return kv
}

// AddString sets type to `TypeObject`, adds a TypeWString child node and returns itself.
func (kv *KeyValue) AddWString(key string, value string) *KeyValue {
	NewKeyValueWString(key, value, kv)

	return kv
}

// AddInt32 sets type to `TypeObject`, adds a TypeInt32 child node and returns itself.
func (kv *KeyValue) AddInt32(key string, value int32) *KeyValue {
	NewKeyValueInt32(key, value, kv)

	return kv
}

// AddColor sets type to `TypeObject`, adds a TypeColor child node and returns itself.
func (kv *KeyValue) AddColor(key string, value int32) *KeyValue {
	NewKeyValueColor(key, value, kv)

	return kv
}

// AddPointer sets type to `TypeObject`, adds a TypePointer child node and returns itself.
func (kv *KeyValue) AddPointer(key string, value int32) *KeyValue {
	NewKeyValuePointer(key, value, kv)

	return kv
}

// AddInt64 sets type to `TypeObject`, adds a TypeInt64 child node and returns itself.
func (kv *KeyValue) AddInt64(key string, value int64) *KeyValue {
	NewKeyValueInt64(key, value, kv)

	return kv
}

// AddUint64 sets type to `TypeObject`, adds a TypeUint64 child node and returns itself.
func (kv *KeyValue) AddUint64(key string, value uint64) *KeyValue {
	NewKeyValueUint64(key, value, kv)

	return kv
}

// AddFloat32 sets type to `TypeObject`, adds a TypeFloat32 child node and returns itself.
func (kv *KeyValue) AddFloat32(key string, value float32) *KeyValue {
	NewKeyValueFloat32(key, value, kv)

	return kv
}
