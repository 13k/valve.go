package examples

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/13k/valve.go/steamres/steampb/dota2/steam/network"
)

func TestProtobufDota2SteamNetwork(t *testing.T) {
	assert.Equal(
		t,
		"k_ESteamDatagramMsg_ConnectRequest",
		network.ESteamDatagramMsgID_k_ESteamDatagramMsg_ConnectRequest.String(),
	)
}
