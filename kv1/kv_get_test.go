package kv1_test

import (
	"strings"

	"github.com/13k/valve.go/kv1"
)

type kvTestCaseGet struct {
	Name     string
	Subject  *kv1.KeyValue
	Expected any
	Err      string
}

func (s *KeyValueSuite) TestGet() {
	testCases := []kvTestCaseGet{
		// TypeString

		{
			Name:     "String/ok",
			Subject:  kv1.NewKeyValueString("foo", "bar", nil),
			Expected: "bar",
		},
		{
			Name:    "String/err",
			Subject: kv1.NewKeyValueInt32("foo", 1, nil),
			Err:     `type Int32 as String`,
		},

		// TypeWString

		{
			Name:     "WString/ok",
			Subject:  kv1.NewKeyValueWString("foo", "bar", nil),
			Expected: "bar",
		},
		{
			Name:    "WString/err",
			Subject: kv1.NewKeyValueInt32("foo", 1, nil),
			Err:     `type Int32 as WString`,
		},

		// TypeInt32

		{
			Name:     "Int32/ok",
			Subject:  kv1.NewKeyValueInt32("foo", 1, nil),
			Expected: int32(1),
		},
		{
			Name:    "Int32/err",
			Subject: kv1.NewKeyValueString("foo", "bar", nil),
			Err:     `type String as Int32`,
		},

		// TypeColor

		{
			Name:     "Color/ok",
			Subject:  kv1.NewKeyValueColor("foo", 1, nil),
			Expected: int32(1),
		},
		{
			Name:    "Color/err",
			Subject: kv1.NewKeyValueString("foo", "bar", nil),
			Err:     `type String as Color`,
		},

		// TypePointer

		{
			Name:     "Pointer/ok",
			Subject:  kv1.NewKeyValuePointer("foo", 1, nil),
			Expected: int32(1),
		},
		{
			Name:    "Pointer/err",
			Subject: kv1.NewKeyValueString("foo", "bar", nil),
			Err:     `type String as Pointer`,
		},

		// TypeInt64

		{
			Name:     "Int64/ok",
			Subject:  kv1.NewKeyValueInt64("foo", 1, nil),
			Expected: int64(1),
		},
		{
			Name:    "Int64/err",
			Subject: kv1.NewKeyValueString("foo", "bar", nil),
			Err:     `type String as Int64`,
		},

		// TypeUint64

		{
			Name:     "Uint64/ok",
			Subject:  kv1.NewKeyValueUint64("foo", 1, nil),
			Expected: uint64(1),
		},
		{
			Name:    "Uint64/err",
			Subject: kv1.NewKeyValueString("foo", "bar", nil),
			Err:     `type String as Uint64`,
		},

		// TypeFloat32

		{
			Name:     "Float32/ok",
			Subject:  kv1.NewKeyValueFloat32("foo", 1, nil),
			Expected: float32(1),
		},
		{
			Name:    "Float32/err",
			Subject: kv1.NewKeyValueString("foo", "bar", nil),
			Err:     `type String as Float32`,
		},
	}

	for _, testCase := range testCases {
		s.testGet(testCase)
	}
}

func (s *KeyValueSuite) testGet(testCase kvTestCaseGet) {
	require := s.Require()
	testFuncName, _, ok := strings.Cut(testCase.Name, "/")

	if !ok {
		require.FailNowf("invalid test case", "%#v", testCase)
	}

	s.Run(testCase.Name, func() {
		var (
			actual any
			err    error
		)

		subject := testCase.Subject

		switch testFuncName {
		case "String":
			actual, err = subject.String()
		case "WString":
			actual, err = subject.WString()
		case "Int32":
			actual, err = subject.Int32()
		case "Color":
			actual, err = subject.Color()
		case "Pointer":
			actual, err = subject.Pointer()
		case "Int64":
			actual, err = subject.Int64()
		case "Uint64":
			actual, err = subject.Uint64()
		case "Float32":
			actual, err = subject.Float32()
		default:
			require.FailNowf("invalid test case", "%#v", testCase)
		}

		if testCase.Err == "" {
			require.NoError(err)
			require.Equal(testCase.Expected, actual)
		} else {
			require.ErrorContains(err, testCase.Err)
		}
	})
}
