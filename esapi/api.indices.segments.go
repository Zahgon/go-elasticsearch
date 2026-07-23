package esapi

import (
	"context"
	"net/http"
)

func newIndicesSegmentsFunc(t Transport) IndicesSegments {
	_ = "STUB: not implemented"
	return *new(IndicesSegments)
}

type IndicesSegments func(o ...func(*IndicesSegmentsRequest)) (*Response, error)

type IndicesSegmentsRequest struct {
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

func (r IndicesSegmentsRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f IndicesSegments) WithContext(v context.Context) func(*IndicesSegmentsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesSegments) WithIndex(v ...string) func(*IndicesSegmentsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesSegments) WithAllowNoIndices(v bool) func(*IndicesSegmentsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesSegments) WithExpandWildcards(v ...string) func(*IndicesSegmentsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesSegments) WithIgnoreUnavailable(v bool) func(*IndicesSegmentsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesSegments) WithPretty() func(*IndicesSegmentsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesSegments) WithHuman() func(*IndicesSegmentsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesSegments) WithErrorTrace() func(*IndicesSegmentsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesSegments) WithFilterPath(v ...string) func(*IndicesSegmentsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesSegments) WithHeader(h map[string]string) func(*IndicesSegmentsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesSegments) WithOpaqueID(s string) func(*IndicesSegmentsRequest) {
	_ = "STUB: not implemented"
	return nil
}
