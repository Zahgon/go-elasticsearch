package esapi

import (
	"context"
	"net/http"
	"time"
)

func newIndicesRemoveBlockFunc(t Transport) IndicesRemoveBlock {
	_ = "STUB: not implemented"
	return *new(IndicesRemoveBlock)
}

type IndicesRemoveBlock func(index []string, block string, o ...func(*IndicesRemoveBlockRequest)) (*Response, error)

type IndicesRemoveBlockRequest struct {
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

func (r IndicesRemoveBlockRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f IndicesRemoveBlock) WithContext(v context.Context) func(*IndicesRemoveBlockRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesRemoveBlock) WithAllowNoIndices(v bool) func(*IndicesRemoveBlockRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesRemoveBlock) WithExpandWildcards(v ...string) func(*IndicesRemoveBlockRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesRemoveBlock) WithIgnoreUnavailable(v bool) func(*IndicesRemoveBlockRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesRemoveBlock) WithMasterTimeout(v time.Duration) func(*IndicesRemoveBlockRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesRemoveBlock) WithTimeout(v time.Duration) func(*IndicesRemoveBlockRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesRemoveBlock) WithPretty() func(*IndicesRemoveBlockRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesRemoveBlock) WithHuman() func(*IndicesRemoveBlockRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesRemoveBlock) WithErrorTrace() func(*IndicesRemoveBlockRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesRemoveBlock) WithFilterPath(v ...string) func(*IndicesRemoveBlockRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesRemoveBlock) WithHeader(h map[string]string) func(*IndicesRemoveBlockRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesRemoveBlock) WithOpaqueID(s string) func(*IndicesRemoveBlockRequest) {
	_ = "STUB: not implemented"
	return nil
}
