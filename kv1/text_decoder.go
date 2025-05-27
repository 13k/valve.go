package kv1

import (
	"bufio"
	"io"

	"github.com/13k/valve.go/kv1/parser"
)

// TextDecoder reads and decodes text-encoded KeyValue nodes from an input stream.
type TextDecoder struct {
	p *parser.TextParser
}

// NewTextDecoder returns a new text decoder that reads from r.
//
// It wraps `r` with a `bufio.Reader`.
func NewTextDecoder(r io.Reader) *TextDecoder {
	p := parser.NewTextParser("", bufio.NewReader(r))

	return &TextDecoder{p: p}
}

// Decode reads the next text-encoded KeyValue node from its input and stores it in the value
// pointed to by kv.
//
// The parser makes no assumptions regarding field types, so all fields are of type TypeString.
func (d *TextDecoder) Decode(kv *KeyValue) error {
	root, err := d.p.Parse()

	if err != nil {
		return textDecError(err)
	}

	applyAST(kv, root)

	return nil
}

func applyAST(kv *KeyValue, node *parser.Node) {
	kv.SetKey(node.Key)

	switch node.Type {
	case parser.NodeTypeObject:
		kv.SetType(TypeObject)

		for _, nodeChild := range node.Children {
			applyAST(kv.NewChild(), nodeChild)
		}
	case parser.NodeTypeField:
		kv.SetString(node.Value)
	}
}
