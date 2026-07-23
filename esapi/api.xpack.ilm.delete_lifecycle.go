package esapi

import (
	"context"
	"net/http"
	"time"
)

func newILMDeleteLifecycleFunc(t Transport) ILMDeleteLifecycle {
	_ = "STUB: not implemented"
	return *new(ILMDeleteLifecycle)
}

type ILMDeleteLifecycle func(policy string, o ...func(*ILMDeleteLifecycleRequest)) (*Response, error)

type ILMDeleteLifecycleRequest struct {
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

func (r ILMDeleteLifecycleRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f ILMDeleteLifecycle) WithContext(v context.Context) func(*ILMDeleteLifecycleRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ILMDeleteLifecycle) WithMasterTimeout(v time.Duration) func(*ILMDeleteLifecycleRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ILMDeleteLifecycle) WithTimeout(v time.Duration) func(*ILMDeleteLifecycleRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ILMDeleteLifecycle) WithPretty() func(*ILMDeleteLifecycleRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ILMDeleteLifecycle) WithHuman() func(*ILMDeleteLifecycleRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ILMDeleteLifecycle) WithErrorTrace() func(*ILMDeleteLifecycleRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ILMDeleteLifecycle) WithFilterPath(v ...string) func(*ILMDeleteLifecycleRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ILMDeleteLifecycle) WithHeader(h map[string]string) func(*ILMDeleteLifecycleRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ILMDeleteLifecycle) WithOpaqueID(s string) func(*ILMDeleteLifecycleRequest) {
	_ = "STUB: not implemented"
	return nil
}
