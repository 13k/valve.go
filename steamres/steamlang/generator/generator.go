package generator

import (
	"io"

	"github.com/13k/valve.go/steamres/steamlang/parser"
)

// Generator is the interface for code generators
type Generator interface {
	Generate(w io.Writer, node *parser.Node) error
}
