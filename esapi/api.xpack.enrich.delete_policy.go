package esapi

import (
	"context"
	"net/http"
	"time"
)

func newEnrichDeletePolicyFunc(t Transport) EnrichDeletePolicy {
	_ = "STUB: not implemented"
	return *new(EnrichDeletePolicy)
}

type EnrichDeletePolicy func(name string, o ...func(*EnrichDeletePolicyRequest)) (*Response, error)

type EnrichDeletePolicyRequest struct {
	Name string

	MasterTimeout time.Duration

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r EnrichDeletePolicyRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f EnrichDeletePolicy) WithContext(v context.Context) func(*EnrichDeletePolicyRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f EnrichDeletePolicy) WithMasterTimeout(v time.Duration) func(*EnrichDeletePolicyRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f EnrichDeletePolicy) WithPretty() func(*EnrichDeletePolicyRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f EnrichDeletePolicy) WithHuman() func(*EnrichDeletePolicyRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f EnrichDeletePolicy) WithErrorTrace() func(*EnrichDeletePolicyRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f EnrichDeletePolicy) WithFilterPath(v ...string) func(*EnrichDeletePolicyRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f EnrichDeletePolicy) WithHeader(h map[string]string) func(*EnrichDeletePolicyRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f EnrichDeletePolicy) WithOpaqueID(s string) func(*EnrichDeletePolicyRequest) {
	_ = "STUB: not implemented"
	return nil
}
