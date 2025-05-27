package steamlib

import (
	"fmt"
	"os"
	"strconv"

	"github.com/13k/valve.go/kv1"
	"github.com/go-viper/mapstructure/v2"
)

func kv1ParseFile(path string) (*kv1.KeyValue, error) {
	f, err := os.Open(path)

	if err != nil {
		return nil, fmt.Errorf("failed to open %q: %w", path, err)
	}

	defer f.Close()

	dec := kv1.NewTextDecoder(f)
	kv := kv1.NewKeyValueEmpty()

	if err := dec.Decode(kv); err != nil {
		return nil, fmt.Errorf("failed to parse kv1 data at %q: %w", path, err)
	}

	return kv, nil
}

func kv1UnmarshalFile[T any](path string) (*T, error) {
	kv, err := kv1ParseFile(path)

	if err != nil {
		return nil, err
	}

	kvMap, err := kv.Map()

	if err != nil {
		return nil, fmt.Errorf("failed to convert KeyValue to map: %w", err)
	}

	output := new(T)

	if err := mapstructure.Decode(kvMap, output); err != nil {
		return nil, fmt.Errorf("failed to decode kv1 struct [%T] from %q: %w", output, path, err)
	}

	return output, nil
}

func kv1Array[T any](value map[string]T) ([]T, error) {
	result := make([]T, len(value))

	for idxStr, item := range value {
		idx, err := strconv.ParseUint(idxStr, 10, 64)

		if err != nil {
			return nil, fmt.Errorf("invalid kv1 array index %q: %w", idxStr, err)
		}

		result[idx] = item
	}

	return result, nil
}
