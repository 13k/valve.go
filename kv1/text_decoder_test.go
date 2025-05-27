package kv1_test

import (
	"bytes"
	"io"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/13k/valve.go/kv1"
)

func TestTextDecoder(t *testing.T) {
	t.Parallel()

	suite.Run(t, new(TextDecoderSuite))
}

type TextDecoderSuite struct {
	Suite
}

type textDecTestCase struct {
	Name            string
	Data            []byte
	Input           io.Reader
	Err             string
	Expected        *kv1.KeyValue
	ExpectedPartial []textDecPartialCase
}

type textDecPartialCase struct {
	Path     []string
	Type     kv1.Type
	Value    string
	Children int
}

func (s *TextDecoderSuite) TestDecode() {
	testCases := []textDecTestCase{
		{
			Name: "input_nil",
			Err:  `1:1: unexpected EOF`,
		},
		{
			Name: "input_empty",
			Data: []byte{},
			Err:  `1:1: unexpected EOF`,
		},
		{
			Name:  "incomplete",
			Input: s.MustOpenFixture("sample.invalid-incomplete.txt"),
			Err:   `7:1: unexpected EOF`,
		},
		{
			Name:  "key_missing",
			Input: s.MustOpenFixture("sample.invalid-missing_key.txt"),
			Err:   `3:4: unexpected token "{"`,
		},
		{
			Name:  "value_missing",
			Input: s.MustOpenFixture("sample.invalid-missing_value.txt"),
			Err:   `7:4: unexpected token "}"`,
		},
		{
			Name:  "root_non_object",
			Input: s.MustOpenFixture("sample.invalid-non_object_root.txt"),
			Err:   `2:1: unexpected EOF`,
		},
		{
			Name:  "valid",
			Input: s.MustOpenFixture("sample.valid.txt"),
			Expected: kv1.NewKeyValueRoot("root").
				AddString("key", "value").
				AddString("qkey1", "value").
				AddString("qkey2", "qvalue").
				AddString("k{ey", "v}alue").
				AddString("esc_quote", `hello "world"`).
				AddString("esc.newline", "hello\nworld").
				AddString("esc,tab", "hello\tworld").
				AddString("esc*backslash", "hello\\world").
				AddString("esc", "hell\to\\\"wo\nrld\\x01\"").
				AddString("int", "13").
				AddString("negint", "-13").
				AddString("float", "1.3").
				AddChild(
					kv1.NewKeyValueObject("seq", nil).
						AddString("0", "don't!").
						AddString("1", "_second").
						AddString("2", "3"),
				).
				AddChild(
					kv1.NewKeyValueObject("nonseq", nil).
						AddString("1a", "one").
						AddString("2b", "two").
						AddString("3c", "three"),
				),
		},
		// no multi-line support yet
		{
			Name:  "value_multiline",
			Input: s.MustOpenFixture("addon_english.txt"),
			Err:   `12:95: literal not terminated`,
		},
		{
			Name:  "addoninfo",
			Input: s.MustOpenFixture("addoninfo.txt"),
			Expected: kv1.NewKeyValueRoot("AddonInfo").
				AddChild(
					kv1.NewKeyValueObject("siege02", nil).
						AddString("MaxPlayers", "5"),
				).
				AddChild(
					kv1.NewKeyValueObject("hero_picker", nil).
						AddString("background_map", "scenes/darkmoon_hero_pick"),
				),
		},
		{
			Name:  "gameinfo",
			Input: s.MustOpenFixture("gameinfo.gi"),
			ExpectedPartial: []textDecPartialCase{
				{
					Path:     nil,
					Type:     kv1.TypeObject,
					Children: 23,
				},
				{
					Path:     []string{"FileSystem"},
					Type:     kv1.TypeObject,
					Children: 5,
				},
				{
					Path:     []string{"MaterialSystem2"},
					Type:     kv1.TypeObject,
					Children: 1,
				},
				{
					Path:     []string{"Engine2"},
					Type:     kv1.TypeObject,
					Children: 21,
				},
				{
					Path:     []string{"SceneFileCache"},
					Type:     kv1.TypeObject,
					Children: 1,
				},
				{
					Path:     []string{"SceneSystem"},
					Type:     kv1.TypeObject,
					Children: 6,
				},
				{
					Path:     []string{"SoundSystem"},
					Type:     kv1.TypeObject,
					Children: 2,
				},
				{
					Path:     []string{"ToolsEnvironment"},
					Type:     kv1.TypeObject,
					Children: 4,
				},
				{
					Path:     []string{"Hammer"},
					Type:     kv1.TypeObject,
					Children: 19,
				},
				{
					Path:  []string{"Hammer", "DefaultTextureScale"},
					Type:  kv1.TypeString,
					Value: "0.250000",
				},
				{
					Path:  []string{"Hammer", "TileGridBlendDefaultColor"},
					Type:  kv1.TypeString,
					Value: "0 255 0",
				},
				{
					Path:     []string{"MaterialEditor"},
					Type:     kv1.TypeObject,
					Children: 2,
				},
				{
					Path:     []string{"ResourceCompiler"},
					Type:     kv1.TypeObject,
					Children: 2,
				},
				{
					Path:  []string{"ResourceCompiler", "DefaultMapBuilders", "gridnav"},
					Type:  kv1.TypeString,
					Value: "1",
				},
				{
					Path:     []string{"RenderPipelineAliases"},
					Type:     kv1.TypeObject,
					Children: 2,
				},
				{
					Path:     []string{"RenderSystem"},
					Type:     kv1.TypeObject,
					Children: 1,
				},
			},
		},
	}

	for _, testCase := range testCases {
		s.subtestDecode(testCase)
	}
}

func (s *TextDecoderSuite) subtestDecode(testCase textDecTestCase) {
	s.Run(testCase.Name, func() {
		require := s.Require()
		input := testCase.Input

		if input == nil {
			input = bytes.NewReader(testCase.Data)
		}

		if closer, ok := input.(io.Closer); ok {
			defer closer.Close()
		}

		dec := kv1.NewTextDecoder(input)
		actual := kv1.NewKeyValueEmpty()
		err := dec.Decode(actual)

		if testCase.Err == "" {
			require.NoError(err)

			if len(testCase.ExpectedPartial) == 0 {
				s.RequireEqualKeyValue(testCase.Expected, actual)
			}
		} else {
			require.ErrorContains(err, testCase.Err)
		}

		for partialCaseIdx, expectedChild := range testCase.ExpectedPartial {
			child := actual

			for _, key := range expectedChild.Path {
				child = child.FindChild(key)

				require.NotNilf(
					child,
					"child index=%d path=%q",
					partialCaseIdx,
					expectedChild.Path,
				)
			}

			require.Equalf(
				expectedChild.Type,
				child.Type(),
				"child index=%d path=%q",
				partialCaseIdx,
				expectedChild.Path,
			)

			if expectedChild.Type == kv1.TypeObject {
				require.Lenf(
					child.Children(),
					expectedChild.Children,
					"child index=%d path=%q",
					partialCaseIdx,
					expectedChild.Path,
				)
			} else {
				require.Equalf(
					expectedChild.Value,
					child.Value(),
					"child index=%d path=%q",
					partialCaseIdx,
					expectedChild.Path,
				)
			}
		}
	})
}
