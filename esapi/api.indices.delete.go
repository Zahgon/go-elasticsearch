package esapi

import (
	"context"
	"net/http"
	"time"
)

func newIndicesDeleteFunc(t Transport) IndicesDelete {
	_ = "STUB: not implemented"
	return *new(IndicesDelete)
}

type IndicesDelete func(index []string, o ...func(*IndicesDeleteRequest)) (*Response, error)

type IndicesDeleteRequest struct {
	Index []string

	AllowNoIndices    *bool
	ExpandWildcards   []string
	IgnoreUnavailable *bool
	MasterTimeout     time.Duration
	Timeout           time.Duration

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r IndicesDeleteRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f IndicesDelete) WithContext(v context.Context) func(*IndicesDeleteRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesDelete) WithAllowNoIndices(v bool) func(*IndicesDeleteRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesDelete) WithExpandWildcards(v ...string) func(*IndicesDeleteRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesDelete) WithIgnoreUnavailable(v bool) func(*IndicesDeleteRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesDelete) WithMasterTimeout(v time.Duration) func(*IndicesDeleteRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesDelete) WithTimeout(v time.Duration) func(*IndicesDeleteRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesDelete) WithPretty() func(*IndicesDeleteRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesDelete) WithHuman() func(*IndicesDeleteRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesDelete) WithErrorTrace() func(*IndicesDeleteRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesDelete) WithFilterPath(v ...string) func(*IndicesDeleteRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesDelete) WithHeader(h map[string]string) func(*IndicesDeleteRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesDelete) WithOpaqueID(s string) func(*IndicesDeleteRequest) {
	_ = "STUB: not implemented"
	return nil
}
