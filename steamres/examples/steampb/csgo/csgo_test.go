package examples

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/13k/valve.go/steamres/steampb/csgo"
	"github.com/13k/valve.go/steamres/steampb/steam"
	"github.com/13k/valve.go/steamres/steampb/steam/client"
)

func TestProtobufCsgo(t *testing.T) {
	assert.Equal(t, int32(2), steam.Default_CMsgProtoBufHeader_Eresult)
	assert.Equal(t, "k_EProtoExecutionSiteSteamClient", client.EProtoExecutionSite_k_EProtoExecutionSiteSteamClient.String())
	assert.Equal(t, "k_EMsgGCCStrike15_v2_MatchmakingStart", csgo.ECsgoGCMsg_k_EMsgGCCStrike15_v2_MatchmakingStart.String())
}
