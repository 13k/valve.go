package kv1_test

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/13k/valve.go/kv1"
)

func TestTextEncoder(t *testing.T) {
	t.Parallel()

	suite.Run(t, new(TextEncoderSuite))
}

type TextEncoderSuite struct {
	Suite
}

type textEncTestCase struct {
	Name     string
	Subject  *kv1.KeyValue
	Expected []byte
	Err      string
}

func (s *TextEncoderSuite) TestEncode() {
	testCases := []textEncTestCase{
		{
			Name:    "type_unsupported_end",
			Subject: kv1.MustNewKeyValue(kv1.TypeEnd, "eof", "", nil),
			Err:     `type is not supported`,
		},
		{
			Name:    "type_unsupported_invalid",
			Subject: kv1.MustNewKeyValue(kv1.TypeInvalid, "foo", "", nil),
			Err:     `type is not supported`,
		},
		{
			Name:     "string_ok",
			Subject:  kv1.NewKeyValueString("K", "S", nil),
			Expected: []byte(`"K" "S"`),
		},
		{
			Name:     "int32_ok",
			Subject:  kv1.NewKeyValueInt32("K", 1, nil),
			Expected: []byte(`"K" "1"`),
		},
		{
			Name:     "color_ok",
			Subject:  kv1.NewKeyValueColor("K", 1, nil),
			Expected: []byte(`"K" "1"`),
		},
		{
			Name:     "pointer_ok",
			Subject:  kv1.NewKeyValuePointer("K", 1, nil),
			Expected: []byte(`"K" "1"`),
		},
		{
			Name:     "int64_ok",
			Subject:  kv1.NewKeyValueInt64("K", 1, nil),
			Expected: []byte(`"K" "1"`),
		},
		{
			Name:     "uint64_ok",
			Subject:  kv1.NewKeyValueUint64("K", 1, nil),
			Expected: []byte(`"K" "1"`),
		},
		{
			Name:     "float32_ok",
			Subject:  kv1.NewKeyValueFloat32("K", 1.23, nil),
			Expected: []byte(`"K" "1.23"`),
		},
		{
			Name:    "object_ok",
			Subject: kv1.NewKeyValueRoot("K").AddString("s", "S"),
			Expected: []byte(
				`"K" {
  "s" "S"
}
`,
			),
		},
		{
			Name: "textdata1",
			Subject: kv1.NewKeyValueRoot("RP").
				AddString("status", "#DOTA_RP_LEAGUE_MATCH_PLAYING_AS").
				AddString("steam_display", "#DOTA_RP_LEAGUE_MATCH_PLAYING_AS").
				AddString("num_params", "3").
				AddString("lobby", "lobby_id: 26628083760328052 lobby_state: RUN password: true game_mode: DOTA_GAMEMODE_CM member_count: 10 max_member_count: 10 name: \"Team Secret vs Vikin.GG \" lobby_type: 1").
				AddString("party", "party_state: IN_MATCH").
				AddString("WatchableGameID", "26628083760328052").
				AddString("param0", "#DOTA_lobby_type_name_lobby").
				AddString("param1", "8").
				AddString("param2", "#npc_dota_hero_grimstroke"),
			Expected: textData1,
		},
		{
			Name: "textdata2",
			Subject: kv1.NewKeyValueRoot("RP").
				AddString("status", "#DOTA_RP_PLAYING_AS").
				AddString("steam_display", "#DOTA_RP_PLAYING_AS").
				AddString("num_params", "3").
				AddString("CustomGameMode", "0").
				AddString("WatchableGameID", "26628083785444387").
				AddString("party", "party_id: 26628083781803523 party_state: IN_MATCH open: false members { steam_id: 76561198054320440 }").
				AddString("param0", "#DOTA_lobby_type_name_ranked").
				AddString("param1", "15").
				AddString("param2", "#npc_dota_hero_zuus"),
			Expected: textData2,
		},
		{
			Name: "textdata3",
			Subject: kv1.NewKeyValueRoot("RP").
				AddString("status", "#DOTA_RP_PLAYING_AS").
				AddString("steam_display", "#DOTA_RP_PLAYING_AS").
				AddString("num_params", "3").
				AddString("watching_server", "[A:1:3441750017:14553]").
				AddString("watching_from_server", "[A:1:1671049217:14554]").
				AddString("party", "party_state: IN_MATCH").
				AddString("WatchableGameID", "26628083799951455").
				AddString("steam_player_group", "26628083752106249").
				AddString("steam_player_group_size", "2").
				AddString("param0", "#DOTA_lobby_type_name_ranked").
				AddString("param1", "2").
				AddString("param2", "#npc_dota_hero_rubick"),
			Expected: textData3,
		},
		{
			Name: "textdata4",
			Subject: kv1.NewKeyValueRoot("RP").
				AddString("status", "#DOTA_RP_HERO_SELECTION").
				AddString("steam_display", "#DOTA_RP_HERO_SELECTION").
				AddString("num_params", "1").
				AddString("WatchableGameID", "26628083824762603").
				AddString("steam_player_group", "26628083765767134").
				AddString("steam_player_group_size", "2").
				AddString("party", "party_id: 26628083765767134 party_state: IN_MATCH open: false members { steam_id: 76561198235766844 } members { steam_id: 76561197978446698 }").
				AddString("watching_server", "[A:1:300033030:14554]").
				AddString("watching_from_server", "[A:1:1361739785:14554]").
				AddString("param0", "#DOTA_lobby_type_name_ranked"),
			Expected: textData4,
		},
	}

	for _, testCase := range testCases {
		s.subtestEncode(testCase)
	}
}

func (s *TextEncoderSuite) subtestEncode(testCase textEncTestCase) {
	s.Run(testCase.Name, func() {
		require := s.Require()
		b := &bytes.Buffer{}
		enc := kv1.NewTextEncoder(b)
		err := enc.Encode(testCase.Subject)
		actual := b.Bytes()

		if testCase.Err == "" {
			require.NoError(err)
			require.Equal(testCase.Expected, actual)
		} else {
			require.ErrorContains(err, testCase.Err)
		}
	})
}
