package parser

import (
	"errors"
	"fmt"
	"text/scanner"
)

var (
	ErrEOF = errors.New("unexpected EOF")
)

var _ error = (*ParseError)(nil)

type ParseError struct {
	pos scanner.Position
	err error
}

func parseError(pos scanner.Position, err error) error {
	if err == nil {
		return nil
	}

	return &ParseError{
		pos: pos,
		err: err,
	}
}

func (err *ParseError) Error() string {
	return fmt.Sprintf("parse error at %s: %s", err.pos, err.err)
}

func (err *ParseError) Unwrap() error {
	return err.err
}

var _ error = (*ScanError)(nil)

type ScanError struct {
	pos scanner.Position
	msg string
}

func scanError(pos scanner.Position, msg string) error {
	return &ScanError{
		pos: pos,
		msg: msg,
	}
}

func scanErrorf(pos scanner.Position, format string, values ...any) error {
	msg := fmt.Sprintf(format, values...)

	return scanError(pos, msg)
}

func tokenError(pos scanner.Position, token rune) error {
	return scanErrorf(pos, "unexpected token %s", scanner.TokenString(token))
}

func (err *ScanError) Error() string {
	return fmt.Sprintf("scan error at %s: %s", err.pos, err.msg)
}
