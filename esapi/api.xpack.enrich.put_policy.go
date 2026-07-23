package esapi

import (
	"context"
	"io"
	"net/http"
	"time"
)

func newEnrichPutPolicyFunc(t Transport) EnrichPutPolicy {
	_ = "STUB: not implemented"
	return *new(EnrichPutPolicy)
}

type EnrichPutPolicy func(name string, body io.Reader, o ...func(*EnrichPutPolicyRequest)) (*Response, error)

type EnrichPutPolicyRequest struct {
	Body io.Reader

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

func (r EnrichPutPolicyRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f EnrichPutPolicy) WithContext(v context.Context) func(*EnrichPutPolicyRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f EnrichPutPolicy) WithMasterTimeout(v time.Duration) func(*EnrichPutPolicyRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f EnrichPutPolicy) WithPretty() func(*EnrichPutPolicyRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f EnrichPutPolicy) WithHuman() func(*EnrichPutPolicyRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f EnrichPutPolicy) WithErrorTrace() func(*EnrichPutPolicyRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f EnrichPutPolicy) WithFilterPath(v ...string) func(*EnrichPutPolicyRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f EnrichPutPolicy) WithHeader(h map[string]string) func(*EnrichPutPolicyRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f EnrichPutPolicy) WithOpaqueID(s string) func(*EnrichPutPolicyRequest) {
	_ = "STUB: not implemented"
	return nil
}
