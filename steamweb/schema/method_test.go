package schema_test

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/13k/valve.go/steamweb/schema"
)

func TestMethod(t *testing.T) {
	t.Parallel()

	suite.Run(t, new(MethodTestSuite))
}

type MethodTestSuite struct {
	Suite
}

func (s *MethodTestSuite) TestNewMethod() {
	assert := s.Assert()
	require := s.Require()

	testCases := []struct {
		testName     string
		name         string
		version      int
		httpMethod   string
		params       schema.MethodParams
		undocumented bool
		key          *schema.MethodKey
		err          string
		errTarget    error
	}{
		{
			testName:     "valid",
			name:         "Method",
			version:      1,
			httpMethod:   http.MethodGet,
			params:       s.createMethodParamsRequired("name", "type"),
			undocumented: true,
			key:          &schema.MethodKey{Name: "Method", Version: 1},
		},
		{
			testName:   "invalid/name",
			name:       "9Method",
			version:    1,
			httpMethod: http.MethodGet,
			err:        `invalid name "9Method"`,
			errTarget:  &schema.InvalidMethodError{},
		},
		{
			testName:   "invalid/version",
			name:       "Method",
			version:    0,
			httpMethod: http.MethodGet,
			err:        `invalid version 0`,
			errTarget:  &schema.InvalidMethodError{},
		},
		{
			testName:   "invalid/http_method",
			name:       "Method",
			version:    1,
			httpMethod: "foo",
			err:        `invalid HTTP method "foo"`,
			errTarget:  &schema.InvalidMethodError{},
		},
	}

	for _, testCase := range testCases {
		s.Run(testCase.testName, func() {
			actual, err := schema.NewMethod(
				testCase.name,
				testCase.version,
				testCase.httpMethod,
				testCase.params,
				testCase.undocumented,
			)

			if testCase.err != "" {
				require.ErrorContains(err, testCase.err)
				require.ErrorAs(err, &testCase.errTarget)
			} else {
				require.NoError(err)
				assert.Equal(testCase.key, actual.Key())
				assert.Equal(testCase.name, actual.Name)
				assert.Equal(testCase.version, actual.Version)
				assert.Equal(testCase.httpMethod, actual.HTTPMethod)
				assert.Equal(testCase.params, actual.Params)
				assert.Equal(testCase.undocumented, actual.Undocumented)
			}
		})
	}
}
