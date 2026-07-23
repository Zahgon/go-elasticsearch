package esapi

import (
	"context"
	"net/http"
	"time"
)

func newIndicesCreateDataStreamFunc(t Transport) IndicesCreateDataStream {
	_ = "STUB: not implemented"
	return *new(IndicesCreateDataStream)
}

type IndicesCreateDataStream func(name string, o ...func(*IndicesCreateDataStreamRequest)) (*Response, error)

type IndicesCreateDataStreamRequest struct {
	Name string

	MasterTimeout time.Duration
	Timeout       time.Duration

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r IndicesCreateDataStreamRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f IndicesCreateDataStream) WithContext(v context.Context) func(*IndicesCreateDataStreamRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesCreateDataStream) WithMasterTimeout(v time.Duration) func(*IndicesCreateDataStreamRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesCreateDataStream) WithTimeout(v time.Duration) func(*IndicesCreateDataStreamRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesCreateDataStream) WithPretty() func(*IndicesCreateDataStreamRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesCreateDataStream) WithHuman() func(*IndicesCreateDataStreamRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesCreateDataStream) WithErrorTrace() func(*IndicesCreateDataStreamRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesCreateDataStream) WithFilterPath(v ...string) func(*IndicesCreateDataStreamRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesCreateDataStream) WithHeader(h map[string]string) func(*IndicesCreateDataStreamRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesCreateDataStream) WithOpaqueID(s string) func(*IndicesCreateDataStreamRequest) {
	_ = "STUB: not implemented"
	return nil
}
