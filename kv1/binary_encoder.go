package kv1

import (
	"bufio"
	"encoding/binary"
	"io"
)

// BinaryEncoder writes binary-encoded KeyValue nodes to an output stream.
type BinaryEncoder struct {
	w *bufio.Writer
}

// NewBinaryEncoder returns a new binary encoder that writes to w.
//
// It wraps `w` with a `bufio.Writer`.
func NewBinaryEncoder(w io.Writer) *BinaryEncoder {
	return &BinaryEncoder{w: bufio.NewWriter(w)}
}

// Encode writes the KeyValue binary encoding of kv to the stream.
func (e *BinaryEncoder) Encode(kv *KeyValue) error { //nolint:cyclop
	ty := kv.Type()

	if ty == TypeInvalid || ty == TypeEnd || ty == TypeWString {
		return binEncUnsupportedError(kv)
	}

	if err := e.writeType(ty); err != nil {
		return binEncError(err, kv)
	}

	if err := e.writeString(kv.Key()); err != nil {
		return binEncError(err, kv)
	}

	var err error

	switch ty {
	case TypeObject:
		err = e.encodeObject(kv)
	case TypeString:
		err = e.encodeString(kv)
	case TypeInt32:
		err = e.encodeInt32(kv)
	case TypeColor:
		err = e.encodeColor(kv)
	case TypePointer:
		err = e.encodePointer(kv)
	case TypeInt64:
		err = e.encodeInt64(kv)
	case TypeUint64:
		err = e.encodeUint64(kv)
	case TypeFloat32:
		err = e.encodeFloat32(kv)
	case TypeEnd, TypeInvalid, TypeWString:
		panic("unreachable")
	}

	if err != nil {
		return err
	}

	return binEncError(e.flush(), kv)
}

func (e *BinaryEncoder) flush() error {
	return writeError(e.w.Flush())
}

func (e *BinaryEncoder) writeType(t Type) error {
	return writeError(e.w.WriteByte(t.Byte()))
}

func (e *BinaryEncoder) writeString(s string) error {
	if _, err := e.w.WriteString(s); err != nil {
		return writeError(err)
	}

	if err := e.w.WriteByte(binDecStringDelim); err != nil {
		return writeError(err)
	}

	return nil
}

func (e *BinaryEncoder) writeInt32(n int32) error {
	return writeError(binary.Write(e.w, binary.LittleEndian, n))
}

func (e *BinaryEncoder) writeInt64(n int64) error {
	return writeError(binary.Write(e.w, binary.LittleEndian, n))
}

func (e *BinaryEncoder) writeUint64(n uint64) error {
	return writeError(binary.Write(e.w, binary.LittleEndian, n))
}

func (e *BinaryEncoder) writeFloat32(n float32) error {
	return writeError(binary.Write(e.w, binary.LittleEndian, n))
}

func (e *BinaryEncoder) encodeObject(kv *KeyValue) error {
	for _, c := range kv.Children() {
		if err := e.Encode(c); err != nil {
			return binEncError(err, kv)
		}
	}

	if err := e.writeType(TypeEnd); err != nil {
		return binEncError(err, kv)
	}

	return nil
}

func (e *BinaryEncoder) encodeString(kv *KeyValue) error {
	s, err := kv.String()

	if err == nil {
		err = e.writeString(s)
	}

	return binEncError(err, kv)
}

func (e *BinaryEncoder) encodeInt32(kv *KeyValue) error {
	n, err := kv.Int32()

	if err == nil {
		err = e.writeInt32(n)
	}

	return binEncError(err, kv)
}

func (e *BinaryEncoder) encodeColor(kv *KeyValue) error {
	n, err := kv.Color()

	if err == nil {
		err = e.writeInt32(n)
	}

	return binEncError(err, kv)
}

func (e *BinaryEncoder) encodePointer(kv *KeyValue) error {
	n, err := kv.Pointer()

	if err == nil {
		err = e.writeInt32(n)
	}

	return binEncError(err, kv)
}

func (e *BinaryEncoder) encodeInt64(kv *KeyValue) error {
	n, err := kv.Int64()

	if err == nil {
		err = e.writeInt64(n)
	}

	return binEncError(err, kv)
}

func (e *BinaryEncoder) encodeUint64(kv *KeyValue) error {
	n, err := kv.Uint64()

	if err == nil {
		err = e.writeUint64(n)
	}

	return binEncError(err, kv)
}

func (e *BinaryEncoder) encodeFloat32(kv *KeyValue) error {
	n, err := kv.Float32()

	if err == nil {
		err = e.writeFloat32(n)
	}

	return binEncError(err, kv)
}
