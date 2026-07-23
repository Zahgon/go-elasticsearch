package esapi

import (
	"context"
	"net/http"
	"time"
)

func newIndicesOpenFunc(t Transport) IndicesOpen {
	_ = "STUB: not implemented"
	return *new(IndicesOpen)
}

type IndicesOpen func(index []string, o ...func(*IndicesOpenRequest)) (*Response, error)

type IndicesOpenRequest struct {
	Index []string

	AllowNoIndices      *bool
	ExpandWildcards     []string
	IgnoreUnavailable   *bool
	MasterTimeout       time.Duration
	Timeout             time.Duration
	WaitForActiveShards string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r IndicesOpenRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f IndicesOpen) WithContext(v context.Context) func(*IndicesOpenRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesOpen) WithAllowNoIndices(v bool) func(*IndicesOpenRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesOpen) WithExpandWildcards(v ...string) func(*IndicesOpenRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesOpen) WithIgnoreUnavailable(v bool) func(*IndicesOpenRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesOpen) WithMasterTimeout(v time.Duration) func(*IndicesOpenRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesOpen) WithTimeout(v time.Duration) func(*IndicesOpenRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesOpen) WithWaitForActiveShards(v string) func(*IndicesOpenRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesOpen) WithPretty() func(*IndicesOpenRequest) { _ = "STUB: not implemented"; return nil }

func (f IndicesOpen) WithHuman() func(*IndicesOpenRequest) { _ = "STUB: not implemented"; return nil }

func (f IndicesOpen) WithErrorTrace() func(*IndicesOpenRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesOpen) WithFilterPath(v ...string) func(*IndicesOpenRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesOpen) WithHeader(h map[string]string) func(*IndicesOpenRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesOpen) WithOpaqueID(s string) func(*IndicesOpenRequest) {
	_ = "STUB: not implemented"
	return nil
}
