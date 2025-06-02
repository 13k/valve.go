package schema_test

import (
	"fmt"
	"net/http"

	"github.com/stretchr/testify/suite"

	"github.com/13k/valve.go/steamlib"
	"github.com/13k/valve.go/steamweb/schema"
)

type Suite struct {
	suite.Suite
}

func (s *Suite) AssertInterfaceEqual(expected, actual *schema.Interface) {
	assert := s.Assert()
	require := s.Require()

	assert.Equal(expected.Key(), actual.Key(), "wrong Interface key")
	assert.Equal(expected.Name, actual.Name, "wrong Interface name")
	require.Len(actual.Methods, len(expected.Methods), "wrong Interface methods count")

	for idx, expectedMethod := range expected.Methods {
		actualMethod := actual.Methods[idx]

		s.AssertMethodEqualf(expectedMethod, actualMethod, "wrong Method at index %d", idx)
	}
}

func (s *Suite) AssertInterfaceEqualf(expected, actual *schema.Interface, msg string, args ...any) {
	assert := s.Assert()
	require := s.Require()

	assert.Equalf(expected.Key(), actual.Key(), msg+": wrong Interface key", args...)
	assert.Equalf(expected.Name, actual.Name, msg+": wrong Interface name", args...)
	require.Lenf(actual.Methods, len(expected.Methods), msg+": wrong Interface methods count", args...)

	for idx, expectedMethod := range expected.Methods {
		actualMethod := actual.Methods[idx]

		s.AssertMethodEqualf(expectedMethod, actualMethod, msg+": wrong Method at index %d", append(args, idx)...)
	}
}

func (s *Suite) AssertInterfacesEqual(expected, actual schema.Interfaces) {
	require := s.Require()

	require.Len(actual, len(expected))

	for idx, expectedInterface := range expected {
		actualInterface := actual[idx]

		s.AssertInterfaceEqualf(expectedInterface, actualInterface, "wrong Interface at index %d", idx)
	}
}

func (s *Suite) AssertMethodEqualf(expected, actual *schema.Method, format string, args ...any) {
	assert := s.Assert()

	assert.Equalf(expected.Key(), actual.Key(), format+": wrong Method key", args...)
	assert.Equalf(expected.Name, actual.Name, format+": wrong Method name", args...)
	assert.Equalf(expected.Version, actual.Version, format+": wrong Method version", args...)
	assert.Equalf(expected.HTTPMethod, actual.HTTPMethod, format+": wrong Method http method", args...)
	assert.Equalf(expected.Params, actual.Params, format+": wrong Method params", args...)
	assert.Equalf(expected.Undocumented, actual.Undocumented, format+": wrong Method undocumented", args...)
}

func (s *Suite) AssertMethodsEqual(expected, actual schema.Methods) {
	require := s.Require()

	require.Len(actual, len(expected))

	for idx, expectedMethod := range expected {
		actualMethod := actual[idx]

		s.AssertMethodEqualf(expectedMethod, actualMethod, "wrong Method at index %d", idx)
	}
}

func (s *Suite) createInterface(name string, appID steamlib.AppID, methods schema.Methods) *schema.Interface {
	if appID != 0 {
		name = fmt.Sprintf("%s_%d", name, appID)
	}

	return schema.MustNewInterface(name, methods)
}

func (s *Suite) createInterfaceNoAppID() *schema.Interface {
	return s.createInterface("IFace", 0, nil)
}

func (s *Suite) createInterfaceAppID() *schema.Interface {
	return s.createInterface("IFace", 123, nil)
}

func (s *Suite) createInterfaceInvalidName() *schema.Interface {
	return &schema.Interface{Name: "invalid"}
}

func (s *Suite) createInterfaceMethods() *schema.Interface {
	return s.createInterface("IFace", 0, s.createMethods())
}

func (s *Suite) createInterfaceMethodsInvalidName() *schema.Interface {
	return &schema.Interface{Name: "IFace", Methods: s.createMethodsInvalidName()}
}

func (s *Suite) createInterfaceMethodsInvalidVersion() *schema.Interface {
	return &schema.Interface{Name: "IFace", Methods: s.createMethodsInvalidVersion()}
}

func (s *Suite) createInterfaceMethodsInvalidHTTPMethod() *schema.Interface {
	return &schema.Interface{Name: "IFace", Methods: s.createMethodsInvalidHTTPMethod()}
}

func (s *Suite) createInterfaces() schema.Interfaces {
	return schema.MustNewInterfaces(s.createInterfaceAppID())
}

func (s *Suite) createInterfacesAppID() schema.Interfaces {
	return schema.Interfaces{s.createInterfaceAppID()}
}

func (s *Suite) createInterfacesMethods() schema.Interfaces {
	return schema.Interfaces{s.createInterfaceMethods()}
}

func (s *Suite) createInterfacesInvalidName() schema.Interfaces {
	return schema.Interfaces{s.createInterfaceInvalidName()}
}

func (s *Suite) createInterfacesMethodsInvalidName() schema.Interfaces {
	return schema.Interfaces{s.createInterfaceMethodsInvalidName()}
}

func (s *Suite) createInterfacesMethodsInvalidVersion() schema.Interfaces {
	return schema.Interfaces{s.createInterfaceMethodsInvalidVersion()}
}

func (s *Suite) createInterfacesMethodsInvalidHTTPMethod() schema.Interfaces {
	return schema.Interfaces{s.createInterfaceMethodsInvalidHTTPMethod()}
}

func (s *Suite) createMethod(name string, version int) *schema.Method {
	return schema.MustNewMethod(name, version, http.MethodGet, nil, false)
}

func (s *Suite) createMethodValid() *schema.Method {
	return s.createMethod("Method", 1)
}

func (s *Suite) createMethodInvalidName() *schema.Method {
	return &schema.Method{Name: "9Method", Version: 1, HTTPMethod: http.MethodGet}
}

func (s *Suite) createMethodInvalidVersion() *schema.Method {
	return &schema.Method{Name: "Method", Version: 0, HTTPMethod: http.MethodGet}
}

func (s *Suite) createMethodInvalidHTTPMethod() *schema.Method {
	return &schema.Method{Name: "Method", Version: 1, HTTPMethod: "xyz"}
}

func (s *Suite) createMethodWithParams() *schema.Method {
	return schema.MustNewMethod(
		"MyMethod",
		1,
		http.MethodGet,
		s.createMethodParamsOptional("name", "type"),
		false,
	)
}

func (s *Suite) createMethodWithRequiredParams() *schema.Method {
	return schema.MustNewMethod(
		"MyMethod",
		1,
		http.MethodGet,
		s.createMethodParamsRequired("name", "type"),
		false,
	)
}

func (s *Suite) createMethods() schema.Methods {
	return schema.MustNewMethods(s.createMethodValid())
}

func (s *Suite) createMethodsInvalidName() schema.Methods {
	return schema.Methods{s.createMethodInvalidName()}
}

func (s *Suite) createMethodsInvalidVersion() schema.Methods {
	return schema.Methods{s.createMethodInvalidVersion()}
}

func (s *Suite) createMethodsInvalidHTTPMethod() schema.Methods {
	return schema.Methods{s.createMethodInvalidHTTPMethod()}
}

func (s *Suite) createMethodParamOptional(name, ty string) *schema.MethodParam {
	return schema.NewMethodParam(name, ty, true, "description")
}

func (s *Suite) createMethodParamRequired(name, ty string) *schema.MethodParam {
	return schema.NewMethodParam(name, ty, false, "description")
}

func (s *Suite) createMethodParamsOptional(name, ty string) schema.MethodParams {
	return schema.NewMethodParams(s.createMethodParamOptional(name, ty))
}

func (s *Suite) createMethodParamsRequired(name, ty string) schema.MethodParams {
	return schema.NewMethodParams(s.createMethodParamRequired(name, ty))
}
