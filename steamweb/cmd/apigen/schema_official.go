package main

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/go-errors/errors"

	"github.com/13k/valve.go/steamweb/schema"
)

const (
	schemaURLOfficial = "https://api.steampowered.com/ISteamWebAPIUtil/GetSupportedAPIList/v1"
)

var (
	errMissingSteamAPIKey = errors.Errorf("API key is required when fetching remote Steam API schema")
)

func newSchemaReqSteam(apiKey string) (*SchemaRequest, error) {
	if apiKey == "" {
		return nil, errors.New(errMissingSteamAPIKey)
	}

	params := url.Values{}
	params.Set("key", apiKey)

	req := &SchemaRequest{
		Kind:   SchemaKindOfficial,
		URL:    schemaURLOfficial,
		Params: params,
	}

	return req, nil
}

func parseSchemaOfficial(data []byte) (*schema.Schema, error) {
	s := &schemaOfficial{}

	if err := json.Unmarshal(data, s); err != nil {
		fmt.Println(string(data))

		return nil, errors.Wrap(err, 0)
	}

	return s.Schema, nil
}

type schemaOfficial struct {
	Schema *schema.Schema `json:"apilist"`
}
