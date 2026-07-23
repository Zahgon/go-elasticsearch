package esapi

import (
	"context"
	"net/http"
	"time"
)

func newIndicesAddBlockFunc(t Transport) IndicesAddBlock {
	_ = "STUB: not implemented"
	return *new(IndicesAddBlock)
}

type IndicesAddBlock func(index []string, block string, o ...func(*IndicesAddBlockRequest)) (*Response, error)

type IndicesAddBlockRequest struct {
	Index []string

	Block string

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

func (r IndicesAddBlockRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f IndicesAddBlock) WithContext(v context.Context) func(*IndicesAddBlockRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesAddBlock) WithAllowNoIndices(v bool) func(*IndicesAddBlockRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesAddBlock) WithExpandWildcards(v ...string) func(*IndicesAddBlockRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesAddBlock) WithIgnoreUnavailable(v bool) func(*IndicesAddBlockRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesAddBlock) WithMasterTimeout(v time.Duration) func(*IndicesAddBlockRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesAddBlock) WithTimeout(v time.Duration) func(*IndicesAddBlockRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesAddBlock) WithPretty() func(*IndicesAddBlockRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesAddBlock) WithHuman() func(*IndicesAddBlockRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesAddBlock) WithErrorTrace() func(*IndicesAddBlockRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesAddBlock) WithFilterPath(v ...string) func(*IndicesAddBlockRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesAddBlock) WithHeader(h map[string]string) func(*IndicesAddBlockRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesAddBlock) WithOpaqueID(s string) func(*IndicesAddBlockRequest) {
	_ = "STUB: not implemented"
	return nil
}
