package steamlib_test

import (
	"math"
	"strconv"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/13k/valve.go/steamlib"
)

func TestAppID(t *testing.T) {
	t.Parallel()

	suite.Run(t, new(AppIDTestSuite))
}

type AppIDTestSuite struct {
	suite.Suite
}

func (s *AppIDTestSuite) TestParse() {
	assert := s.Assert()
	require := s.Require()

	testCases := []struct {
		name     string
		subject  string
		expected steamlib.AppID
		err      string
	}{
		{
			name:    "empty",
			subject: "",
			err:     `AppID from string ""`,
		},
		{
			name:    "non_number",
			subject: "foo",
			err:     `AppID from string "foo"`,
		},
		{
			name:    "negative_number",
			subject: "-1",
			err:     `AppID from string "-1"`,
		},
		{
			name:    "overflow",
			subject: strconv.FormatUint(math.MaxUint32+1, 10),
			err:     `AppID from string "4294967296"`,
		},
		{
			name:     "valid",
			subject:  "42",
			expected: steamlib.AppID(42),
		},
	}

	for _, testCase := range testCases {
		s.Run(testCase.name, func() {
			actual, err := steamlib.ParseAppID(testCase.subject)

			if testCase.err != "" {
				require.ErrorContains(err, testCase.err)
			} else {
				require.NoError(err)
				assert.Equal(testCase.expected, actual)
			}
		})
	}
}

func (s *AppIDTestSuite) TestMarshalBinary() {
	assert := s.Assert()
	require := s.Require()

	testCases := []struct {
		name     string
		subject  steamlib.AppID
		expected []byte
		err      string
	}{
		{
			name:     "min",
			subject:  steamlib.AppID(0),
			expected: []byte{0x00, 0x00, 0x00, 0x00},
		},
		{
			name:     "max",
			subject:  steamlib.AppID(math.MaxUint32),
			expected: []byte{0xff, 0xff, 0xff, 0xff},
		},
		{
			name:     "mid",
			subject:  steamlib.AppID(1337),
			expected: []byte{0x39, 0x05, 0x00, 0x00},
		},
	}

	for _, testCase := range testCases {
		s.Run(testCase.name, func() {
			actual, err := testCase.subject.MarshalBinary()

			if testCase.err != "" {
				require.ErrorContains(err, testCase.err)
			} else {
				require.NoError(err)
				assert.Equal(testCase.expected, actual)
			}
		})
	}
}

func (s *AppIDTestSuite) TestUnmarshalBinary() {
	assert := s.Assert()
	require := s.Require()

	testCases := []struct {
		name     string
		subject  []byte
		expected steamlib.AppID
		err      string
	}{
		{
			name:    "nil",
			subject: nil,
			err:     `AppID from binary data ""`,
		},
		{
			name:    "empty",
			subject: []byte{},
			err:     `AppID from binary data ""`,
		},
		{
			name:    "incomplete",
			subject: []byte{0x00},
			err:     `AppID from binary data "\x00"`,
		},
		{
			name:     "min",
			subject:  []byte{0x00, 0x00, 0x00, 0x00},
			expected: steamlib.AppID(0),
		},
		{
			name:     "max",
			subject:  []byte{0xff, 0xff, 0xff, 0xff},
			expected: steamlib.AppID(math.MaxUint32),
		},
		{
			name:     "mid",
			subject:  []byte{0x39, 0x05, 0x00, 0x00},
			expected: steamlib.AppID(1337),
		},
		{
			name:     "excess",
			subject:  []byte{0x39, 0x05, 0x00, 0x00, 0x01, 0x01},
			expected: steamlib.AppID(1337),
		},
	}

	for _, testCase := range testCases {
		s.Run(testCase.name, func() {
			actual := steamlib.AppID(0)
			err := actual.UnmarshalBinary(testCase.subject)

			if testCase.err != "" {
				require.ErrorContains(err, testCase.err)
			} else {
				require.NoError(err)
				assert.Equal(testCase.expected, actual)
			}
		})
	}
}

func (s *AppIDTestSuite) TestMarshalText() {
	assert := s.Assert()
	require := s.Require()

	testCases := []struct {
		name     string
		subject  steamlib.AppID
		expected []byte
		err      string
	}{
		{
			name:     "min",
			subject:  steamlib.AppID(0),
			expected: []byte("0"),
		},
		{
			name:     "max",
			subject:  steamlib.AppID(math.MaxUint32),
			expected: []byte("4294967295"),
		},
		{
			name:     "mid",
			subject:  steamlib.AppID(1337),
			expected: []byte("1337"),
		},
	}

	for _, testCase := range testCases {
		s.Run(testCase.name, func() {
			actual, err := testCase.subject.MarshalText()

			if testCase.err != "" {
				require.ErrorContains(err, testCase.err)
			} else {
				require.NoError(err)
				assert.Equal(testCase.expected, actual)
			}
		})
	}
}

func (s *AppIDTestSuite) TestUnmarshalText() {
	assert := s.Assert()
	require := s.Require()

	testCases := []struct {
		name     string
		subject  []byte
		expected steamlib.AppID
		err      string
	}{
		{
			name:    "nil",
			subject: nil,
			err:     `AppID from text data ""`,
		},
		{
			name:    "empty",
			subject: []byte{},
			err:     `AppID from text data ""`,
		},
		{
			name:    "incomplete",
			subject: []byte(""),
			err:     `AppID from text data ""`,
		},
		{
			name:    "non_number",
			subject: []byte("foo"),
			err:     `AppID from text data "foo"`,
		},
		{
			name:    "negative_number",
			subject: []byte("-1"),
			err:     `AppID from text data "-1"`,
		},
		{
			name:    "overflow",
			subject: []byte("4294967296"),
			err:     `AppID from text data "4294967296"`,
		},
		{
			name:     "min",
			subject:  []byte("0"),
			expected: steamlib.AppID(0),
		},
		{
			name:     "max",
			subject:  []byte("4294967295"),
			expected: steamlib.AppID(math.MaxUint32),
		},
		{
			name:     "mid",
			subject:  []byte("1337"),
			expected: steamlib.AppID(1337),
		},
	}

	for _, testCase := range testCases {
		s.Run(testCase.name, func() {
			actual := steamlib.AppID(0)
			err := actual.UnmarshalText(testCase.subject)

			if testCase.err != "" {
				require.ErrorContains(err, testCase.err)
			} else {
				require.NoError(err)
				assert.Equal(testCase.expected, actual)
			}
		})
	}
}

func (s *AppIDTestSuite) TestMarshalJSON() {
	assert := s.Assert()
	require := s.Require()

	testCases := []struct {
		name     string
		subject  steamlib.AppID
		expected []byte
		err      string
	}{
		{
			name:     "min",
			subject:  steamlib.AppID(0),
			expected: []byte("0"),
		},
		{
			name:     "max",
			subject:  steamlib.AppID(math.MaxUint32),
			expected: []byte("4294967295"),
		},
		{
			name:     "mid",
			subject:  steamlib.AppID(1337),
			expected: []byte("1337"),
		},
	}

	for _, testCase := range testCases {
		s.Run(testCase.name, func() {
			actual, err := testCase.subject.MarshalJSON()

			if testCase.err != "" {
				require.ErrorContains(err, testCase.err)
			} else {
				require.NoError(err)
				assert.Equal(testCase.expected, actual)
			}
		})
	}
}

func (s *AppIDTestSuite) TestUnmarshalJSON() {
	assert := s.Assert()
	require := s.Require()

	testCases := []struct {
		name     string
		subject  []byte
		expected steamlib.AppID
		err      string
	}{
		{
			name:    "nil",
			subject: nil,
			err:     `AppID from JSON data ""`,
		},
		{
			name:    "empty",
			subject: []byte{},
			err:     `AppID from JSON data ""`,
		},
		{
			name:    "incomplete",
			subject: []byte(""),
			err:     `AppID from JSON data ""`,
		},
		{
			name:    "non_number",
			subject: []byte("foo"),
			err:     `AppID from JSON data "foo"`,
		},
		{
			name:    "negative_number",
			subject: []byte("-1"),
			err:     `AppID from JSON data "-1"`,
		},
		{
			name:    "overflow",
			subject: []byte("4294967296"),
			err:     `AppID from JSON data "4294967296"`,
		},
		{
			name:     "min",
			subject:  []byte("0"),
			expected: steamlib.AppID(0),
		},
		{
			name:     "max",
			subject:  []byte("4294967295"),
			expected: steamlib.AppID(math.MaxUint32),
		},
		{
			name:     "mid",
			subject:  []byte("1337"),
			expected: steamlib.AppID(1337),
		},
	}

	for _, testCase := range testCases {
		s.Run(testCase.name, func() {
			actual := steamlib.AppID(0)
			err := actual.UnmarshalJSON(testCase.subject)

			if testCase.err != "" {
				require.ErrorContains(err, testCase.err)
			} else {
				require.NoError(err)
				assert.Equal(testCase.expected, actual)
			}
		})
	}
}
