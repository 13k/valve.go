// Package dota2 implements an HTTP client for the (undocumented) Dota 2 Web API.
//
// For more information, refer to the main "steamweb" package documentation.
package dota2

import (
	"github.com/13k/valve.go/steamweb"
)

const (
	// HostURL is the base URL for the Dota 2 Web API client.
	HostURL = "https://www.dota2.com/webapi"
)

var _ steamweb.Client = (*Client)(nil)

// Client is a thin wrapper around a base `steamweb.Client` that automatically sets the base
// `HostURL`.
type Client struct {
	steamweb.Client
}

// New creates a new Dota2 API client.
func New(options ...steamweb.ClientOption) (*Client, error) {
	options = append(options, steamweb.WithHostURL(HostURL))
	base, err := steamweb.New(options...)

	if err != nil {
		return nil, err
	}

	client := &Client{Client: base}

	return client, nil
}
