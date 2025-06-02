package schema_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/13k/valve.go/steamweb/schema"
)

func TestMethodParam(t *testing.T) {
	t.Parallel()

	suite.Run(t, new(MethodParamTestSuite))
}

type MethodParamTestSuite struct {
	Suite
}

func (s *MethodParamTestSuite) TestValidateValue() {
	require := s.Require()

	testCases := []struct {
		testName  string
		subject   *schema.MethodParam
		value     string
		err       string
		errTarget error
	}{
		{
			testName: "optional/valid",
			subject:  s.createMethodParamOptional("param", "type"),
			value:    "foo",
		},
		// no type validation
		{
			testName: "optional/invalid",
			subject:  s.createMethodParamOptional("param", "type"),
			value:    "<invalid>",
		},
		{
			testName: "optional/missing",
			subject:  s.createMethodParamOptional("param", "type"),
			value:    "",
		},
		{
			testName: "required/valid",
			subject:  s.createMethodParamRequired("param", "type"),
			value:    "foo",
		},
		// no type validation
		{
			testName: "required/invalid",
			subject:  s.createMethodParamRequired("param", "type"),
			value:    "<invalid>",
		},
		{
			testName:  "required/missing",
			subject:   s.createMethodParamRequired("param", "type"),
			err:       `missing required parameter "param"`,
			errTarget: &schema.RequiredParameterError{},
		},
	}

	for _, testCase := range testCases {
		s.Run(testCase.testName, func() {
			err := testCase.subject.ValidateValue(testCase.value)

			if testCase.err != "" {
				require.ErrorContains(err, testCase.err)
				require.ErrorAs(err, &testCase.errTarget)
			} else {
				require.NoError(err)
			}
		})
	}
}
