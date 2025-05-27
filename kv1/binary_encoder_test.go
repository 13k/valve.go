package kv1_test

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/13k/valve.go/kv1"
)

func TestBinaryEncoder(t *testing.T) {
	t.Parallel()

	suite.Run(t, new(BinaryEncoderSuite))
}

type BinaryEncoderSuite struct {
	Suite
}

type binEncTestCase struct {
	Name     string
	Subject  *kv1.KeyValue
	Err      string
	Expected []byte
}

func (s *BinaryEncoderSuite) TestEncode() {
	testCases := []binEncTestCase{
		{
			Name:    "type_unsupported_end",
			Subject: kv1.MustNewKeyValue(kv1.TypeEnd, "", "", nil),
			Err:     "type is not supported",
		},
		{
			Name:    "type_unsupported_wstring",
			Subject: kv1.MustNewKeyValue(kv1.TypeWString, "foo", "", nil),
			Err:     "type is not supported",
		},
		{
			Name:    "type_unsupported_invalid",
			Subject: kv1.MustNewKeyValue(kv1.TypeInvalid, "foo", "", nil),
			Err:     "type is not supported",
		},
		{
			Name:    "string_ok",
			Subject: kv1.NewKeyValueString("K", "S", nil),
			Expected: []byte{
				kv1.TypeString.Byte(), // type
				'K', 0x00,             // key
				'S', 0x00, // string value
			},
		},
		{
			Name:    "int32_ok",
			Subject: kv1.NewKeyValueInt32("K", 1, nil),
			Expected: []byte{
				kv1.TypeInt32.Byte(), // type
				'K', 0x00,            // key
				0x01, 0x00, 0x00, 0x00, // int32 value
			},
		},
		{
			Name:    "color_ok",
			Subject: kv1.NewKeyValueColor("K", 1, nil),
			Expected: []byte{
				kv1.TypeColor.Byte(), // type
				'K', 0x00,            // key
				0x01, 0x00, 0x00, 0x00, // int32 value
			},
		},
		{
			Name:    "pointer_ok",
			Subject: kv1.NewKeyValuePointer("K", 1, nil),
			Expected: []byte{
				kv1.TypePointer.Byte(), // type
				'K', 0x00,              // key
				0x01, 0x00, 0x00, 0x00, // int32 value
			},
		},
		{
			Name:    "int64_ok",
			Subject: kv1.NewKeyValueInt64("K", 1, nil),
			Expected: []byte{
				kv1.TypeInt64.Byte(), // type
				'K', 0x00,            // key
				0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // int64 value
			},
		},
		{
			Name:    "uint64_ok",
			Subject: kv1.NewKeyValueUint64("K", 1, nil),
			Expected: []byte{
				kv1.TypeUint64.Byte(), // type
				'K', 0x00,             // key
				0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // uint64 value
			},
		},
		{
			Name:    "object_ok",
			Subject: kv1.NewKeyValueRoot("K").AddString("s", "S"),
			Expected: []byte{
				kv1.TypeObject.Byte(), // type
				'K', 0x00,             // key
				0x01,      // child type
				's', 0x00, // child key
				'S', 0x00, // child string value
				kv1.TypeEnd.Byte(), // end type
			},
		},
		{
			Name: "bindata1",
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
			Expected: binData1,
		},
		{
			Name: "bindata2",
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
			Expected: binData2,
		},
		{
			Name: "bindata3",
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
			Expected: binData3,
		},
		{
			Name: "bindata4",
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
			Expected: binData4,
		},
	}

	for _, testCase := range testCases {
		s.subtestEncode(testCase)
	}
}

func (s *BinaryEncoderSuite) subtestEncode(testCase binEncTestCase) {
	s.Run(testCase.Name, func() {
		require := s.Require()
		b := &bytes.Buffer{}
		enc := kv1.NewBinaryEncoder(b)
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
