package examples

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/13k/valve.go/steamres/steampb/steam"
	"github.com/13k/valve.go/steamres/steampb/steam/client"
	"github.com/13k/valve.go/steamres/steampb/tf2"
)

func TestProtobufTf2(t *testing.T) {
	assert.Equal(t, int32(2), steam.Default_CMsgProtoBufHeader_Eresult)
	assert.Equal(t, "k_EProtoExecutionSiteSteamClient", client.EProtoExecutionSite_k_EProtoExecutionSiteSteamClient.String())
	assert.Equal(t, "k_EMsgGCReportWarKill", tf2.ETFGCMsg_k_EMsgGCReportWarKill.String())
}
