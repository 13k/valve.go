package schema_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/13k/valve.go/steamweb/schema"
)

func TestSchema(t *testing.T) {
	t.Parallel()

	suite.Run(t, new(SchemaTestSuite))
}

type SchemaTestSuite struct {
	Suite
}

func (s *SchemaTestSuite) TestNewSchema() {
	require := s.Require()

	testCases := []struct {
		testName  string
		subject   schema.Interfaces
		err       string
		errTarget error
	}{
		{
			testName: "valid/nil",
			subject:  nil,
		},
		{
			testName: "valid/empty",
			subject:  schema.Interfaces{},
		},
		{
			testName: "valid/interfaces",
			subject:  s.createInterfaces(),
		},
		{
			testName:  "invalid/interfaces/name",
			subject:   s.createInterfacesInvalidName(),
			err:       `invalid name "invalid"`,
			errTarget: &schema.InvalidSchemaError{},
		},
		{
			testName: "valid/interfaces/methods",
			subject:  s.createInterfacesMethods(),
		},
		{
			testName:  "invalid/interfaces/methods/name",
			subject:   s.createInterfacesMethodsInvalidName(),
			err:       `invalid name "9Method"`,
			errTarget: &schema.InvalidSchemaError{},
		},
		{
			testName:  "invalid/interfaces/methods/version",
			subject:   s.createInterfacesMethodsInvalidVersion(),
			err:       `invalid version 0`,
			errTarget: &schema.InvalidSchemaError{},
		},
		{
			testName:  "invalid/interfaces/methods/http_method",
			subject:   s.createInterfacesMethodsInvalidHTTPMethod(),
			err:       `invalid HTTP method "xyz"`,
			errTarget: &schema.InvalidSchemaError{},
		},
	}

	for _, testCase := range testCases {
		s.Run(testCase.testName, func() {
			actual, err := schema.NewSchema(testCase.subject)

			if testCase.err != "" {
				require.ErrorContains(err, testCase.err)
				require.ErrorAs(err, &testCase.errTarget)
			} else {
				require.NoError(err)
				s.AssertInterfacesEqual(testCase.subject, actual.Interfaces)
			}
		})
	}
}
