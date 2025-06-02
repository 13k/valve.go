package schema_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/13k/valve.go/steamlib"
	"github.com/13k/valve.go/steamweb/schema"
)

//// Interfaces {{{

func TestInterfaces(t *testing.T) {
	t.Parallel()

	suite.Run(t, new(InterfacesTestSuite))
}

type InterfacesTestSuite struct {
	Suite
}

func (s *InterfacesTestSuite) TestNewInterfaces() {
	require := s.Require()

	testCases := []struct {
		testName string
		subject  []*schema.Interface
		expected schema.Interfaces
		err      string
	}{
		{
			testName: "nil",
			subject:  nil,
			expected: schema.Interfaces{},
		},
		{
			testName: "empty",
			subject:  []*schema.Interface{},
			expected: schema.Interfaces{},
		},
		{
			testName: "valid",
			subject:  []*schema.Interface{s.createInterfaceNoAppID(), s.createInterfaceAppID()},
			expected: schema.Interfaces{s.createInterfaceNoAppID(), s.createInterfaceAppID()},
		},
		{
			testName: "invalid",
			subject:  []*schema.Interface{s.createInterfaceInvalidName()},
			err:      `invalid name "invalid"`,
		},
	}

	for _, testCase := range testCases {
		s.Run(testCase.testName, func() {
			actual, err := schema.NewInterfaces(testCase.subject...)

			if testCase.err != "" {
				require.ErrorContains(err, testCase.err)
			} else {
				require.NoError(err)
				s.AssertInterfacesEqual(testCase.expected, actual)
			}
		})
	}
}

func (s *InterfacesTestSuite) TestGet() {
	assert := s.Assert()
	require := s.Require()

	key := &schema.InterfaceKey{
		Name:  "IFace",
		AppID: 123,
	}

	testCases := []struct {
		testName  string
		subject   schema.Interfaces
		key       *schema.InterfaceKey
		index     int
		err       string
		errTarget error
	}{
		{
			testName:  "nil",
			subject:   schema.Interfaces(nil),
			key:       key,
			err:       `interface not found ("IFace", 123)`,
			errTarget: &schema.InterfaceNotFoundError{},
		},
		{
			testName:  "empty",
			subject:   schema.Interfaces{},
			key:       key,
			err:       `interface not found ("IFace", 123)`,
			errTarget: &schema.InterfaceNotFoundError{},
		},
		{
			testName:  "not_found/name",
			subject:   s.createInterfaces(),
			key:       &schema.InterfaceKey{Name: "IFoo", AppID: 123},
			err:       `interface not found ("IFoo", 123)`,
			errTarget: &schema.InterfaceNotFoundError{},
		},
		{
			testName:  "not_found/appid",
			subject:   s.createInterfaces(),
			key:       &schema.InterfaceKey{Name: "IFace", AppID: 456},
			err:       `interface not found ("IFace", 456)`,
			errTarget: &schema.InterfaceNotFoundError{},
		},
		{
			testName:  "not_found/appid_zero",
			subject:   s.createInterfaces(),
			key:       &schema.InterfaceKey{Name: "IFace", AppID: 0},
			err:       `interface not found ("IFace", 0)`,
			errTarget: &schema.InterfaceNotFoundError{},
		},
		{
			testName: "found",
			subject:  s.createInterfaces(),
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

//// Interfaces }}}
//// InterfacesIndex {{{

func TestInterfacesIndex(t *testing.T) {
	t.Parallel()

	suite.Run(t, new(InterfacesIndexTestSuite))
}

type InterfacesIndexTestSuite struct {
	Suite
}

func (s *InterfacesIndexTestSuite) TestNewInterfacesIndex() {
	assert := s.Assert()

	interfaces := schema.MustNewInterfaces(
		s.createInterface("IFoo", 0, nil),
		s.createInterface("IFoo", 16, nil),
		s.createInterface("IFoo", 32, nil),
		s.createInterface("IBar", 32, nil),
		s.createInterface("IBar_A", 32, nil),
		s.createInterface("IBar_B", 32, nil),
		s.createInterface("IBaz", 48, nil),
	)

	testCases := []struct {
		testName string
		subject  schema.Interfaces
		expected schema.InterfacesIndex
	}{
		{
			testName: "nil",
			subject:  nil,
			expected: nil,
		},
		{
			testName: "empty",
			subject:  schema.Interfaces{},
			expected: schema.InterfacesIndex{},
		},
		{
			testName: "filled",
			subject:  interfaces,
			expected: schema.InterfacesIndex{
				schema.InterfaceKey{Name: "IFoo", AppID: 0}:    interfaces[0],
				schema.InterfaceKey{Name: "IFoo", AppID: 16}:   interfaces[1],
				schema.InterfaceKey{Name: "IFoo", AppID: 32}:   interfaces[2],
				schema.InterfaceKey{Name: "IBar", AppID: 32}:   interfaces[3],
				schema.InterfaceKey{Name: "IBar_A", AppID: 32}: interfaces[4],
				schema.InterfaceKey{Name: "IBar_B", AppID: 32}: interfaces[5],
				schema.InterfaceKey{Name: "IBaz", AppID: 48}:   interfaces[6],
			},
		},
	}

	for _, testCase := range testCases {
		s.Run(testCase.testName, func() {
			actual := schema.NewInterfacesIndex(testCase.subject)

			assert.Equal(testCase.expected, actual)
		})
	}
}

func (s *InterfacesIndexTestSuite) TestName() {
	assert := s.Assert()

	interfaces := schema.MustNewInterfaces(
		s.createInterface("IFoo", 0, nil),
		s.createInterface("IFoo", 16, nil),
		s.createInterface("IFoo", 32, nil),
		s.createInterface("IBaz", 48, nil),
	)

	testCases := []struct {
		testName string
		subject  schema.InterfacesIndex
		expected string
	}{
		{
			testName: "nil",
			subject:  nil,
			expected: "",
		},
		{
			testName: "empty",
			subject:  schema.InterfacesIndex{},
			expected: "",
		},
		{
			testName: "filled",
			subject:  schema.NewInterfacesIndex(interfaces),
			expected: "IFoo",
		},
	}

	for _, testCase := range testCases {
		s.Run(testCase.testName, func() {
			actual := testCase.subject.Name()

			assert.Equal(testCase.expected, actual)
		})
	}
}

func (s *InterfacesIndexTestSuite) TestAppIDs() {
	assert := s.Assert()

	interfaces := schema.MustNewInterfaces(
		s.createInterface("IFoo", 0, nil),
		s.createInterface("IFoo", 16, nil),
		s.createInterface("IFoo", 32, nil),
		s.createInterface("IBaz", 48, nil),
	)

	testCases := []struct {
		testName string
		subject  schema.InterfacesIndex
		expected []steamlib.AppID
	}{
		{
			testName: "nil",
			subject:  nil,
			expected: nil,
		},
		{
			testName: "empty",
			subject:  schema.InterfacesIndex{},
			expected: []steamlib.AppID{},
		},
		{
			testName: "filled",
			subject:  schema.NewInterfacesIndex(interfaces),
			expected: []steamlib.AppID{
				steamlib.AppID(16),
				steamlib.AppID(32),
				steamlib.AppID(48),
			},
		},
	}

	for _, testCase := range testCases {
		s.Run(testCase.testName, func() {
			actual := testCase.subject.AppIDs()

			assert.ElementsMatch(testCase.expected, actual)
		})
	}
}

func (s *InterfacesIndexTestSuite) TestGroupMethodsByName() {
	assert := s.Assert()

	interfaces := schema.MustNewInterfaces(
		schema.MustNewInterface(
			"Interface_1",
			schema.MustNewMethods(
				s.createMethod("Method1", 1),
				s.createMethod("Method2", 1),
				s.createMethod("Method3", 1),
			),
		),
		schema.MustNewInterface(
			"Interface_2",
			schema.MustNewMethods(
				s.createMethod("Method1", 1),
				s.createMethod("Method1", 2),
				s.createMethod("Method1", 3),
			),
		),
		schema.MustNewInterface(
			"Interface_3",
			schema.MustNewMethods(
				s.createMethod("Method1", 1),
				s.createMethod("Method2", 2),
				s.createMethod("Method3", 3),
			),
		),
	)

	testCases := []struct {
		testName string
		subject  schema.InterfacesIndex
		expected schema.MethodGroupsByName
	}{
		{
			testName: "nil",
			subject:  nil,
			expected: nil,
		},
		{
			testName: "empty",
			subject:  schema.InterfacesIndex{},
			expected: nil,
		},
		{
			testName: "filled",
			subject:  schema.NewInterfacesIndex(interfaces),
			expected: schema.MethodGroupsByName{
				"Method1": {
					schema.MethodKey{Name: "Method1", Version: 1}: interfaces[2].Methods[0],
					schema.MethodKey{Name: "Method1", Version: 2}: interfaces[1].Methods[1],
					schema.MethodKey{Name: "Method1", Version: 3}: interfaces[1].Methods[2],
				},
				"Method2": {
					schema.MethodKey{Name: "Method2", Version: 1}: interfaces[0].Methods[1],
					schema.MethodKey{Name: "Method2", Version: 2}: interfaces[2].Methods[1],
				},
				"Method3": {
					schema.MethodKey{Name: "Method3", Version: 1}: interfaces[0].Methods[2],
					schema.MethodKey{Name: "Method3", Version: 3}: interfaces[2].Methods[2],
				},
			},
		},
	}

	for _, testCase := range testCases {
		s.Run(testCase.testName, func() {
			actual := testCase.subject.GroupMethodsByName()

			assert.Equal(testCase.expected, actual)
		})
	}
}

//// InterfacesIndex }}}
//// InterfaceGroupsByName {{{

func TestInterfaceGroupsByName(t *testing.T) {
	t.Parallel()

	suite.Run(t, new(InterfaceGroupsByNameTestSuite))
}

type InterfaceGroupsByNameTestSuite struct {
	Suite
}

func (s *InterfaceGroupsByNameTestSuite) TestNewInterfaceGroupsByName() {
	assert := s.Assert()

	interfaces := schema.MustNewInterfaces(
		s.createInterface("IFoo", 0, nil),
		s.createInterface("IFoo", 16, nil),
		s.createInterface("IFoo", 32, nil),
		s.createInterface("IBar", 32, nil),
		s.createInterface("IBar_A", 32, nil),
		s.createInterface("IBar_B", 32, nil),
		s.createInterface("IBaz", 48, nil),
	)

	testCases := []struct {
		testName string
		subject  schema.Interfaces
		expected schema.InterfaceGroupsByName
	}{
		{
			testName: "nil",
			subject:  nil,
			expected: nil,
		},
		{
			testName: "empty",
			subject:  schema.Interfaces{},
			expected: schema.InterfaceGroupsByName{},
		},
		{
			testName: "filled",
			subject:  interfaces,
			expected: schema.InterfaceGroupsByName{
				"IFoo": {
					schema.InterfaceKey{Name: "IFoo", AppID: 0}:  interfaces[0],
					schema.InterfaceKey{Name: "IFoo", AppID: 16}: interfaces[1],
					schema.InterfaceKey{Name: "IFoo", AppID: 32}: interfaces[2],
				},
				"IBar": {
					schema.InterfaceKey{Name: "IBar", AppID: 32}: interfaces[3],
				},
				"IBar_A": {
					schema.InterfaceKey{Name: "IBar_A", AppID: 32}: interfaces[4],
				},
				"IBar_B": {
					schema.InterfaceKey{Name: "IBar_B", AppID: 32}: interfaces[5],
				},
				"IBaz": {
					schema.InterfaceKey{Name: "IBaz", AppID: 48}: interfaces[6],
				},
			},
		},
	}

	for _, testCase := range testCases {
		s.Run(testCase.testName, func() {
			actual := schema.NewInterfaceGroupsByName(testCase.subject)

			assert.Equal(testCase.expected, actual)
		})
	}
}

//// InterfaceGroupsByName }}}
