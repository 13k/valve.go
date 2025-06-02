package main

import (
	"net/http"
	"strings"

	j "github.com/dave/jennifer/jen"
)

var (
	jenImportNames = map[string]string{
		pkgPathTestifyAssert:  pkgNameTestifyAssert,
		pkgPathTestifyRequire: pkgNameTestifyRequire,

		pkgBasePath:           pkgNameRoot,
		pkgPathSteamwebSchema: pkgNameSchema,
		pkgPathSteamwebSteam:  pkgNameSteam,
		pkgPathSteamwebDota2:  pkgNameDota2,
	}
)

func jFile(pkgPath, pkgName string) *j.File {
	f := j.NewFilePathName(pkgPath, pkgName)

	f.ImportNames(jenImportNames)

	return f
}

//// Go statements

func jVarErrError() *j.Statement {
	return j.Var().Err().Error()
}

func jIfErrRet(values ...j.Code) *j.Statement {
	return j.If(j.Err().Op("!=").Nil()).Block(j.Return(values...))
}

func jIfErrRetNilErr() *j.Statement {
	return jIfErrRet(j.Nil(), j.Err())
}

type jBagItem struct {
	ID   string
	Stmt j.Code
}

type jBag []jBagItem

func (bag *jBag) Declare(id string, stmt j.Code) bool {
	for _, item := range *bag {
		if item.ID == id {
			return false
		}
	}

	*bag = append(*bag, jBagItem{ID: id, Stmt: stmt})

	return true
}

func (bag jBag) Codes() []j.Code {
	if bag == nil {
		return nil
	}

	s := make([]j.Code, len(bag))

	for i, item := range bag {
		s[i] = item.Stmt
	}

	return s
}

//// package http statements

func jHTTPMethod(method string) *j.Statement {
	method = strings.ToUpper(method)

	switch method {
	case http.MethodGet:
		return j.Qual(pkgPathNetHTTP, "MethodGet")
	case http.MethodPost:
		return j.Qual(pkgPathNetHTTP, "MethodPost")
	case http.MethodPut:
		return j.Qual(pkgPathNetHTTP, "MethodPut")
	default:
		return nil
	}
}

//// package testing statements

func jTestingTIdPtr(id string) *j.Statement {
	return j.Id(id).Op("*").Qual(pkgPathTesting, "T")
}

func jTestifyAssert(assertion, tID string, args ...j.Code) *j.Statement { //nolint:unparam
	args = append([]j.Code{j.Id(tID)}, args...)
	return j.Qual(pkgPathTestifyAssert, assertion).Call(args...)
}

func jTestifyRequire(assertion, tID string, args ...j.Code) *j.Statement { //nolint:unparam
	args = append([]j.Code{j.Id(tID)}, args...)
	return j.Qual(pkgPathTestifyRequire, assertion).Call(args...)
}

//// package steamweb statements

func jSteamwebID(id string) *j.Statement {
	return j.Qual(pkgBasePath, id)
}

func jSteamwebTypeOp(op, id string) *j.Statement {
	return j.Op(op).Add(jSteamwebID(id))
}

func jSteamwebTypePtr(id string) *j.Statement {
	return jSteamwebTypeOp("*", id)
}

func jRequestCtorID() *j.Statement {
	return jSteamwebID(srcRequestCtor)
}

func jRequestPtr() *j.Statement {
	return jSteamwebTypePtr(srcRequest)
}

//// packages that contain a "Client" identifier

func jClientID(pkgPath string) *j.Statement {
	if pkgPath != "" {
		return j.Qual(pkgPath, srcClient)
	}

	return j.Id(srcClient)
}

func jClientCtorID(pkgPath string) *j.Statement {
	if pkgPath != "" {
		return j.Qual(pkgPath, srcClientCtor)
	}

	return j.Id(srcClientCtor)
}

func jClientOp(op, pkgPath string) *j.Statement {
	return j.Op(op).Add(jClientID(pkgPath))
}

func jClientPtr(pkgPath string) *j.Statement {
	return jClientOp("*", pkgPath)
}

//// package schema statements

func jGeyserSchemaID(id string) *j.Statement {
	return j.Qual(pkgPathSteamwebSchema, id)
}

func jGeyserSchemaTypeOp(op, id string) *j.Statement {
	return j.Op(op).Add(jGeyserSchemaID(id))
}

func jGeyserSchemaTypeAddr(id string) *j.Statement {
	return jGeyserSchemaTypeOp("&", id)
}

func jGeyserSchemaTypePtr(id string) *j.Statement {
	return jGeyserSchemaTypeOp("*", id)
}

func jSchemaInterfaceAddr() *j.Statement {
	return jGeyserSchemaTypeAddr(srcSchemaInterface)
}

func jSchemaInterfacePtr() *j.Statement {
	return jGeyserSchemaTypePtr(srcSchemaInterface)
}

func jSchemaMethodAddr() *j.Statement {
	return jGeyserSchemaTypeAddr(srcSchemaMethod)
}

func jSchemaMethodParamAddr() *j.Statement {
	return jGeyserSchemaTypeAddr(srcSchemaMethodParam)
}

func jSchemaInterfacesCtorID() *j.Statement {
	return jGeyserSchemaID(srcSchemaInterfacesCtor)
}

func jSchemaMethodsCtorID() *j.Statement {
	return jGeyserSchemaID(srcSchemaMethodsCtor)
}

func jSchemaMethodParamsCtor() *j.Statement {
	return jGeyserSchemaID(srcSchemaMethodParamsCtor)
}

func jSchemaInterfaceKey(name string, appID j.Code) *j.Statement {
	keyValues := j.Dict{
		j.Id("Name"): j.Lit(name),
	}

	if appID != nil {
		keyValues[j.Id("AppID")] = appID
	}

	return jGeyserSchemaID(srcSchemaInterfaceKey).Values(keyValues)
}

func jSchemaMethodKey(name string, version j.Code) *j.Statement {
	return jGeyserSchemaID(srcSchemaMethodKey).Values(j.Dict{
		j.Id("Name"):    j.Lit(name),
		j.Id("Version"): version,
	})
}

func jInterfaceMethodNotFoundErrorPtr() *j.Statement {
	return jGeyserSchemaTypePtr(srcIntefaceMethodNotFoundError)
}
