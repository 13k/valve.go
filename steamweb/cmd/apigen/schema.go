package main

import (
	"bufio"
	"encoding/json"
	"io/fs"
	"log"
	"net/url"
	"os"
	"sort"

	"github.com/go-errors/errors"
	"github.com/go-resty/resty/v2"

	"github.com/13k/valve.go/steamweb/schema"
)

type SchemaKind uint8

const (
	SchemaKindOfficial SchemaKind = iota
	SchemaKindXpaw
)

const (
	warnfSchemaMergeMethodConflict = "warning: Schema.Merge(): method %s/%s/%d already exists on destination schema"
)

type Schema struct {
	Schema *schema.Schema `json:"apilist"`

	relPath     string
	pkgPath     string
	pkgName     string
	filenames   map[string]string
	keyedByName map[string]*schema.Interface
}

func GetSchema(cacheFile string, apiKey string) (*Schema, error) {
	var (
		schema *Schema
		err    error
	)

	if cacheFile != "" {
		schema = &Schema{}

		if err = schema.Load(cacheFile); err != nil {
			if errors.Is(err, fs.ErrNotExist) {
				schema = nil
			} else {
				return nil, errors.Wrap(err, 0)
			}
		}
	}

	if schema == nil {
		schema, err = fetchSchema(apiKey)

		if err != nil {
			return nil, errors.Wrap(err, 0)
		}

		if cacheFile != "" {
			if err = schema.Dump(cacheFile); err != nil {
				return nil, errors.Wrap(err, 0)
			}
		}
	}

	return schema, nil
}

func newSchema(ss *schema.Schema) (*Schema, error) {
	spec := schemaSpecs["steam"]

	schema := &Schema{
		Schema:      ss,
		relPath:     spec.RelPath,
		pkgPath:     spec.PkgPath,
		pkgName:     spec.PkgName,
		filenames:   spec.Filenames,
		keyedByName: nil,
	}

	if err := schema.build(); err != nil {
		return nil, errors.Wrap(err, 0)
	}

	return schema, nil
}

func fetchSchema(apiKey string) (*Schema, error) {
	req, err := newSchemaReqSteam(apiKey)

	if err != nil {
		return nil, errors.Wrap(err, 0)
	}

	client := resty.New()
	schema, err := req.Fetch(client)

	if err != nil {
		return nil, errors.Wrap(err, 0)
	}

	req, err = newSchemaReqXpaw()

	if err != nil {
		return nil, errors.Wrap(err, 0)
	}

	schemaXpaw, err := req.Fetch(client)

	if err != nil {
		return nil, errors.Wrap(err, 0)
	}

	if err := schema.Merge(schemaXpaw); err != nil {
		return nil, errors.Wrap(err, 0)
	}

	return schema, nil
}

func parseSchema(kind SchemaKind, data []byte) (*Schema, error) {
	var (
		ss  *schema.Schema
		err error
	)

	switch kind {
	case SchemaKindOfficial:
		ss, err = parseSchemaOfficial(data)
	case SchemaKindXpaw:
		ss, err = parseSchemaXpaw(data)
	}

	if err != nil {
		return nil, errors.Wrap(err, 0)
	}

	schema, err := newSchema(ss)

	if err != nil {
		return nil, errors.Wrap(err, 0)
	}

	return schema, nil
}

func (s *Schema) build() error {
	s.normalize()

	if err := s.validate(); err != nil {
		return errors.Wrap(err, 0)
	}

	if err := s.buildIndex(); err != nil {
		return errors.Wrap(err, 0)
	}

	return nil
}

func (s *Schema) normalize() {
	for _, i := range s.Schema.Interfaces {
		for _, m := range i.Methods {
			if m.HTTPMethod == "" {
				m.HTTPMethod = "GET"
			}
		}
	}
}

func (s *Schema) validate() error {
	if err := s.Schema.Validate(); err != nil {
		return errors.Wrap(err, 0)
	}

	return nil
}

func (s *Schema) buildIndex() error {
	if s.keyedByName != nil {
		return nil
	}

	s.keyedByName = make(map[string]*schema.Interface)

	for _, sif := range s.Schema.Interfaces {
		if _, ok := s.keyedByName[sif.Name]; ok {
			return errors.Errorf("schema contains duplicate interfaces (%q)", sif.Name)
		}

		s.keyedByName[sif.Name] = sif
	}

	return nil
}

func (s *Schema) Filename(group schema.InterfacesIndex) string {
	return s.filenames[group.Name()]
}

func (s *Schema) Merge(other *Schema) error {
	if err := s.buildIndex(); err != nil {
		return errors.Wrap(err, 0)
	}

	if err := other.buildIndex(); err != nil {
		return errors.Wrap(err, 0)
	}

	var result []*schema.Interface

	// append all interfaces in `s` that are not in `other`
	for _, si := range s.Schema.Interfaces {
		if _, ok := other.keyedByName[si.Name]; !ok {
			result = append(result, si)
		}
	}

	for _, oi := range other.Schema.Interfaces {
		si := s.keyedByName[oi.Name]

		// append interface in `other` that is not in `s`
		if si == nil {
			result = append(result, oi)
			continue
		}

		// merge interface that belongs to both

		var oMethods []*schema.Method

		// append all methods in `other` that are not in `s`
		for _, om := range oi.Methods {
			if _, err := si.Methods.Get(om.Key()); err == nil {
				// ignore conflicting method (keep method in `s`)
				log.Printf(warnfSchemaMergeMethodConflict, oi.Name, om.Name, om.Version)
				continue
			}

			oMethods = append(oMethods, om)
		}

		// prepend all methods in `s`
		methods, err := schema.NewMethods(append(si.Methods, oMethods...)...)

		if err != nil {
			return errors.Wrap(err, 0)
		}

		// create a new interface with merged methods
		mergedInterface := &schema.Interface{
			Name:    si.Name,
			Methods: methods,
		}

		result = append(result, mergedInterface)
	}

	interfaces, err := schema.NewInterfaces(result...)

	if err != nil {
		return errors.Wrap(err, 0)
	}

	s.Schema.Interfaces = interfaces

	return nil
}

func (s *Schema) Dump(path string) error {
	log.Printf("Writing schema to %q", path)

	f, err := os.Create(path)

	if err != nil {
		return errors.Wrap(err, 0)
	}

	defer f.Close()

	buf := bufio.NewWriter(f)

	defer buf.Flush()

	if err := json.NewEncoder(buf).Encode(s); err != nil {
		return errors.Wrap(err, 0)
	}

	return nil
}

func (s *Schema) Load(path string) error {
	f, err := os.Open(path)

	if err != nil {
		return errors.Wrap(err, 0)
	}

	defer f.Close()

	log.Printf("Loading schema from %q", path)

	buf := bufio.NewReader(f)

	if err := json.NewDecoder(buf).Decode(s); err != nil {
		return errors.Wrap(err, 0)
	}

	return nil
}

type interfaceGroupIterator func(string, schema.InterfacesIndex) error

func (s *Schema) eachSortedInterfaceGroup(fn interfaceGroupIterator) error {
	groups := s.Schema.Interfaces.GroupByBaseName()
	groupNames := make([]string, 0, len(groups))

	for groupName := range groups {
		groupNames = append(groupNames, groupName)
	}

	sort.Strings(groupNames)

	for _, groupName := range groupNames {
		group := groups[groupName]

		if err := fn(groupName, group); err != nil {
			return errors.Wrap(err, 0)
		}
	}

	return nil
}

type SchemaRequest struct {
	Kind   SchemaKind
	URL    string
	Params url.Values
}

func (req *SchemaRequest) Fetch(client *resty.Client) (*Schema, error) {
	log.Printf("Fetching schema from %q", req.URL)

	resp, err := client.R().
		SetQueryParamsFromValues(req.Params).
		Get(req.URL)

	if err != nil {
		return nil, errors.Wrap(err, 0)
	}

	if !resp.IsSuccess() {
		return nil, errors.Errorf("error fetching remote schema: %s", resp.Status())
	}

	body := resp.Body()
	schema, err := parseSchema(req.Kind, body)

	if err != nil {
		return nil, errors.Wrap(err, 0)
	}

	return schema, nil
}
