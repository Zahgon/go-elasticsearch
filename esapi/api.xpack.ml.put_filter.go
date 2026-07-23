package esapi

import (
	"context"
	"io"
	"net/http"
)

func newMLPutFilterFunc(t Transport) MLPutFilter {
	_ = "STUB: not implemented"
	return *new(MLPutFilter)
}

type MLPutFilter func(body io.Reader, filter_id string, o ...func(*MLPutFilterRequest)) (*Response, error)

type MLPutFilterRequest struct {
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

func (r MLPutFilterRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f MLPutFilter) WithContext(v context.Context) func(*MLPutFilterRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLPutFilter) WithPretty() func(*MLPutFilterRequest) { _ = "STUB: not implemented"; return nil }

func (f MLPutFilter) WithHuman() func(*MLPutFilterRequest) { _ = "STUB: not implemented"; return nil }

func (f MLPutFilter) WithErrorTrace() func(*MLPutFilterRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLPutFilter) WithFilterPath(v ...string) func(*MLPutFilterRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLPutFilter) WithHeader(h map[string]string) func(*MLPutFilterRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLPutFilter) WithOpaqueID(s string) func(*MLPutFilterRequest) {
	_ = "STUB: not implemented"
	return nil
}
