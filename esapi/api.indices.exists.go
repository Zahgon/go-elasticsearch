package esapi

import (
	"context"
	"net/http"
)

func newIndicesExistsFunc(t Transport) IndicesExists {
	_ = "STUB: not implemented"
	return *new(IndicesExists)
}

type IndicesExists func(index []string, o ...func(*IndicesExistsRequest)) (*Response, error)

type IndicesExistsRequest struct {
	Index []string

	AllowNoIndices    *bool
	ExpandWildcards   []string
	FlatSettings      *bool
	IgnoreUnavailable *bool
	IncludeDefaults   *bool
	Local             *bool

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r IndicesExistsRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f IndicesExists) WithContext(v context.Context) func(*IndicesExistsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesExists) WithAllowNoIndices(v bool) func(*IndicesExistsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesExists) WithExpandWildcards(v ...string) func(*IndicesExistsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesExists) WithFlatSettings(v bool) func(*IndicesExistsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesExists) WithIgnoreUnavailable(v bool) func(*IndicesExistsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesExists) WithIncludeDefaults(v bool) func(*IndicesExistsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesExists) WithLocal(v bool) func(*IndicesExistsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesExists) WithPretty() func(*IndicesExistsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesExists) WithHuman() func(*IndicesExistsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesExists) WithErrorTrace() func(*IndicesExistsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesExists) WithFilterPath(v ...string) func(*IndicesExistsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesExists) WithHeader(h map[string]string) func(*IndicesExistsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesExists) WithOpaqueID(s string) func(*IndicesExistsRequest) {
	_ = "STUB: not implemented"
	return nil
}
