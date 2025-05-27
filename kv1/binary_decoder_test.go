package kv1_test

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/13k/valve.go/kv1"
)

func TestBinaryDecoder(t *testing.T) {
	t.Parallel()

	suite.Run(t, new(BinaryDecoderSuite))
}

type BinaryDecoderSuite struct {
	Suite
}

type binDecTestCase struct {
	Name     string
	Data     []byte
	Err      string
	Expected *kv1.KeyValue
}

func (s *BinaryDecoderSuite) TestDecode() {
	testCases := []binDecTestCase{
		{
			Name: "data_nil",
			Err:  `binary decode error: failed to read node type: EOF`,
		},
		{
			Name: "data_empty",
			Data: []byte{},
			Err:  `binary decode error: failed to read node type: EOF`,
		},
		{
			Name: "end",
			Data: []byte{
				kv1.TypeEnd.Byte(), // type
			},
			Expected: kv1.MustNewKeyValue(kv1.TypeEnd, "", int32(0), nil),
		},
		{
			Name: "string_no_key",
			Data: []byte{
				kv1.TypeString.Byte(), // type
			},
			Err: `binary decode error: failed to decode value with type String: failed to read node key: failed to read string value: EOF`,
		},
		{
			Name: "string_key_incomplete",
			Data: []byte{
				kv1.TypeString.Byte(), // type
				'K',                   // non-terminated key
			},
			Err: `binary decode error: failed to decode value with type String: failed to read node key: failed to read string value: EOF`,
		},
		{
			Name: "string_no_value",
			Data: []byte{
				kv1.TypeString.Byte(), // type
				'K', 0x00,             // key
				// missing value
			},
			Err: `binary decode error: failed to decode key "K" with type String: failed to read string value: EOF`,
		},
		{
			Name: "string_value_incomplete",
			Data: []byte{
				kv1.TypeString.Byte(),
				'K', 0x00, // key
				'V', // non-terminated string value
			},
			Err: `binary decode error: failed to decode key "K" with type String: failed to read string value: EOF`,
		},
		{
			Name: "string_ok",
			Data: []byte{
				kv1.TypeString.Byte(),
				'K', 0x00, // key
				'V', 0x00, // string value
			},
			Expected: kv1.MustNewKeyValue(kv1.TypeString, "K", "V", nil),
		},
		{
			Name: "int32_incomplete",
			Data: []byte{
				kv1.TypeInt32.Byte(), // type
				'K', 0x00,            // key
				0x01, 0x00, 0x00, // incomplete int32 value
			},
			Err: `binary decode error: failed to decode key "K" with type Int32: failed to read int32 value: unexpected EOF`,
		},
		{
			Name: "int32_ok",
			Data: []byte{
				kv1.TypeInt32.Byte(), // type
				'K', 0x00,            // key
				0x01, 0x00, 0x00, 0x00, // int32 value
			},
			Expected: kv1.MustNewKeyValue(kv1.TypeInt32, "K", int32(1), nil),
		},
		{
			Name: "object_incomplete",
			Data: []byte{
				kv1.TypeObject.Byte(), // type
				'K', 0x00,             // key
				0x01,      // child string type
				's', 0x00, // child key
				'S', // child non-terminated string value
			},
			Err: `binary decode error: failed to decode key "K" with type Object: failed to read object value: binary decode error: failed to decode key "s" with type String: failed to read string value: EOF`,
		},
		{
			Name: "object_ok",
			Data: []byte{
				kv1.TypeObject.Byte(), // type
				'K', 0x00,             // key
				0x01,      // child string type
				's', 0x00, // child key
				'S', 0x00, // child string value
				kv1.TypeEnd.Byte(), // end
			},
			Expected: kv1.NewKeyValueRoot("K").AddString("s", "S"),
		},
		{
			Name: "bindata1",
			Data: binData1,
			Expected: kv1.NewKeyValueRoot("RP").
				AddString("status", "#DOTA_RP_LEAGUE_MATCH_PLAYING_AS").
				AddString("steam_display", "#DOTA_RP_LEAGUE_MATCH_PLAYING_AS").
				AddString("num_params", "3").
				AddString("lobby", "lobby_id: 26628083760328052 lobby_state: RUN password: true game_mode: DOTA_GAMEMODE_CM member_count: 10 max_member_count: 10 name: \"Team Secret vs Vikin.GG \" lobby_type: 1").
				AddString("party", "party_state: IN_MATCH").
				AddString("WatchableGameID", "26628083760328052").
				AddString("param0", "#DOTA_lobby_type_name_lobby").
				AddString("param1", "8").
				AddString("param2", "#npc_dota_hero_grimstroke"),
		},
		{
			Name: "bindata2",
			Data: binData2,
			Expected: kv1.NewKeyValueRoot("RP").
				AddString("status", "#DOTA_RP_PLAYING_AS").
				AddString("steam_display", "#DOTA_RP_PLAYING_AS").
				AddString("num_params", "3").
				AddString("CustomGameMode", "0").
				AddString("WatchableGameID", "26628083785444387").
				AddString("party", "party_id: 26628083781803523 party_state: IN_MATCH open: false members { steam_id: 76561198054320440 }").
				AddString("param0", "#DOTA_lobby_type_name_ranked").
				AddString("param1", "15").
				AddString("param2", "#npc_dota_hero_zuus"),
		},
		{
			Name: "bindata3",
			Data: binData3,
			Expected: kv1.NewKeyValueRoot("RP").
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
		},
		{
			Name: "bindata4",
			Data: binData4,
			Expected: kv1.NewKeyValueRoot("RP").
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
		},
	}

	for _, testCase := range testCases {
		s.subtestDecode(testCase)
	}
}

func (s *BinaryDecoderSuite) subtestDecode(testCase binDecTestCase) {
	s.Run(testCase.Name, func() {
		require := s.Require()
		actual := kv1.NewKeyValueEmpty()
		dec := kv1.NewBinaryDecoder(bytes.NewReader(testCase.Data))
		err := dec.Decode(actual)

		if testCase.Err == "" {
			require.NoError(err)
			s.RequireEqualKeyValue(testCase.Expected, actual)
		} else {
			require.ErrorContains(err, testCase.Err)
		}
	})
}
