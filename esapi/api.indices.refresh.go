package esapi

import (
	"context"
	"net/http"
)

func newIndicesRefreshFunc(t Transport) IndicesRefresh {
	_ = "STUB: not implemented"
	return *new(IndicesRefresh)
}

type IndicesRefresh func(o ...func(*IndicesRefreshRequest)) (*Response, error)

type IndicesRefreshRequest struct {
	Index []string

	AllowNoIndices    *bool
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

func (r IndicesRefreshRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f IndicesRefresh) WithContext(v context.Context) func(*IndicesRefreshRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesRefresh) WithIndex(v ...string) func(*IndicesRefreshRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesRefresh) WithAllowNoIndices(v bool) func(*IndicesRefreshRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesRefresh) WithExpandWildcards(v ...string) func(*IndicesRefreshRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesRefresh) WithIgnoreUnavailable(v bool) func(*IndicesRefreshRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesRefresh) WithPretty() func(*IndicesRefreshRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesRefresh) WithHuman() func(*IndicesRefreshRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesRefresh) WithErrorTrace() func(*IndicesRefreshRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesRefresh) WithFilterPath(v ...string) func(*IndicesRefreshRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesRefresh) WithHeader(h map[string]string) func(*IndicesRefreshRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesRefresh) WithOpaqueID(s string) func(*IndicesRefreshRequest) {
	_ = "STUB: not implemented"
	return nil
}
