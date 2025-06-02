package schema_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/13k/valve.go/steamweb/schema"
)

//// Methods {{{

func TestMethods(t *testing.T) {
	t.Parallel()

	suite.Run(t, new(MethodsTestSuite))
}

type MethodsTestSuite struct {
	Suite
}

func (s *MethodsTestSuite) TestNewMethods() {
	require := s.Require()

	testCases := []struct {
		testName string
		subject  []*schema.Method
		expected schema.Methods
		err      string
	}{
		{
			testName: "nil",
			subject:  nil,
			expected: schema.Methods{},
		},
		{
			testName: "empty",
			subject:  []*schema.Method{},
			expected: schema.Methods{},
		},
		{
			testName: "valid",
			subject:  []*schema.Method{s.createMethodValid()},
			expected: schema.Methods{s.createMethodValid()},
		},
		{
			testName: "invalid/name",
			subject:  []*schema.Method{s.createMethodInvalidName()},
			err:      `invalid name "9Method"`,
		},
		{
			testName: "invalid/version",
			subject:  []*schema.Method{s.createMethodInvalidVersion()},
			err:      `invalid version 0`,
		},
		{
			testName: "invalid/http_method",
			subject:  []*schema.Method{s.createMethodInvalidHTTPMethod()},
			err:      `invalid HTTP method "xyz"`,
		},
	}

	for _, testCase := range testCases {
		s.Run(testCase.testName, func() {
			actual, err := schema.NewMethods(testCase.subject...)

			if testCase.err != "" {
				require.ErrorContains(err, testCase.err)
			} else {
				require.NoError(err)
				s.AssertMethodsEqual(testCase.expected, actual)
			}
		})
	}
}

func (s *MethodsTestSuite) TestGet() {
	assert := s.Assert()
	require := s.Require()

	key := &schema.MethodKey{
		Name:    "Method",
		Version: 1,
	}

	testCases := []struct {
		testName  string
		subject   schema.Methods
		key       *schema.MethodKey
		index     int
		err       string
		errTarget error
	}{
		{
			testName:  "nil",
			subject:   schema.Methods(nil),
			key:       key,
			err:       `method not found ("Method", 1)`,
			errTarget: &schema.MethodNotFoundError{},
		},
		{
			testName:  "empty",
			subject:   schema.Methods{},
			key:       key,
			err:       `method not found ("Method", 1)`,
			errTarget: &schema.MethodNotFoundError{},
		},
		{
			testName:  "not_found/name",
			subject:   s.createMethods(),
			key:       &schema.MethodKey{Name: "Bar", Version: 1},
			err:       `method not found ("Bar", 1)`,
			errTarget: &schema.MethodNotFoundError{},
		},
		{
			testName:  "not_found/version",
			subject:   s.createMethods(),
			key:       &schema.MethodKey{Name: "Method", Version: 2},
			err:       `method not found ("Method", 2)`,
			errTarget: &schema.MethodNotFoundError{},
		},
		{
			testName:  "not_found/version_zero",
			subject:   s.createMethods(),
			key:       &schema.MethodKey{Name: "Method", Version: 0},
			err:       `method not found ("Method", 0)`,
			errTarget: &schema.MethodNotFoundError{},
		},
		{
			testName: "found",
			subject:  s.createMethods(),
			key:      key,
			index:    0,
		},
	}

	for _, testCase := range testCases {
		s.Run(testCase.testName, func() {
			actual, err := testCase.subject.Get(testCase.key)

			if testCase.err != "" {
				require.ErrorContains(err, testCase.err)
				require.ErrorAs(err, &testCase.errTarget)
			} else {
				require.NoError(err)
				assert.Equal(testCase.key, actual.Key())
				assert.Same(testCase.subject[testCase.index], actual)
			}
		})
	}
}

//// Methods }}}
//// MethodsIndex {{{

func TestMethodsIndex(t *testing.T) {
	t.Parallel()

	suite.Run(t, new(MethodsIndexTestSuite))
}

type MethodsIndexTestSuite struct {
	Suite
}

func (s *MethodsIndexTestSuite) TestNewMethodsIndex() {
	assert := s.Assert()

	methods := schema.MustNewMethods(
		s.createMethod("Foo", 1),
		s.createMethod("Foo", 16),
		s.createMethod("Foo", 32),
		s.createMethod("Bar", 32),
		s.createMethod("BarA", 32),
		s.createMethod("BarB", 32),
		s.createMethod("Baz", 48),
	)

	testCases := []struct {
		testName string
		subject  schema.Methods
		expected schema.MethodsIndex
	}{
		{
			testName: "nil",
			subject:  nil,
			expected: nil,
		},
		{
			testName: "empty",
			subject:  schema.Methods{},
			expected: schema.MethodsIndex{},
		},
		{
			testName: "filled",
			subject:  methods,
			expected: schema.MethodsIndex{
				schema.MethodKey{Name: "Foo", Version: 1}:   methods[0],
				schema.MethodKey{Name: "Foo", Version: 16}:  methods[1],
				schema.MethodKey{Name: "Foo", Version: 32}:  methods[2],
				schema.MethodKey{Name: "Bar", Version: 32}:  methods[3],
				schema.MethodKey{Name: "BarA", Version: 32}: methods[4],
				schema.MethodKey{Name: "BarB", Version: 32}: methods[5],
				schema.MethodKey{Name: "Baz", Version: 48}:  methods[6],
			},
		},
	}

	for _, testCase := range testCases {
		s.Run(testCase.testName, func() {
			actual := schema.NewMethodsIndex(testCase.subject)

			assert.Equal(testCase.expected, actual)
		})
	}
}

func (s *MethodsIndexTestSuite) TestVersions() {
	assert := s.Assert()

	methods := schema.MustNewMethods(
		s.createMethod("Foo", 1),
		s.createMethod("Foo", 16),
		s.createMethod("Foo", 32),
		s.createMethod("Baz", 48),
	)

	testCases := []struct {
		testName string
		subject  schema.MethodsIndex
		expected []int
	}{
		{
			testName: "nil",
			subject:  nil,
			expected: nil,
		},
		{
			testName: "empty",
			subject:  schema.MethodsIndex{},
			expected: []int{},
		},
		{
			testName: "filled",
			subject:  schema.NewMethodsIndex(methods),
			expected: []int{1, 16, 32, 48},
		},
	}

	for _, testCase := range testCases {
		s.Run(testCase.testName, func() {
			actual := testCase.subject.Versions()

			assert.ElementsMatch(testCase.expected, actual)
		})
	}
}

//// MethodsIndex }}}
//// MethodGroupsByName {{{

func TestMethodGroupsByName(t *testing.T) {
	t.Parallel()

	suite.Run(t, new(MethodGroupsByNameTestSuite))
}

type MethodGroupsByNameTestSuite struct {
	Suite
}

func (s *MethodGroupsByNameTestSuite) TestNewMethodGroupsByName() {
	assert := s.Assert()

	methods := schema.MustNewMethods(
		s.createMethod("Foo", 1),
		s.createMethod("Foo", 16),
		s.createMethod("Foo", 32),
		s.createMethod("Bar", 32),
		s.createMethod("BarA", 32),
		s.createMethod("BarB", 32),
		s.createMethod("Baz", 48),
	)

	testCases := []struct {
		testName string
		subject  schema.Methods
		expected schema.MethodGroupsByName
	}{
		{
			testName: "nil",
			subject:  nil,
			expected: nil,
		},
		{
			testName: "empty",
			subject:  schema.Methods{},
			expected: schema.MethodGroupsByName{},
		},
		{
			testName: "filled",
			subject:  methods,
			expected: schema.MethodGroupsByName{
				"Foo": {
					schema.MethodKey{Name: "Foo", Version: 1}:  methods[0],
					schema.MethodKey{Name: "Foo", Version: 16}: methods[1],
					schema.MethodKey{Name: "Foo", Version: 32}: methods[2],
				},
				"Bar": {
					schema.MethodKey{Name: "Bar", Version: 32}: methods[3],
				},
				"BarA": {
					schema.MethodKey{Name: "BarA", Version: 32}: methods[4],
				},
				"BarB": {
					schema.MethodKey{Name: "BarB", Version: 32}: methods[5],
				},
				"Baz": {
					schema.MethodKey{Name: "Baz", Version: 48}: methods[6],
				},
			},
		},
	}

	for _, testCase := range testCases {
		s.Run(testCase.testName, func() {
			actual := schema.NewMethodGroupsByName(testCase.subject)

			assert.Equal(testCase.expected, actual)
		})
	}
}

//// MethodGroupsByName }}}
