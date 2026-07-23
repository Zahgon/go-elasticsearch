package esapi

import (
	"context"
	"net/http"
	"time"
)

func newILMStopFunc(t Transport) ILMStop { _ = "STUB: not implemented"; return *new(ILMStop) }

type ILMStop func(o ...func(*ILMStopRequest)) (*Response, error)

type ILMStopRequest struct {
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

func (r ILMStopRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f ILMStop) WithContext(v context.Context) func(*ILMStopRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ILMStop) WithMasterTimeout(v time.Duration) func(*ILMStopRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ILMStop) WithTimeout(v time.Duration) func(*ILMStopRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ILMStop) WithPretty() func(*ILMStopRequest) { _ = "STUB: not implemented"; return nil }

func (f ILMStop) WithHuman() func(*ILMStopRequest) { _ = "STUB: not implemented"; return nil }

func (f ILMStop) WithErrorTrace() func(*ILMStopRequest) { _ = "STUB: not implemented"; return nil }

func (f ILMStop) WithFilterPath(v ...string) func(*ILMStopRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ILMStop) WithHeader(h map[string]string) func(*ILMStopRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ILMStop) WithOpaqueID(s string) func(*ILMStopRequest) {
	_ = "STUB: not implemented"
	return nil
}
