package schema_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/13k/valve.go/steamlib"
	"github.com/13k/valve.go/steamweb/schema"
)

func TestInterface(t *testing.T) {
	t.Parallel()

	suite.Run(t, new(InterfaceTestSuite))
}

type InterfaceTestSuite struct {
	Suite
}

func (s *InterfaceTestSuite) TestNewInterface() {
	assert := s.Assert()
	require := s.Require()

	testCases := []struct {
		testName  string
		name      string
		methods   schema.Methods
		key       *schema.InterfaceKey
		err       string
		errTarget error
	}{
		{
			testName:  "invalid/name",
			name:      "iface",
			methods:   nil,
			err:       `invalid name "iface"`,
			errTarget: &schema.InvalidInterfaceError{},
		},
		{
			testName: "valid/no_appid",
			name:     "IFace",
			methods:  nil,
			key:      &schema.InterfaceKey{Name: "IFace", AppID: steamlib.AppID(0)},
		},
		{
			testName: "valid/appid",
			name:     "IFace_42",
			methods:  nil,
			key:      &schema.InterfaceKey{Name: "IFace", AppID: steamlib.AppID(42)},
		},
		{
			testName:  "invalid/methods/name",
			name:      "IFace_42",
			methods:   s.createMethodsInvalidName(),
			err:       `invalid name "9Method"`,
			errTarget: &schema.InvalidInterfaceError{},
		},
		{
			testName:  "invalid/methods/version",
			name:      "IFace_42",
			methods:   s.createMethodsInvalidVersion(),
			err:       `invalid version 0`,
			errTarget: &schema.InvalidInterfaceError{},
		},
		{
			testName:  "invalid/methods/http_method",
			name:      "IFace_42",
			methods:   s.createMethodsInvalidHTTPMethod(),
			err:       `invalid HTTP method "xyz"`,
			errTarget: &schema.InvalidInterfaceError{},
		},
		{
			testName: "valid/methods",
			name:     "IFace_42",
			methods:  s.createMethods(),
			key:      &schema.InterfaceKey{Name: "IFace", AppID: steamlib.AppID(42)},
		},
	}

	for _, testCase := range testCases {
		s.Run(testCase.testName, func() {
			actual, err := schema.NewInterface(testCase.name, testCase.methods)

			if testCase.err != "" {
				require.ErrorContains(err, testCase.err)
				require.ErrorAs(err, &testCase.errTarget)
			} else {
				require.NoError(err)
				assert.Equal(testCase.key, actual.Key())
				assert.Equal(testCase.name, actual.Name)
				s.AssertMethodsEqual(testCase.methods, actual.Methods)
			}
		})
	}
}
