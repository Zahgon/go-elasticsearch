package esapi

import (
	"context"
	"net/http"
)

func newIndicesReloadSearchAnalyzersFunc(t Transport) IndicesReloadSearchAnalyzers {
	_ = "STUB: not implemented"
	return *new(IndicesReloadSearchAnalyzers)
}

type IndicesReloadSearchAnalyzers func(index []string, o ...func(*IndicesReloadSearchAnalyzersRequest)) (*Response, error)

type IndicesReloadSearchAnalyzersRequest struct {
	Index []string

	AllowNoIndices    *bool
	ExpandWildcards   []string
	IgnoreUnavailable *bool
	Resource          string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r IndicesReloadSearchAnalyzersRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f IndicesReloadSearchAnalyzers) WithContext(v context.Context) func(*IndicesReloadSearchAnalyzersRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesReloadSearchAnalyzers) WithAllowNoIndices(v bool) func(*IndicesReloadSearchAnalyzersRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesReloadSearchAnalyzers) WithExpandWildcards(v ...string) func(*IndicesReloadSearchAnalyzersRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesReloadSearchAnalyzers) WithIgnoreUnavailable(v bool) func(*IndicesReloadSearchAnalyzersRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesReloadSearchAnalyzers) WithResource(v string) func(*IndicesReloadSearchAnalyzersRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesReloadSearchAnalyzers) WithPretty() func(*IndicesReloadSearchAnalyzersRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesReloadSearchAnalyzers) WithHuman() func(*IndicesReloadSearchAnalyzersRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesReloadSearchAnalyzers) WithErrorTrace() func(*IndicesReloadSearchAnalyzersRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesReloadSearchAnalyzers) WithFilterPath(v ...string) func(*IndicesReloadSearchAnalyzersRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesReloadSearchAnalyzers) WithHeader(h map[string]string) func(*IndicesReloadSearchAnalyzersRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesReloadSearchAnalyzers) WithOpaqueID(s string) func(*IndicesReloadSearchAnalyzersRequest) {
	_ = "STUB: not implemented"
	return nil
}
