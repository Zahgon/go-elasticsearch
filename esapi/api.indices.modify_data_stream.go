package esapi

import (
	"context"
	"io"
	"net/http"
)

func newIndicesModifyDataStreamFunc(t Transport) IndicesModifyDataStream {
	_ = "STUB: not implemented"
	return *new(IndicesModifyDataStream)
}

type IndicesModifyDataStream func(body io.Reader, o ...func(*IndicesModifyDataStreamRequest)) (*Response, error)

type IndicesModifyDataStreamRequest struct {
	Body io.Reader

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r IndicesModifyDataStreamRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f IndicesModifyDataStream) WithContext(v context.Context) func(*IndicesModifyDataStreamRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesModifyDataStream) WithPretty() func(*IndicesModifyDataStreamRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesModifyDataStream) WithHuman() func(*IndicesModifyDataStreamRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesModifyDataStream) WithErrorTrace() func(*IndicesModifyDataStreamRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesModifyDataStream) WithFilterPath(v ...string) func(*IndicesModifyDataStreamRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesModifyDataStream) WithHeader(h map[string]string) func(*IndicesModifyDataStreamRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesModifyDataStream) WithOpaqueID(s string) func(*IndicesModifyDataStreamRequest) {
	_ = "STUB: not implemented"
	return nil
}
