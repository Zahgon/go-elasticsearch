package esapi

import (
	"context"
	"net/http"
)

func newMLInfoFunc(t Transport) MLInfo { _ = "STUB: not implemented"; return *new(MLInfo) }

type MLInfo func(o ...func(*MLInfoRequest)) (*Response, error)

type MLInfoRequest struct {
	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r MLInfoRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f MLInfo) WithContext(v context.Context) func(*MLInfoRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLInfo) WithPretty() func(*MLInfoRequest) { _ = "STUB: not implemented"; return nil }

func (f MLInfo) WithHuman() func(*MLInfoRequest) { _ = "STUB: not implemented"; return nil }

func (f MLInfo) WithErrorTrace() func(*MLInfoRequest) { _ = "STUB: not implemented"; return nil }

func (f MLInfo) WithFilterPath(v ...string) func(*MLInfoRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLInfo) WithHeader(h map[string]string) func(*MLInfoRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLInfo) WithOpaqueID(s string) func(*MLInfoRequest) { _ = "STUB: not implemented"; return nil }
