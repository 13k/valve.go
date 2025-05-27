package parser

import (
	"io"
	"text/scanner"
)

// TextParser is a parser for KeyValue in text format.
type TextParser struct {
	s *scanner.Scanner
}

// NewTextParser creates a TextParser.
func NewTextParser(fname string, r io.Reader) *TextParser {
	s := new(scanner.Scanner).Init(r)

	s.Whitespace = 1<<tokenSpace | 1<<tokenTab | 1<<tokenCarriageReturn | 1<<tokenNewline

	s.Mode = scanner.ScanIdents |
		scanner.ScanStrings |
		scanner.ScanComments |
		scanner.SkipComments

	s.IsIdentRune = isIdentRune

	if fname == "" {
		if n, ok := r.(namer); ok {
			fname = n.Name()
		}
	}

	s.Filename = fname

	return &TextParser{s: s}
}

// Parse reads parses the text-encoded KeyValue values from the input stream, generating an AST
// tree.
func (p *TextParser) Parse() (*Node, error) { //nolint:cyclop,funlen
	var scanErr error

	root := NewNode()
	scope := root
	node := root

	p.s.Error = func(s *scanner.Scanner, msg string) {
		scanErr = scanError(s.Pos(), msg)
	}

	for {
		tok := p.s.Scan()

		if tok == scanner.EOF {
			if node.Type != NodeTypeObject && node.Key != "" && node.Value == "" {
				return root, parseError(p.s.Pos(), ErrEOF)
			}

			if scope != nil {
				return root, parseError(p.s.Pos(), ErrEOF)
			}

			break
		}

		if scanErr != nil {
			return root, scanErr
		}

		text := p.s.TokenText()

		switch {
		case node.Key == "":
			switch tok {
			case tokenObjectEnd.Rune():
				scope = scope.Parent
			case scanner.String, scanner.Ident:
				node.Key = parseString(text)
			default:
				return root, tokenError(p.s.Pos(), tok)
			}
		case node.Value == "":
			switch tok {
			case tokenObjectStart.Rune():
				node.Type = NodeTypeObject
			case scanner.String:
				node.Type = NodeTypeField
				node.Value = parseString(text)
			case scanner.Ident:
				node.Type = NodeTypeField
				node.Value = text
			default:
				return root, tokenError(p.s.Pos(), tok)
			}

			if node != scope {
				scope.addChild(node)
			}

			if node.Type == NodeTypeObject {
				scope = node
			}

			node = NewNode()
		}
	}

	return root, nil
}
