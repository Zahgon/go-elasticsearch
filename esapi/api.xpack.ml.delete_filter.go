package esapi

import (
	"context"
	"net/http"
)

func newMLDeleteFilterFunc(t Transport) MLDeleteFilter {
	_ = "STUB: not implemented"
	return *new(MLDeleteFilter)
}

type MLDeleteFilter func(filter_id string, o ...func(*MLDeleteFilterRequest)) (*Response, error)

type MLDeleteFilterRequest struct {
	FilterID string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r MLDeleteFilterRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f MLDeleteFilter) WithContext(v context.Context) func(*MLDeleteFilterRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLDeleteFilter) WithPretty() func(*MLDeleteFilterRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLDeleteFilter) WithHuman() func(*MLDeleteFilterRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLDeleteFilter) WithErrorTrace() func(*MLDeleteFilterRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLDeleteFilter) WithFilterPath(v ...string) func(*MLDeleteFilterRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLDeleteFilter) WithHeader(h map[string]string) func(*MLDeleteFilterRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLDeleteFilter) WithOpaqueID(s string) func(*MLDeleteFilterRequest) {
	_ = "STUB: not implemented"
	return nil
}
