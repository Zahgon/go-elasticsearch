package esapi

import (
	"context"
	"net/http"
	"time"
)

func newIndicesCloseFunc(t Transport) IndicesClose {
	_ = "STUB: not implemented"
	return *new(IndicesClose)
}

type IndicesClose func(index []string, o ...func(*IndicesCloseRequest)) (*Response, error)

type IndicesCloseRequest struct {
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

func (r IndicesCloseRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f IndicesClose) WithContext(v context.Context) func(*IndicesCloseRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesClose) WithAllowNoIndices(v bool) func(*IndicesCloseRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesClose) WithExpandWildcards(v ...string) func(*IndicesCloseRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesClose) WithIgnoreUnavailable(v bool) func(*IndicesCloseRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesClose) WithMasterTimeout(v time.Duration) func(*IndicesCloseRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesClose) WithTimeout(v time.Duration) func(*IndicesCloseRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesClose) WithWaitForActiveShards(v string) func(*IndicesCloseRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesClose) WithPretty() func(*IndicesCloseRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesClose) WithHuman() func(*IndicesCloseRequest) { _ = "STUB: not implemented"; return nil }

func (f IndicesClose) WithErrorTrace() func(*IndicesCloseRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesClose) WithFilterPath(v ...string) func(*IndicesCloseRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesClose) WithHeader(h map[string]string) func(*IndicesCloseRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesClose) WithOpaqueID(s string) func(*IndicesCloseRequest) {
	_ = "STUB: not implemented"
	return nil
}
