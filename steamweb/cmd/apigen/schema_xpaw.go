package main

import (
	"encoding/json"

	"github.com/go-errors/errors"

	"github.com/13k/valve.go/steamweb/schema"
)

const (
	schemaURLXpaw = "https://github.com/xPaw/SteamWebAPIDocumentation/raw/refs/heads/master/api.json"
)

func newSchemaReqXpaw() (*SchemaRequest, error) {
	req := &SchemaRequest{
		Kind:   SchemaKindXpaw,
		URL:    schemaURLXpaw,
		Params: nil,
	}

	return req, nil
}

func parseSchemaXpaw(data []byte) (*schema.Schema, error) {
	xs := make(schemaXpaw)

	if err := json.Unmarshal(data, &xs); err != nil {
		return nil, errors.Wrap(err, 0)
	}

	ss, err := xs.ToSchema()

	if err != nil {
		return nil, errors.Wrap(err, 0)
	}

	return ss, nil
}

// interface_name: interface
type schemaXpaw map[string]*schemaXpawInterface

func (xs schemaXpaw) ToSchema() (*schema.Schema, error) {
	interfaces := make(schema.Interfaces, 0, len(xs))

	for name, xi := range xs {
		si, err := xi.ToInterface(name)

		if err != nil {
			return nil, errors.Wrap(err, 0)
		}

		interfaces = append(interfaces, si)
	}

	ss := &schema.Schema{Interfaces: interfaces}

	return ss, nil
}

// method_name: method
type schemaXpawInterface map[string]*schemaXpawMethod

func (xi schemaXpawInterface) ToInterface(name string) (*schema.Interface, error) {
	si := &schema.Interface{
		Name:    name,
		Methods: make(schema.Methods, 0, len(xi)),
	}

	for name, xm := range xi {
		sm, err := xm.ToMethod(name)

		if err != nil {
			return nil, errors.Wrap(err, 0)
		}

		si.Methods = append(si.Methods, sm)
	}

	return si, nil
}

type schemaXpawMethod struct {
	Type       string                   `json:"_type"`
	Version    int                      `json:"version"`
	HTTPMethod string                   `json:"httpmethod"`
	Params     []*schemaXpawMethodParam `json:"parameters"`
}

func (xm *schemaXpawMethod) ToMethod(name string) (*schema.Method, error) {
	sm := &schema.Method{
		Name:         name,
		Version:      xm.Version,
		HTTPMethod:   xm.HTTPMethod,
		Params:       make(schema.MethodParams, len(xm.Params)),
		Undocumented: xm.Type == "undocumented",
	}

	for i, xmp := range xm.Params {
		smp, err := xmp.ToMethodParam()

		if err != nil {
			return nil, errors.Wrap(err, 0)
		}

		sm.Params[i] = smp
	}

	return sm, nil
}

type schemaXpawMethodParam struct {
	Name        string                   `json:"name"`
	Type        string                   `json:"type"`
	Optional    bool                     `json:"optional"`
	Description string                   `json:"description"`
	Extra       []*schemaXpawMethodParam `json:"extra"`
}

func (xmp *schemaXpawMethodParam) ToMethodParam() (*schema.MethodParam, error) {
	smp := &schema.MethodParam{
		Name:        xmp.Name,
		Type:        xmp.Type,
		Optional:    xmp.Optional,
		Description: xmp.Description,
	}

	return smp, nil
}
