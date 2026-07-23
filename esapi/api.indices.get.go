package esapi

import (
	"context"
	"net/http"
	"time"
)

func newIndicesGetFunc(t Transport) IndicesGet { _ = "STUB: not implemented"; return *new(IndicesGet) }

type IndicesGet func(index []string, o ...func(*IndicesGetRequest)) (*Response, error)

type IndicesGetRequest struct {
	Index []string

	AllowNoIndices    *bool
	ExpandWildcards   []string
	Features          []string
	FlatSettings      *bool
	IgnoreUnavailable *bool
	IncludeDefaults   *bool
	Local             *bool
	MasterTimeout     time.Duration

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r IndicesGetRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f IndicesGet) WithContext(v context.Context) func(*IndicesGetRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesGet) WithAllowNoIndices(v bool) func(*IndicesGetRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesGet) WithExpandWildcards(v ...string) func(*IndicesGetRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesGet) WithFeatures(v ...string) func(*IndicesGetRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesGet) WithFlatSettings(v bool) func(*IndicesGetRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesGet) WithIgnoreUnavailable(v bool) func(*IndicesGetRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesGet) WithIncludeDefaults(v bool) func(*IndicesGetRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesGet) WithLocal(v bool) func(*IndicesGetRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesGet) WithMasterTimeout(v time.Duration) func(*IndicesGetRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesGet) WithPretty() func(*IndicesGetRequest) { _ = "STUB: not implemented"; return nil }

func (f IndicesGet) WithHuman() func(*IndicesGetRequest) { _ = "STUB: not implemented"; return nil }

func (f IndicesGet) WithErrorTrace() func(*IndicesGetRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesGet) WithFilterPath(v ...string) func(*IndicesGetRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesGet) WithHeader(h map[string]string) func(*IndicesGetRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesGet) WithOpaqueID(s string) func(*IndicesGetRequest) {
	_ = "STUB: not implemented"
	return nil
}
