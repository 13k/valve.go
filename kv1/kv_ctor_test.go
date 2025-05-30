package kv1_test

import (
	"github.com/13k/valve.go/kv1"
)

type kvTestCaseCtor struct {
	Name   string
	Type   kv1.Type
	Key    string
	Value  any
	Parent *kv1.KeyValue
	Err    string
}

func (s *KeyValueSuite) TestNewKeyValue() { //nolint:maintidx
	testCases := []kvTestCaseCtor{
		// TypeInvalid

		{
			Name:   "TypeInvalid/ok",
			Type:   kv1.TypeInvalid,
			Key:    "",
			Value:  nil,
			Parent: nil,
		},
		{
			Name:   "TypeInvalid/key_err",
			Type:   kv1.TypeInvalid,
			Key:    "foo",
			Value:  nil,
			Parent: nil,
			Err:    `invalid key`,
		},
		{
			Name:   "TypeInvalid/value_err",
			Type:   kv1.TypeInvalid,
			Key:    "",
			Value:  1,
			Parent: nil,
			Err:    `type Invalid with value 1 (int)`,
		},
		{
			Name:   "TypeInvalid/child",
			Type:   kv1.TypeInvalid,
			Key:    "",
			Value:  nil,
			Parent: kv1.NewKeyValueObjectRoot("root"),
		},

		// TypeEnd

		{
			Name:   "TypeEnd/ok",
			Type:   kv1.TypeEnd,
			Key:    "",
			Value:  nil,
			Parent: nil,
		},
		{
			Name:   "TypeEnd/key_err",
			Type:   kv1.TypeEnd,
			Key:    "foo",
			Value:  nil,
			Parent: nil,
			Err:    `invalid key`,
		},
		{
			Name:   "TypeEnd/value_err",
			Type:   kv1.TypeEnd,
			Key:    "",
			Value:  1,
			Parent: nil,
			Err:    `type End with value 1 (int)`,
		},
		{
			Name:   "TypeEnd/child",
			Type:   kv1.TypeEnd,
			Key:    "",
			Value:  nil,
			Parent: kv1.NewKeyValueObjectRoot("root"),
		},

		// TypeObject

		{
			Name:   "TypeObject/ok",
			Type:   kv1.TypeObject,
			Key:    "foo",
			Value:  nil,
			Parent: nil,
		},
		{
			Name:   "TypeObject/key_err",
			Type:   kv1.TypeObject,
			Key:    "",
			Value:  nil,
			Parent: nil,
			Err:    `invalid key`,
		},
		{
			Name:   "TypeObject/value_err",
			Type:   kv1.TypeObject,
			Key:    "foo",
			Value:  1,
			Parent: nil,
			Err:    `type Object with value 1 (int)`,
		},
		{
			Name:   "TypeObject/child",
			Type:   kv1.TypeObject,
			Key:    "foo",
			Value:  nil,
			Parent: kv1.NewKeyValueObjectRoot("root"),
		},

		// TypeString

		{
			Name:   "TypeString/ok",
			Type:   kv1.TypeString,
			Key:    "foo",
			Value:  "bar",
			Parent: nil,
		},
		{
			Name:   "TypeString/key_err",
			Type:   kv1.TypeString,
			Key:    "",
			Value:  "foo",
			Parent: nil,
			Err:    `invalid key`,
		},
		{
			Name:   "TypeString/value_err",
			Type:   kv1.TypeString,
			Key:    "foo",
			Value:  1,
			Parent: nil,
			Err:    `type String with value 1 (int)`,
		},
		{
			Name:   "TypeString/child",
			Type:   kv1.TypeString,
			Key:    "foo",
			Value:  "bar",
			Parent: kv1.NewKeyValueObjectRoot("root"),
		},

		// TypeWString

		{
			Name:   "TypeWString/ok",
			Type:   kv1.TypeWString,
			Key:    "foo",
			Value:  "bar",
			Parent: nil,
		},
		{
			Name:   "TypeWString/key_err",
			Type:   kv1.TypeWString,
			Key:    "",
			Value:  "foo",
			Parent: nil,
			Err:    `invalid key`,
		},
		{
			Name:   "TypeWString/value_err",
			Type:   kv1.TypeWString,
			Key:    "foo",
			Value:  1,
			Parent: nil,
			Err:    `type WString with value 1 (int)`,
		},
		{
			Name:   "TypeWString/child",
			Type:   kv1.TypeWString,
			Key:    "foo",
			Value:  "bar",
			Parent: kv1.NewKeyValueObjectRoot("root"),
		},

		// TypeInt32

		{
			Name:   "TypeInt32/ok",
			Type:   kv1.TypeInt32,
			Key:    "foo",
			Value:  int32(1),
			Parent: nil,
		},
		{
			Name:   "TypeInt32/key_err",
			Type:   kv1.TypeInt32,
			Key:    "",
			Value:  int32(1),
			Parent: nil,
			Err:    `invalid key`,
		},
		{
			Name:   "TypeInt32/value_err",
			Type:   kv1.TypeInt32,
			Key:    "foo",
			Value:  "1",
			Parent: nil,
			Err:    `type Int32 with value 1 (string)`,
		},
		{
			Name:   "TypeInt32/child",
			Type:   kv1.TypeInt32,
			Key:    "foo",
			Value:  int32(1),
			Parent: kv1.NewKeyValueObjectRoot("root"),
		},

		// TypeColor

		{
			Name:   "TypeColor/ok",
			Type:   kv1.TypeColor,
			Key:    "foo",
			Value:  int32(1),
			Parent: nil,
		},
		{
			Name:   "TypeColor/key_err",
			Type:   kv1.TypeColor,
			Key:    "",
			Value:  int32(1),
			Parent: nil,
			Err:    `invalid key`,
		},
		{
			Name:   "TypeColor/value_err",
			Type:   kv1.TypeColor,
			Key:    "foo",
			Value:  "1",
			Parent: nil,
			Err:    `type Color with value 1 (string)`,
		},
		{
			Name:   "TypeColor/child",
			Type:   kv1.TypeColor,
			Key:    "foo",
			Value:  int32(1),
			Parent: kv1.NewKeyValueObjectRoot("root"),
		},

		// TypePointer

		{
			Name:   "TypePointer/ok",
			Type:   kv1.TypePointer,
			Key:    "foo",
			Value:  int32(1),
			Parent: nil,
		},
		{
			Name:   "TypePointer/key_err",
			Type:   kv1.TypePointer,
			Key:    "",
			Value:  int32(1),
			Parent: nil,
			Err:    `invalid key`,
		},
		{
			Name:   "TypePointer/value_err",
			Type:   kv1.TypePointer,
			Key:    "foo",
			Value:  "1",
			Parent: nil,
			Err:    `type Pointer with value 1 (string)`,
		},
		{
			Name:   "TypePointer/child",
			Type:   kv1.TypePointer,
			Key:    "foo",
			Value:  int32(1),
			Parent: kv1.NewKeyValueObjectRoot("root"),
		},

		// TypeInt64

		{
			Name:   "TypeInt64/ok",
			Type:   kv1.TypeInt64,
			Key:    "foo",
			Value:  int64(1),
			Parent: nil,
		},
		{
			Name:   "TypeInt64/key_err",
			Type:   kv1.TypeInt64,
			Key:    "",
			Value:  int64(1),
			Parent: nil,
			Err:    `invalid key`,
		},
		{
			Name:   "TypeInt64/value_err",
			Type:   kv1.TypeInt64,
			Key:    "foo",
			Value:  "1",
			Parent: nil,
			Err:    `type Int64 with value 1 (string)`,
		},
		{
			Name:   "TypeInt64/child",
			Type:   kv1.TypeInt64,
			Key:    "foo",
			Value:  int64(1),
			Parent: kv1.NewKeyValueObjectRoot("root"),
		},

		// TypeUint64

		{
			Name:   "TypeUint64/ok",
			Type:   kv1.TypeUint64,
			Key:    "foo",
			Value:  uint64(1),
			Parent: nil,
		},
		{
			Name:   "TypeUint64/key_err",
			Type:   kv1.TypeUint64,
			Key:    "",
			Value:  uint64(1),
			Parent: nil,
			Err:    `invalid key`,
		},
		{
			Name:   "TypeUint64/value_err",
			Type:   kv1.TypeUint64,
			Key:    "foo",
			Value:  "1",
			Parent: nil,
			Err:    `type Uint64 with value 1 (string)`,
		},
		{
			Name:   "TypeUint64/child",
			Type:   kv1.TypeUint64,
			Key:    "foo",
			Value:  uint64(1),
			Parent: kv1.NewKeyValueObjectRoot("root"),
		},

		// TypeFloat32

		{
			Name:   "TypeFloat32/ok",
			Type:   kv1.TypeFloat32,
			Key:    "foo",
			Value:  float32(1),
			Parent: nil,
		},
		{
			Name:   "TypeFloat32/key_err",
			Type:   kv1.TypeFloat32,
			Key:    "",
			Value:  float32(1),
			Parent: nil,
			Err:    `invalid key`,
		},
		{
			Name:   "TypeFloat32/value_err",
			Type:   kv1.TypeFloat32,
			Key:    "foo",
			Value:  "1",
			Parent: nil,
			Err:    `type Float32 with value 1 (string)`,
		},
		{
			Name:   "TypeFloat32/child",
			Type:   kv1.TypeFloat32,
			Key:    "foo",
			Value:  float32(1),
			Parent: kv1.NewKeyValueObjectRoot("root"),
		},
	}

	for _, testCase := range testCases {
		s.testCtor(testCase)
	}
}

func (s *KeyValueSuite) testCtor(testCase kvTestCaseCtor) {
	s.Run(testCase.Name, func() {
		require := s.Require()

		actual, err := kv1.NewKeyValue(
			testCase.Type,
			testCase.Key,
			testCase.Value,
			testCase.Parent,
		)

		if testCase.Err == "" {
			require.NoError(err)

			require.Equal(testCase.Type, actual.Type())
			require.Equal(testCase.Key, actual.Key())
			require.Equal(testCase.Value, actual.Value())
			require.Same(testCase.Parent, actual.Parent())
		} else {
			require.ErrorContains(err, testCase.Err)
		}
	})
}
