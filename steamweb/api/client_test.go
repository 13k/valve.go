package api_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/13k/valve.go/steamweb/api"
)

func TestNew(t *testing.T) {
	client, err := api.New()

	require.NoError(t, err)
	require.NotNil(t, client)
	require.Equal(t, api.HostURL, client.HostURL())
}
