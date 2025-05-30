package kv1_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/13k/valve.go/kv1"
)

func TestKeyValue(t *testing.T) {
	t.Parallel()

	suite.Run(t, new(KeyValueSuite))
}

type KeyValueSuite struct {
	Suite
}

func (s *KeyValueSuite) TestType() {
	var kv *kv1.KeyValue

	require := s.Require()

	{
		kv = kv1.NewKeyValueObject("Object", nil)

		require.Equal(kv1.TypeObject, kv.Type())
		require.True(kv.IsObject())
		require.False(kv.IsField())
	}

	{
		kv = kv1.NewKeyValueString("String", "String", nil)

		require.Equal(kv1.TypeString, kv.Type())
		require.True(kv.IsString())
		require.True(kv.IsField())
	}

	{
		kv = kv1.NewKeyValueWString("WString", "WString", nil)

		require.Equal(kv1.TypeWString, kv.Type())
		require.True(kv.IsWString())
		require.True(kv.IsField())
	}

	{
		kv = kv1.NewKeyValueInt32("Int32", 1, nil)

		require.Equal(kv1.TypeInt32, kv.Type())
		require.True(kv.IsInt32())
		require.True(kv.IsField())
	}

	{
		kv = kv1.NewKeyValueColor("Color", 1, nil)

		require.Equal(kv1.TypeColor, kv.Type())
		require.True(kv.IsColor())
		require.True(kv.IsField())
	}

	{
		kv = kv1.NewKeyValuePointer("Pointer", 1, nil)

		require.Equal(kv1.TypePointer, kv.Type())
		require.True(kv.IsPointer())
		require.True(kv.IsField())
	}

	{
		kv = kv1.NewKeyValueInt64("Int64", 1, nil)

		require.Equal(kv1.TypeInt64, kv.Type())
		require.True(kv.IsInt64())
		require.True(kv.IsField())
	}

	{
		kv = kv1.NewKeyValueUint64("Uint64", 1, nil)

		require.Equal(kv1.TypeUint64, kv.Type())
		require.True(kv.IsUint64())
		require.True(kv.IsField())
	}

	{
		kv = kv1.NewKeyValueFloat32("Float32", 1, nil)

		require.Equal(kv1.TypeFloat32, kv.Type())
		require.True(kv.IsFloat32())
		require.True(kv.IsField())
	}

	{
		kv = kv1.NewKeyValueEnd()

		require.Equal(kv1.TypeEnd, kv.Type())
		require.True(kv.IsEnd())
		require.False(kv.IsField())
	}

	{
		kv = kv1.NewKeyValueEmpty()

		require.Equal(kv1.TypeInvalid, kv.Type())
		require.True(kv.IsInvalid())
		require.False(kv.IsField())
	}
}
