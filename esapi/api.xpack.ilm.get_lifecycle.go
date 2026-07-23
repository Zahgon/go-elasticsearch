package esapi

import (
	"context"
	"net/http"
	"time"
)

func newILMGetLifecycleFunc(t Transport) ILMGetLifecycle {
	_ = "STUB: not implemented"
	return *new(ILMGetLifecycle)
}

type ILMGetLifecycle func(o ...func(*ILMGetLifecycleRequest)) (*Response, error)

type ILMGetLifecycleRequest struct {
	Policy string

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

func (r ILMGetLifecycleRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f ILMGetLifecycle) WithContext(v context.Context) func(*ILMGetLifecycleRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ILMGetLifecycle) WithPolicy(v string) func(*ILMGetLifecycleRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ILMGetLifecycle) WithMasterTimeout(v time.Duration) func(*ILMGetLifecycleRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ILMGetLifecycle) WithTimeout(v time.Duration) func(*ILMGetLifecycleRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ILMGetLifecycle) WithPretty() func(*ILMGetLifecycleRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ILMGetLifecycle) WithHuman() func(*ILMGetLifecycleRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ILMGetLifecycle) WithErrorTrace() func(*ILMGetLifecycleRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ILMGetLifecycle) WithFilterPath(v ...string) func(*ILMGetLifecycleRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ILMGetLifecycle) WithHeader(h map[string]string) func(*ILMGetLifecycleRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ILMGetLifecycle) WithOpaqueID(s string) func(*ILMGetLifecycleRequest) {
	_ = "STUB: not implemented"
	return nil
}
