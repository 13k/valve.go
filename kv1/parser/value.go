package parser

import "strings"

//nolint:gochecknoglobals
var escapeSeqs = [][2]string{
	{`\\`, `\`},
	{`\n`, "\n"},
	{`\t`, "\t"},
	{`\"`, `"`},
}

func unquoteString(s string) string {
	q := tokenQuote.String()
	// `strings.Trim` eats too much in the case "hello \"world\""
	// `strconv.Unquote` errors with non-quoted strings
	s = strings.TrimPrefix(s, q)
	s = strings.TrimSuffix(s, q)

	return s
}

func unescapeString(s string) string {
	for _, esc := range escapeSeqs {
		s = strings.ReplaceAll(s, esc[0], esc[1])
	}

	return s
}

func parseString(s string) string {
	s = unquoteString(s)
	s = unescapeString(s)

	return s
}
