package esutil

import (
	"bytes"
	"io"
)

func NewJSONReader(v interface{}) io.Reader { _ = "STUB: not implemented"; return *new(io.Reader) }

type JSONEncoder interface {
	EncodeJSON(io.Writer) error
}

type JSONReader struct {
	val interface{}
	buf *bytes.Reader
}

func (r *JSONReader) ensureEncoded() error { _ = "STUB: not implemented"; return nil }

func (r *JSONReader) Read(p []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (r *JSONReader) Seek(offset int64, whence int) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (r *JSONReader) WriteTo(w io.Writer) (int64, error) { _ = "STUB: not implemented"; return 0, nil }

func (r *JSONReader) encode(w io.Writer) error { _ = "STUB: not implemented"; return nil }

type countingWriter struct {
	io.Writer
	n int
}

func (cw *countingWriter) Write(p []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }
