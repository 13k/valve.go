package kv1

// Map returns the KeyValue as `map[string]any`.
//
// It returns an error if the KeyValue's type is not `TypeObject`.
func (kv *KeyValue) Map() (map[string]any, error) {
	if !kv.IsObject() {
		return nil, kvMapTypeError(kv)
	}

	result := make(map[string]any)

	if err := kvMap(kv, result); err != nil {
		return nil, err
	}

	return result, nil
}

func kvMap(kv *KeyValue, m map[string]any) error {
	switch {
	case kv.IsInvalid():
		return kvMapTypeError(kv)
	case kv.IsEnd():
		return nil
	case kv.IsObject():
		{
			children := make(map[string]any, len(kv.children))
			m[kv.key] = children

			for _, child := range kv.children {
				if err := kvMap(child, children); err != nil {
					return err
				}
			}
		}
	default:
		{
			m[kv.key] = kv.value
		}
	}

	return nil
}
