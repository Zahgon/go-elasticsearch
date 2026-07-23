package esapi

import (
	"context"
	"net/http"
	"time"
)

func newEnrichExecutePolicyFunc(t Transport) EnrichExecutePolicy {
	_ = "STUB: not implemented"
	return *new(EnrichExecutePolicy)
}

type EnrichExecutePolicy func(name string, o ...func(*EnrichExecutePolicyRequest)) (*Response, error)

type EnrichExecutePolicyRequest struct {
	Name string

	MasterTimeout     time.Duration
	WaitForCompletion *bool

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r EnrichExecutePolicyRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f EnrichExecutePolicy) WithContext(v context.Context) func(*EnrichExecutePolicyRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f EnrichExecutePolicy) WithMasterTimeout(v time.Duration) func(*EnrichExecutePolicyRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f EnrichExecutePolicy) WithWaitForCompletion(v bool) func(*EnrichExecutePolicyRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f EnrichExecutePolicy) WithPretty() func(*EnrichExecutePolicyRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f EnrichExecutePolicy) WithHuman() func(*EnrichExecutePolicyRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f EnrichExecutePolicy) WithErrorTrace() func(*EnrichExecutePolicyRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f EnrichExecutePolicy) WithFilterPath(v ...string) func(*EnrichExecutePolicyRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f EnrichExecutePolicy) WithHeader(h map[string]string) func(*EnrichExecutePolicyRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f EnrichExecutePolicy) WithOpaqueID(s string) func(*EnrichExecutePolicyRequest) {
	_ = "STUB: not implemented"
	return nil
}
