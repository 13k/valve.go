package kv1

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

const (
	textIndent      = "  "
	textObjectStart = "{"
	textObjectEnd   = "}"
)

// TextEncoder writes text-encoded KeyValue nodes to an output stream.
type TextEncoder struct {
	w *bufio.Writer
}

// NewTextEncoder returns a new text encoder that writes to w.
//
// It wraps `w` with a `bufio.Writer`.
func NewTextEncoder(w io.Writer) *TextEncoder {
	return &TextEncoder{w: bufio.NewWriter(w)}
}

// Encode writes the KeyValue text encoding of kv to the stream.
func (e *TextEncoder) Encode(kv *KeyValue) error {
	return e.encode(kv, 0)
}

func (e *TextEncoder) encode(kv *KeyValue, level int) error {
	ty := kv.Type()

	if ty == TypeEnd || ty == TypeInvalid {
		return textEncUnsupportedError(kv)
	}

	indent := strings.Repeat(textIndent, level)
	qKey := strconv.Quote(kv.Key())

	if _, err := fmt.Fprintf(e.w, "%s%s ", indent, qKey); err != nil {
		return textEncError(writeError(err), kv)
	}

	var err error

	switch ty {
	case TypeObject:
		{
			err = e.encodeObject(kv, indent, level)
		}
	case
		TypeString, TypeWString,
		TypeInt32, TypeColor, TypePointer,
		TypeInt64,
		TypeUint64,
		TypeFloat32:
		{
			err = e.encodeValue(kv)
		}
	case TypeInvalid, TypeEnd:
		panic("unreachable")
	}

	if err != nil {
		return err
	}

	return textEncError(e.flush(), kv)
}

func (e *TextEncoder) flush() error {
	return writeError(e.w.Flush())
}

func (e *TextEncoder) writeByte(b byte) error {
	return writeError(e.w.WriteByte(b))
}

func (e *TextEncoder) writeString(s string) error {
	_, err := e.w.WriteString(s)

	return writeError(err)
}

func (e *TextEncoder) encodeObject(kv *KeyValue, indent string, level int) error {
	if err := e.writeString(textObjectStart); err != nil {
		return textEncError(err, kv)
	}

	if err := e.writeByte('\n'); err != nil {
		return textEncError(err, kv)
	}

	nextLevel := level + 1

	for _, c := range kv.Children() {
		if err := e.encode(c, nextLevel); err != nil {
			return textEncError(err, kv)
		}

		if err := e.writeByte('\n'); err != nil {
			return textEncError(err, kv)
		}
	}

	if err := e.writeString(indent); err != nil {
		return textEncError(err, kv)
	}

	if err := e.writeString(textObjectEnd); err != nil {
		return textEncError(err, kv)
	}

	if err := e.writeByte('\n'); err != nil {
		return textEncError(err, kv)
	}

	return nil
}

func (e *TextEncoder) encodeValue(kv *KeyValue) error {
	s, err := kv.ToString()

	if err != nil {
		return textEncError(err, kv)
	}

	s = strconv.Quote(s)
	err = e.writeString(s)

	return textEncError(err, kv)
}
