package examples

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/13k/valve.go/steamres/steampb/underlords/steamworks"
)

func TestProtobufUnderlordsSteamworks(t *testing.T) {
	assert.Equal(
		t,
		"k_EProtoExecutionSiteSteamClient",
		steamworks.EProtoExecutionSite_k_EProtoExecutionSiteSteamClient.String(),
	)
}
