package esapi

import (
	"context"
	"net/http"
)

func newIndicesRecoveryFunc(t Transport) IndicesRecovery {
	_ = "STUB: not implemented"
	return *new(IndicesRecovery)
}

type IndicesRecovery func(o ...func(*IndicesRecoveryRequest)) (*Response, error)

type IndicesRecoveryRequest struct {
	Index []string

	ActiveOnly        *bool
	AllowNoIndices    *bool
	Detailed          *bool
	ExpandWildcards   []string
	IgnoreUnavailable *bool

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r IndicesRecoveryRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f IndicesRecovery) WithContext(v context.Context) func(*IndicesRecoveryRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesRecovery) WithIndex(v ...string) func(*IndicesRecoveryRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesRecovery) WithActiveOnly(v bool) func(*IndicesRecoveryRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesRecovery) WithAllowNoIndices(v bool) func(*IndicesRecoveryRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesRecovery) WithDetailed(v bool) func(*IndicesRecoveryRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesRecovery) WithExpandWildcards(v ...string) func(*IndicesRecoveryRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesRecovery) WithIgnoreUnavailable(v bool) func(*IndicesRecoveryRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesRecovery) WithPretty() func(*IndicesRecoveryRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesRecovery) WithHuman() func(*IndicesRecoveryRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesRecovery) WithErrorTrace() func(*IndicesRecoveryRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesRecovery) WithFilterPath(v ...string) func(*IndicesRecoveryRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesRecovery) WithHeader(h map[string]string) func(*IndicesRecoveryRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesRecovery) WithOpaqueID(s string) func(*IndicesRecoveryRequest) {
	_ = "STUB: not implemented"
	return nil
}
