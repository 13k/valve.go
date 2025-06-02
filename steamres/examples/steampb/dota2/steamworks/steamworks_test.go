package examples

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/13k/valve.go/steamres/steampb/dota2/steamworks"
)

func TestProtobufDota2Steamworks(t *testing.T) {
	assert.Equal(
		t,
		"k_EProtoExecutionSiteSteamClient",
		steamworks.EProtoExecutionSite_k_EProtoExecutionSiteSteamClient.String(),
	)
}
