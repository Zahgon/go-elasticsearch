package esapi

import (
	"context"
	"net/http"
	"time"
)

func newILMExplainLifecycleFunc(t Transport) ILMExplainLifecycle {
	_ = "STUB: not implemented"
	return *new(ILMExplainLifecycle)
}

type ILMExplainLifecycle func(index string, o ...func(*ILMExplainLifecycleRequest)) (*Response, error)

type ILMExplainLifecycleRequest struct {
	Index string

	MasterTimeout time.Duration
	OnlyErrors    *bool
	OnlyManaged   *bool

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r ILMExplainLifecycleRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f ILMExplainLifecycle) WithContext(v context.Context) func(*ILMExplainLifecycleRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ILMExplainLifecycle) WithMasterTimeout(v time.Duration) func(*ILMExplainLifecycleRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ILMExplainLifecycle) WithOnlyErrors(v bool) func(*ILMExplainLifecycleRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ILMExplainLifecycle) WithOnlyManaged(v bool) func(*ILMExplainLifecycleRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ILMExplainLifecycle) WithPretty() func(*ILMExplainLifecycleRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ILMExplainLifecycle) WithHuman() func(*ILMExplainLifecycleRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ILMExplainLifecycle) WithErrorTrace() func(*ILMExplainLifecycleRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ILMExplainLifecycle) WithFilterPath(v ...string) func(*ILMExplainLifecycleRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ILMExplainLifecycle) WithHeader(h map[string]string) func(*ILMExplainLifecycleRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ILMExplainLifecycle) WithOpaqueID(s string) func(*ILMExplainLifecycleRequest) {
	_ = "STUB: not implemented"
	return nil
}
