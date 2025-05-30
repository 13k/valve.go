package kv1_test

import (
	"math"
	"strings"

	"github.com/13k/valve.go/kv1"
)

type kvTestCaseConvert struct {
	Name     string
	Subject  *kv1.KeyValue
	Expected any
	Err      string
}

func (s *KeyValueSuite) TestToString() {
	testCases := []kvTestCaseConvert{
		{
			Name:    "ToString/invalid",
			Subject: kv1.NewKeyValueEmpty(),
			Err:     `cannot format type Invalid`,
		},
		{
			Name:    "ToString/end",
			Subject: kv1.NewKeyValueEnd(),
			Err:     `cannot format type End`,
		},
		{
			Name:    "ToString/object",
			Subject: kv1.NewKeyValueObject("object", nil),
			Err:     `cannot format type Object`,
		},
		{
			Name:     "ToString/string",
			Subject:  kv1.NewKeyValueString("string", "foo", nil),
			Expected: "foo",
		},
		{
			Name:     "ToString/wstring",
			Subject:  kv1.NewKeyValueWString("wstring", "foo", nil),
			Expected: "foo",
		},
		{
			Name:     "ToString/int32",
			Subject:  kv1.NewKeyValueInt32("int32", -13, nil),
			Expected: "-13",
		},
		{
			Name:     "ToString/color",
			Subject:  kv1.NewKeyValueColor("color", -13, nil),
			Expected: "-13",
		},
		{
			Name:     "ToString/pointer",
			Subject:  kv1.NewKeyValuePointer("pointer", 13, nil),
			Expected: "13",
		},
		{
			Name:     "ToString/int64",
			Subject:  kv1.NewKeyValueInt64("int64", 13, nil),
			Expected: "13",
		},
		{
			Name:     "ToString/uint64",
			Subject:  kv1.NewKeyValueUint64("uint64", 13, nil),
			Expected: "13",
		},
		{
			Name:     "ToString/float32",
			Subject:  kv1.NewKeyValueFloat32("float32", 13.31, nil),
			Expected: "13.31",
		},
	}

	for _, testCase := range testCases {
		s.testConvert(testCase)
	}
}

func (s *KeyValueSuite) TestToInt32() {
	testCases := []kvTestCaseConvert{
		{
			Name:    "ToInt32/invalid",
			Subject: kv1.NewKeyValueEmpty(),
			Err:     `type Invalid to int32`,
		},
		{
			Name:    "ToInt32/end",
			Subject: kv1.NewKeyValueEnd(),
			Err:     `type End to int32`,
		},
		{
			Name:    "ToInt32/object",
			Subject: kv1.NewKeyValueObject("object", nil),
			Err:     `type Object to int32`,
		},
		{
			Name:    "ToInt32/string_err",
			Subject: kv1.NewKeyValueString("string", "foo", nil),
			Err:     `failed to convert`,
		},
		{
			Name:     "ToInt32/string_ok",
			Subject:  kv1.NewKeyValueString("string", "13", nil),
			Expected: int32(13),
		},
		{
			Name:    "ToInt32/wstring_err",
			Subject: kv1.NewKeyValueWString("wstring", "foo", nil),
			Err:     `failed to convert`,
		},
		{
			Name:     "ToInt32/wstring_ok",
			Subject:  kv1.NewKeyValueWString("wstring", "13", nil),
			Expected: int32(13),
		},
		{
			Name:     "ToInt32/int32",
			Subject:  kv1.NewKeyValueInt32("int32", -13, nil),
			Expected: int32(-13),
		},
		{
			Name:     "ToInt32/color",
			Subject:  kv1.NewKeyValueColor("color", -13, nil),
			Expected: int32(-13),
		},
		{
			Name:     "ToInt32/pointer",
			Subject:  kv1.NewKeyValuePointer("pointer", 13, nil),
			Expected: int32(13),
		},
		{
			Name:     "ToInt32/int64_ok",
			Subject:  kv1.NewKeyValueInt64("int64", 13, nil),
			Expected: int32(13),
		},
		{
			Name:    "ToInt32/int64_overflow",
			Subject: kv1.NewKeyValueInt64("int64", math.MaxInt32+1, nil),
			Err:     `failed to convert`,
		},
		{
			Name:     "ToInt32/uint64_ok",
			Subject:  kv1.NewKeyValueUint64("uint64", 13, nil),
			Expected: int32(13),
		},
		{
			Name:    "ToInt32/uint64_overflow",
			Subject: kv1.NewKeyValueUint64("uint64", math.MaxInt32+1, nil),
			Err:     `failed to convert`,
		},
		{
			Name:     "ToInt32/float32_ok",
			Subject:  kv1.NewKeyValueFloat32("float32", 13.31, nil),
			Expected: int32(13),
		},
		{
			Name:    "ToInt32/float32_overflow",
			Subject: kv1.NewKeyValueFloat32("float32", math.MaxInt32+1, nil),
			Err:     `failed to convert`,
		},
	}

	for _, testCase := range testCases {
		s.testConvert(testCase)
	}
}

func (s *KeyValueSuite) TestToInt64() {
	testCases := []kvTestCaseConvert{
		{
			Name:    "ToInt64/invalid",
			Subject: kv1.NewKeyValueEmpty(),
			Err:     `type Invalid to int64`,
		},
		{
			Name:    "ToInt64/end",
			Subject: kv1.NewKeyValueEnd(),
			Err:     `type End to int64`,
		},
		{
			Name:    "ToInt64/object",
			Subject: kv1.NewKeyValueObject("object", nil),
			Err:     `type Object to int64`,
		},
		{
			Name:    "ToInt64/string_err",
			Subject: kv1.NewKeyValueString("string", "foo", nil),
			Err:     `failed to convert`,
		},
		{
			Name:     "ToInt64/string_ok",
			Subject:  kv1.NewKeyValueString("string", "13", nil),
			Expected: int64(13),
		},
		{
			Name:    "ToInt64/wstring_err",
			Subject: kv1.NewKeyValueWString("wstring", "foo", nil),
			Err:     `failed to convert`,
		},
		{
			Name:     "ToInt64/wstring_ok",
			Subject:  kv1.NewKeyValueWString("wstring", "13", nil),
			Expected: int64(13),
		},
		{
			Name:     "ToInt64/int32",
			Subject:  kv1.NewKeyValueInt32("int32", -13, nil),
			Expected: int64(-13),
		},
		{
			Name:     "ToInt64/color",
			Subject:  kv1.NewKeyValueColor("color", -13, nil),
			Expected: int64(-13),
		},
		{
			Name:     "ToInt64/pointer",
			Subject:  kv1.NewKeyValuePointer("pointer", 13, nil),
			Expected: int64(13),
		},
		{
			Name:     "ToInt64/int64",
			Subject:  kv1.NewKeyValueInt64("int64", 13, nil),
			Expected: int64(13),
		},
		{
			Name:     "ToInt64/uint64_ok",
			Subject:  kv1.NewKeyValueUint64("uint64", 13, nil),
			Expected: int64(13),
		},
		{
			Name:    "ToInt64/uint64_overflow",
			Subject: kv1.NewKeyValueUint64("uint64", math.MaxInt64+1, nil),
			Err:     `failed to convert`,
		},
		{
			Name:     "ToInt64/float32_ok",
			Subject:  kv1.NewKeyValueFloat32("float32", 13.31, nil),
			Expected: int64(13),
		},
		{
			Name:    "ToInt64/float32_overflow",
			Subject: kv1.NewKeyValueFloat32("float32", math.MaxInt64+1, nil),
			Err:     `failed to convert`,
		},
	}

	for _, testCase := range testCases {
		s.testConvert(testCase)
	}
}

func (s *KeyValueSuite) TestToUint64() {
	testCases := []kvTestCaseConvert{
		{
			Name:    "ToUint64/invalid",
			Subject: kv1.NewKeyValueEmpty(),
			Err:     `type Invalid to uint64`,
		},
		{
			Name:    "ToUint64/end",
			Subject: kv1.NewKeyValueEnd(),
			Err:     `type End to uint64`,
		},
		{
			Name:    "ToUint64/object",
			Subject: kv1.NewKeyValueObject("object", nil),
			Err:     `type Object to uint64`,
		},
		{
			Name:    "ToUint64/string_err",
			Subject: kv1.NewKeyValueString("string", "foo", nil),
			Err:     `failed to convert`,
		},
		{
			Name:    "ToUint64/string_neg",
			Subject: kv1.NewKeyValueString("string", "-13", nil),
			Err:     `failed to convert`,
		},
		{
			Name:     "ToUint64/string_ok",
			Subject:  kv1.NewKeyValueString("string", "13", nil),
			Expected: uint64(13),
		},
		{
			Name:    "ToUint64/wstring_err",
			Subject: kv1.NewKeyValueWString("wstring", "foo", nil),
			Err:     `failed to convert`,
		},
		{
			Name:    "ToUint64/wstring_neg",
			Subject: kv1.NewKeyValueWString("wstring", "-13", nil),
			Err:     `failed to convert`,
		},
		{
			Name:     "ToUint64/wstring_ok",
			Subject:  kv1.NewKeyValueWString("wstring", "13", nil),
			Expected: uint64(13),
		},
		{
			Name:     "ToUint64/int32_ok",
			Subject:  kv1.NewKeyValueInt32("int32", 13, nil),
			Expected: uint64(13),
		},
		{
			Name:    "ToUint64/int32_neg",
			Subject: kv1.NewKeyValueInt32("int32", -13, nil),
			Err:     `failed to convert`,
		},
		{
			Name:     "ToUint64/color",
			Subject:  kv1.NewKeyValueColor("color", 13, nil),
			Expected: uint64(13),
		},
		{
			Name:     "ToUint64/pointer",
			Subject:  kv1.NewKeyValuePointer("pointer", 13, nil),
			Expected: uint64(13),
		},
		{
			Name:     "ToUint64/int64_ok",
			Subject:  kv1.NewKeyValueInt64("int64", 13, nil),
			Expected: uint64(13),
		},
		{
			Name:    "ToUint64/int64_neg",
			Subject: kv1.NewKeyValueInt64("int64", -13, nil),
			Err:     `failed to convert`,
		},
		{
			Name:     "ToUint64/uint64",
			Subject:  kv1.NewKeyValueUint64("uint64", 13, nil),
			Expected: uint64(13),
		},
		{
			Name:     "ToUint64/float32_ok",
			Subject:  kv1.NewKeyValueFloat32("float32", 13.31, nil),
			Expected: uint64(13),
		},
		{
			Name:    "ToUint64/float32_overflow",
			Subject: kv1.NewKeyValueFloat32("float32", math.MaxUint64+1, nil),
			Err:     `failed to convert`,
		},
	}

	for _, testCase := range testCases {
		s.testConvert(testCase)
	}
}

func (s *KeyValueSuite) TestToFloat32() {
	testCases := []kvTestCaseConvert{
		{
			Name:    "ToFloat32/invalid",
			Subject: kv1.NewKeyValueEmpty(),
			Err:     `type Invalid to float32`,
		},
		{
			Name:    "ToFloat32/end",
			Subject: kv1.NewKeyValueEnd(),
			Err:     `type End to float32`,
		},
		{
			Name:    "ToFloat32/object",
			Subject: kv1.NewKeyValueObject("object", nil),
			Err:     `type Object to float32`,
		},
		{
			Name:    "ToFloat32/string_err",
			Subject: kv1.NewKeyValueString("string", "foo", nil),
			Err:     `failed to convert`,
		},
		{
			Name:     "ToFloat32/string_neg",
			Subject:  kv1.NewKeyValueString("string", "-13.31", nil),
			Expected: float32(-13.31),
		},
		{
			Name:     "ToFloat32/string_pos",
			Subject:  kv1.NewKeyValueString("string", "13.31", nil),
			Expected: float32(13.31),
		},
		{
			Name:    "ToFloat32/wstring_err",
			Subject: kv1.NewKeyValueWString("wstring", "foo", nil),
			Err:     `failed to convert`,
		},
		{
			Name:     "ToFloat32/wstring_neg",
			Subject:  kv1.NewKeyValueWString("wstring", "-13.31", nil),
			Expected: float32(-13.31),
		},
		{
			Name:     "ToFloat32/wstring_pos",
			Subject:  kv1.NewKeyValueWString("wstring", "13.31", nil),
			Expected: float32(13.31),
		},
		{
			Name:     "ToFloat32/int32_pos",
			Subject:  kv1.NewKeyValueInt32("int32", 13, nil),
			Expected: float32(13),
		},
		{
			Name:     "ToFloat32/int32_neg",
			Subject:  kv1.NewKeyValueInt32("int32", -13, nil),
			Expected: float32(-13),
		},
		{
			Name:     "ToFloat32/color",
			Subject:  kv1.NewKeyValueColor("color", 13, nil),
			Expected: float32(13),
		},
		{
			Name:     "ToFloat32/pointer",
			Subject:  kv1.NewKeyValuePointer("pointer", 13, nil),
			Expected: float32(13),
		},
		{
			Name:     "ToFloat32/int64_pos",
			Subject:  kv1.NewKeyValueInt64("int64", 13, nil),
			Expected: float32(13),
		},
		{
			Name:     "ToFloat32/int64_neg",
			Subject:  kv1.NewKeyValueInt64("int64", -13, nil),
			Expected: float32(-13),
		},
		{
			Name:     "ToFloat32/uint64",
			Subject:  kv1.NewKeyValueUint64("uint64", 13, nil),
			Expected: float32(13),
		},
		{
			Name:     "ToFloat32/float32",
			Subject:  kv1.NewKeyValueFloat32("float32", 13.31, nil),
			Expected: float32(13.31),
		},
	}

	for _, testCase := range testCases {
		s.testConvert(testCase)
	}
}

func (s *KeyValueSuite) testConvert(testCase kvTestCaseConvert) {
	require := s.Require()
	testFuncName, testName, ok := strings.Cut(testCase.Name, "/")

	if !ok {
		require.FailNowf("invalid test case", "%#v", testCase)
	}

	s.Run(testName, func() {
		var (
			actual any
			err    error
		)

		subject := testCase.Subject

		switch testFuncName {
		case "ToString":
			actual, err = subject.ToString()
		case "ToInt32":
			actual, err = subject.ToInt32()
		case "ToInt64":
			actual, err = subject.ToInt64()
		case "ToUint64":
			actual, err = subject.ToUint64()
		case "ToFloat32":
			actual, err = subject.ToFloat32()
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
