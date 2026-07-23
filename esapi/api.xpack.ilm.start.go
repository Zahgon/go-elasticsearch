package esapi

import (
	"context"
	"net/http"
	"time"
)

func newILMStartFunc(t Transport) ILMStart { _ = "STUB: not implemented"; return *new(ILMStart) }

type ILMStart func(o ...func(*ILMStartRequest)) (*Response, error)

type ILMStartRequest struct {
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

func (r ILMStartRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f ILMStart) WithContext(v context.Context) func(*ILMStartRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ILMStart) WithMasterTimeout(v time.Duration) func(*ILMStartRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ILMStart) WithTimeout(v time.Duration) func(*ILMStartRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ILMStart) WithPretty() func(*ILMStartRequest) { _ = "STUB: not implemented"; return nil }

func (f ILMStart) WithHuman() func(*ILMStartRequest) { _ = "STUB: not implemented"; return nil }

func (f ILMStart) WithErrorTrace() func(*ILMStartRequest) { _ = "STUB: not implemented"; return nil }

func (f ILMStart) WithFilterPath(v ...string) func(*ILMStartRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ILMStart) WithHeader(h map[string]string) func(*ILMStartRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ILMStart) WithOpaqueID(s string) func(*ILMStartRequest) {
	_ = "STUB: not implemented"
	return nil
}
