package esapi

import (
	"context"
	"io"
	"net/http"
)

func newMLUpdateFilterFunc(t Transport) MLUpdateFilter {
	_ = "STUB: not implemented"
	return *new(MLUpdateFilter)
}

type MLUpdateFilter func(body io.Reader, filter_id string, o ...func(*MLUpdateFilterRequest)) (*Response, error)

type MLUpdateFilterRequest struct {
	Body io.Reader

	FilterID string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r MLUpdateFilterRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f MLUpdateFilter) WithContext(v context.Context) func(*MLUpdateFilterRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLUpdateFilter) WithPretty() func(*MLUpdateFilterRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLUpdateFilter) WithHuman() func(*MLUpdateFilterRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLUpdateFilter) WithErrorTrace() func(*MLUpdateFilterRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLUpdateFilter) WithFilterPath(v ...string) func(*MLUpdateFilterRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLUpdateFilter) WithHeader(h map[string]string) func(*MLUpdateFilterRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLUpdateFilter) WithOpaqueID(s string) func(*MLUpdateFilterRequest) {
	_ = "STUB: not implemented"
	return nil
}
