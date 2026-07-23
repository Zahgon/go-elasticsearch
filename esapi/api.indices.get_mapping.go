package esapi

import (
	"context"
	"net/http"
	"time"
)

func newIndicesGetMappingFunc(t Transport) IndicesGetMapping {
	_ = "STUB: not implemented"
	return *new(IndicesGetMapping)
}

type IndicesGetMapping func(o ...func(*IndicesGetMappingRequest)) (*Response, error)

type IndicesGetMappingRequest struct {
	Index []string

	AllowNoIndices    *bool
	ExpandWildcards   []string
	IgnoreUnavailable *bool
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

func (r IndicesGetMappingRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f IndicesGetMapping) WithContext(v context.Context) func(*IndicesGetMappingRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesGetMapping) WithIndex(v ...string) func(*IndicesGetMappingRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesGetMapping) WithAllowNoIndices(v bool) func(*IndicesGetMappingRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesGetMapping) WithExpandWildcards(v ...string) func(*IndicesGetMappingRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesGetMapping) WithIgnoreUnavailable(v bool) func(*IndicesGetMappingRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesGetMapping) WithLocal(v bool) func(*IndicesGetMappingRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesGetMapping) WithMasterTimeout(v time.Duration) func(*IndicesGetMappingRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesGetMapping) WithPretty() func(*IndicesGetMappingRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesGetMapping) WithHuman() func(*IndicesGetMappingRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesGetMapping) WithErrorTrace() func(*IndicesGetMappingRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesGetMapping) WithFilterPath(v ...string) func(*IndicesGetMappingRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesGetMapping) WithHeader(h map[string]string) func(*IndicesGetMappingRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesGetMapping) WithOpaqueID(s string) func(*IndicesGetMappingRequest) {
	_ = "STUB: not implemented"
	return nil
}
