package esapi

import (
	"context"
	"net/http"
)

func newIndicesResolveIndexFunc(t Transport) IndicesResolveIndex {
	_ = "STUB: not implemented"
	return *new(IndicesResolveIndex)
}

type IndicesResolveIndex func(name []string, o ...func(*IndicesResolveIndexRequest)) (*Response, error)

type IndicesResolveIndexRequest struct {
	Name []string

	AllowNoIndices    *bool
	ExpandWildcards   []string
	IgnoreUnavailable *bool
	Mode              []string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r IndicesResolveIndexRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f IndicesResolveIndex) WithContext(v context.Context) func(*IndicesResolveIndexRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesResolveIndex) WithAllowNoIndices(v bool) func(*IndicesResolveIndexRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesResolveIndex) WithExpandWildcards(v ...string) func(*IndicesResolveIndexRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesResolveIndex) WithIgnoreUnavailable(v bool) func(*IndicesResolveIndexRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesResolveIndex) WithMode(v ...string) func(*IndicesResolveIndexRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesResolveIndex) WithPretty() func(*IndicesResolveIndexRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesResolveIndex) WithHuman() func(*IndicesResolveIndexRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesResolveIndex) WithErrorTrace() func(*IndicesResolveIndexRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesResolveIndex) WithFilterPath(v ...string) func(*IndicesResolveIndexRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesResolveIndex) WithHeader(h map[string]string) func(*IndicesResolveIndexRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesResolveIndex) WithOpaqueID(s string) func(*IndicesResolveIndexRequest) {
	_ = "STUB: not implemented"
	return nil
}
