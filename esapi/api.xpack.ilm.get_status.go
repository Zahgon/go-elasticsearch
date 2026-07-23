package esapi

import (
	"context"
	"net/http"
)

func newILMGetStatusFunc(t Transport) ILMGetStatus {
	_ = "STUB: not implemented"
	return *new(ILMGetStatus)
}

type ILMGetStatus func(o ...func(*ILMGetStatusRequest)) (*Response, error)

type ILMGetStatusRequest struct {
	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r ILMGetStatusRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f ILMGetStatus) WithContext(v context.Context) func(*ILMGetStatusRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ILMGetStatus) WithPretty() func(*ILMGetStatusRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ILMGetStatus) WithHuman() func(*ILMGetStatusRequest) { _ = "STUB: not implemented"; return nil }

func (f ILMGetStatus) WithErrorTrace() func(*ILMGetStatusRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ILMGetStatus) WithFilterPath(v ...string) func(*ILMGetStatusRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ILMGetStatus) WithHeader(h map[string]string) func(*ILMGetStatusRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ILMGetStatus) WithOpaqueID(s string) func(*ILMGetStatusRequest) {
	_ = "STUB: not implemented"
	return nil
}
