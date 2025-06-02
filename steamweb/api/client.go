// Package steam implements an HTTP client for the Steam Web API.
//
// For more information, refer to the main "steamweb" package documentation.
package api

import (
	"github.com/13k/valve.go/steamweb"
)

const (
	// HostURL is the base URL for the Steam Web API client.
	HostURL = "https://api.steampowered.com"
)

var _ steamweb.Client = (*Client)(nil)

// Client is a thin wrapper around a base `steamweb.Client` that automatically sets the base
// `HostURL`.
type Client struct {
	steamweb.Client
}

// New creates a new Steam API client.
func New(options ...steamweb.ClientOption) (*Client, error) {
	options = append(options, steamweb.WithHostURL(HostURL))
	base, err := steamweb.New(options...)

	if err != nil {
		return nil, err
	}

	client := &Client{Client: base}

	return client, nil
}
