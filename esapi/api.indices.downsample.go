package esapi

import (
	"context"
	"io"
	"net/http"
)

func newIndicesDownsampleFunc(t Transport) IndicesDownsample {
	_ = "STUB: not implemented"
	return *new(IndicesDownsample)
}

type IndicesDownsample func(index string, body io.Reader, target_index string, o ...func(*IndicesDownsampleRequest)) (*Response, error)

type IndicesDownsampleRequest struct {
	Index string

	Body io.Reader

	TargetIndex string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r IndicesDownsampleRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f IndicesDownsample) WithContext(v context.Context) func(*IndicesDownsampleRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesDownsample) WithPretty() func(*IndicesDownsampleRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesDownsample) WithHuman() func(*IndicesDownsampleRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesDownsample) WithErrorTrace() func(*IndicesDownsampleRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesDownsample) WithFilterPath(v ...string) func(*IndicesDownsampleRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesDownsample) WithHeader(h map[string]string) func(*IndicesDownsampleRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesDownsample) WithOpaqueID(s string) func(*IndicesDownsampleRequest) {
	_ = "STUB: not implemented"
	return nil
}
