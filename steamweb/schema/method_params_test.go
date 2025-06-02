package schema_test

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/13k/valve.go/steamweb/schema"
)

func TestMethodParams(t *testing.T) {
	t.Parallel()

	suite.Run(t, new(MethodParamsTestSuite))
}

type MethodParamsTestSuite struct {
	Suite
}

func (s *MethodParamsTestSuite) TestValidateParams() {
	require := s.Require()

	testCases := []struct {
		testName  string
		subject   schema.MethodParams
		params    url.Values
		err       string
		errTarget error
	}{
		{
			testName: "optional/valid",
			subject:  s.createMethodParamsOptional("param", "type"),
			params:   url.Values{"param": {"foo"}},
		},
		// no type validation
		{
			testName: "optional/invalid",
			subject:  s.createMethodParamsOptional("param", "type"),
			params:   url.Values{"param": {"<invalid>"}},
		},
		{
			testName: "optional/missing",
			subject:  s.createMethodParamsOptional("param", "type"),
			params:   nil,
		},
		{
			testName: "required/valid",
			subject:  s.createMethodParamsRequired("param", "type"),
			params:   url.Values{"param": {"foo"}},
		},
		// no type validation
		{
			testName: "required/invalid",
			subject:  s.createMethodParamsRequired("param", "type"),
			params:   url.Values{"param": {"<invalid>"}},
		},
		{
			testName:  "required/missing",
			subject:   s.createMethodParamsRequired("param", "type"),
			err:       `missing required parameter "param"`,
			errTarget: &schema.RequiredParameterError{},
		},
	}

	for _, testCase := range testCases {
		s.Run(testCase.testName, func() {
			err := testCase.subject.ValidateParams(testCase.params)

			if testCase.err != "" {
				require.ErrorContains(err, testCase.err)
				require.ErrorAs(err, &testCase.errTarget)
			} else {
				require.NoError(err)
			}
		})
	}
}
