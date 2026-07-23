package esapi

import (
	"context"
	"io"
	"net/http"
)

func newIndicesCreateFromFunc(t Transport) IndicesCreateFrom {
	_ = "STUB: not implemented"
	return *new(IndicesCreateFrom)
}

type IndicesCreateFrom func(dest string, source string, o ...func(*IndicesCreateFromRequest)) (*Response, error)

type IndicesCreateFromRequest struct {
	Body io.Reader

	Dest   string
	Source string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r IndicesCreateFromRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f IndicesCreateFrom) WithContext(v context.Context) func(*IndicesCreateFromRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesCreateFrom) WithBody(v io.Reader) func(*IndicesCreateFromRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesCreateFrom) WithPretty() func(*IndicesCreateFromRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesCreateFrom) WithHuman() func(*IndicesCreateFromRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesCreateFrom) WithErrorTrace() func(*IndicesCreateFromRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesCreateFrom) WithFilterPath(v ...string) func(*IndicesCreateFromRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesCreateFrom) WithHeader(h map[string]string) func(*IndicesCreateFromRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesCreateFrom) WithOpaqueID(s string) func(*IndicesCreateFromRequest) {
	_ = "STUB: not implemented"
	return nil
}
