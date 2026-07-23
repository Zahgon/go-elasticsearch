package esapi

import (
	"context"
	"net/http"
	"time"
)

func newEnrichGetPolicyFunc(t Transport) EnrichGetPolicy {
	_ = "STUB: not implemented"
	return *new(EnrichGetPolicy)
}

type EnrichGetPolicy func(o ...func(*EnrichGetPolicyRequest)) (*Response, error)

type EnrichGetPolicyRequest struct {
	Name []string

	MasterTimeout time.Duration

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r EnrichGetPolicyRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f EnrichGetPolicy) WithContext(v context.Context) func(*EnrichGetPolicyRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f EnrichGetPolicy) WithName(v ...string) func(*EnrichGetPolicyRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f EnrichGetPolicy) WithMasterTimeout(v time.Duration) func(*EnrichGetPolicyRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f EnrichGetPolicy) WithPretty() func(*EnrichGetPolicyRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f EnrichGetPolicy) WithHuman() func(*EnrichGetPolicyRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f EnrichGetPolicy) WithErrorTrace() func(*EnrichGetPolicyRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f EnrichGetPolicy) WithFilterPath(v ...string) func(*EnrichGetPolicyRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f EnrichGetPolicy) WithHeader(h map[string]string) func(*EnrichGetPolicyRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f EnrichGetPolicy) WithOpaqueID(s string) func(*EnrichGetPolicyRequest) {
	_ = "STUB: not implemented"
	return nil
}
