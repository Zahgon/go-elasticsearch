package esapi

import (
	"context"
	"net/http"
)

func newILMRemovePolicyFunc(t Transport) ILMRemovePolicy {
	_ = "STUB: not implemented"
	return *new(ILMRemovePolicy)
}

type ILMRemovePolicy func(index string, o ...func(*ILMRemovePolicyRequest)) (*Response, error)

type ILMRemovePolicyRequest struct {
	Index string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r ILMRemovePolicyRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f ILMRemovePolicy) WithContext(v context.Context) func(*ILMRemovePolicyRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ILMRemovePolicy) WithPretty() func(*ILMRemovePolicyRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ILMRemovePolicy) WithHuman() func(*ILMRemovePolicyRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ILMRemovePolicy) WithErrorTrace() func(*ILMRemovePolicyRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ILMRemovePolicy) WithFilterPath(v ...string) func(*ILMRemovePolicyRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ILMRemovePolicy) WithHeader(h map[string]string) func(*ILMRemovePolicyRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ILMRemovePolicy) WithOpaqueID(s string) func(*ILMRemovePolicyRequest) {
	_ = "STUB: not implemented"
	return nil
}
