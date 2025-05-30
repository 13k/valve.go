package kv1_test

import (
	"bytes"
	"io"

	"github.com/13k/valve.go/kv1"
)

type kvTestCaseMap struct {
	Name      string
	Subject   *kv1.KeyValue
	TextData  []byte
	TextInput io.Reader
	Expected  map[string]any
	Err       string
}

func (s *KeyValueSuite) TestMap() {
	testCases := []kvTestCaseMap{
		{
			Name:    "ty_invalid",
			Subject: kv1.NewKeyValueEmpty(),
			Err:     `Invalid to map`,
		},
		{
			Name:    "ty_end",
			Subject: kv1.NewKeyValueEnd(),
			Err:     `End to map`,
		},
		{
			Name:    "non_object",
			Subject: kv1.NewKeyValueString("foo", "bar", nil),
			Err:     `String to map`,
		},
		{
			Name: "object_ok",
			Subject: kv1.
				NewKeyValueObjectRoot("root").
				AddString("str", "foo1").
				AddWString("wstr", "bar1").
				AddInt32("int32", 0).
				AddColor("color", 1).
				AddPointer("pointer", 2).
				AddInt64("int64", 3).
				AddUint64("uint64", 4).
				AddFloat32("float32", 5).
				AddChild(
					kv1.NewKeyValueObjectRoot("object").
						AddString("str", "foo2").
						AddWString("wstr", "bar2").
						AddInt32("int32", 10).
						AddColor("color", 11).
						AddPointer("pointer", 12).
						AddInt64("int64", 13).
						AddUint64("uint64", 14).
						AddFloat32("float32", 15),
				),
			Expected: map[string]any{
				"root": map[string]any{
					"str":     "foo1",
					"wstr":    "bar1",
					"int32":   int32(0),
					"color":   int32(1),
					"pointer": int32(2),
					"int64":   int64(3),
					"uint64":  uint64(4),
					"float32": float32(5),
					"object": map[string]any{
						"str":     "foo2",
						"wstr":    "bar2",
						"int32":   int32(10),
						"color":   int32(11),
						"pointer": int32(12),
						"int64":   int64(13),
						"uint64":  uint64(14),
						"float32": float32(15),
					},
				},
			},
		},
		{
			Name:      "sample.valid.txt",
			TextInput: s.MustOpenFixture("sample.valid.txt"),
			Expected: map[string]any{
				"root": map[string]any{
					"key":           "value",
					"qkey1":         "value",
					"qkey2":         "qvalue",
					"k{ey":          "v}alue",
					"esc_quote":     `hello "world"`,
					"esc.newline":   "hello\nworld",
					"esc,tab":       "hello\tworld",
					"esc*backslash": "hello\\world",
					"esc":           "hell\to\\\"wo\nrld\\x01\"",
					"int":           "13",
					"negint":        "-13",
					"float":         "1.3",
					"seq": map[string]any{
						"0": "don't!",
						"1": "_second",
						"2": "3",
					},
					"nonseq": map[string]any{
						"1a": "one",
						"2b": "two",
						"3c": "three",
					},
				},
			},
		},
		{
			Name:      "addoninfo.txt",
			TextInput: s.MustOpenFixture("addoninfo.txt"),
			Expected: map[string]any{
				"AddonInfo": map[string]any{
					"siege02": map[string]any{
						"MaxPlayers": "5",
					},
					"hero_picker": map[string]any{
						"background_map": "scenes/darkmoon_hero_pick",
					},
				},
			},
		},
		{
			Name:     "textData1",
			TextData: textData1,
			Expected: map[string]any{
				"RP": map[string]any{
					"status":          `#DOTA_RP_LEAGUE_MATCH_PLAYING_AS`,
					"steam_display":   `#DOTA_RP_LEAGUE_MATCH_PLAYING_AS`,
					"num_params":      `3`,
					"lobby":           `lobby_id: 26628083760328052 lobby_state: RUN password: true game_mode: DOTA_GAMEMODE_CM member_count: 10 max_member_count: 10 name: "Team Secret vs Vikin.GG " lobby_type: 1`,
					"party":           `party_state: IN_MATCH`,
					"WatchableGameID": `26628083760328052`,
					"param0":          `#DOTA_lobby_type_name_lobby`,
					"param1":          `8`,
					"param2":          `#npc_dota_hero_grimstroke`,
				},
			},
		},
	}

	for _, testCase := range testCases {
		s.testMap(testCase)
	}
}

func (s *KeyValueSuite) testMap(testCase kvTestCaseMap) {
	s.Run(testCase.Name, func() {
		var subject *kv1.KeyValue

		require := s.Require()

		switch {
		case testCase.Subject != nil:
			{
				subject = testCase.Subject
			}
		case testCase.TextData != nil:
			{
				input := bytes.NewReader(testCase.TextData)
				dec := kv1.NewTextDecoder(input)
				subject = kv1.NewKeyValueEmpty()

				require.NoError(dec.Decode(subject))
			}
		case testCase.TextInput != nil:
			{
				input := testCase.TextInput

				if closer, ok := input.(io.Closer); ok {
					defer closer.Close()
				}

				dec := kv1.NewTextDecoder(input)
				subject = kv1.NewKeyValueEmpty()

				require.NoError(dec.Decode(subject))
			}
		default:
			{
				s.FailNow("invalid test case")
			}
		}

		actual, err := subject.Map()

		if testCase.Err == "" {
			require.NoError(err)
			require.Equal(testCase.Expected, actual)
		} else {
			require.ErrorContains(err, testCase.Err)
		}
	})
}
