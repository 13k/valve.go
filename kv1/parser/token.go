package parser

import "unicode"

type token rune

const (
	tokenSpace          token = ' '
	tokenTab            token = '\t'
	tokenCarriageReturn token = '\r'
	tokenNewline        token = '\n'
	tokenObjectStart    token = '{'
	tokenObjectEnd      token = '}'
	tokenQuote          token = '"'
	tokenComment        token = '/'
)

func (t token) Rune() rune {
	return rune(t)
}

func (t token) String() string {
	return string(t)
}

func isIdentRune(ch rune, _ int) bool {
	return unicode.In(ch, unicode.Number, unicode.Letter, unicode.Punct) &&
		ch != tokenObjectStart.Rune() &&
		ch != tokenObjectEnd.Rune() &&
		ch != tokenQuote.Rune() &&
		ch != tokenComment.Rune()
}
