package esapi

import (
	"context"
	"net/http"
)

func newIndicesGetFieldMappingFunc(t Transport) IndicesGetFieldMapping {
	_ = "STUB: not implemented"
	return *new(IndicesGetFieldMapping)
}

type IndicesGetFieldMapping func(fields []string, o ...func(*IndicesGetFieldMappingRequest)) (*Response, error)

type IndicesGetFieldMappingRequest struct {
	Index []string

	Fields []string

	AllowNoIndices    *bool
	ExpandWildcards   []string
	IgnoreUnavailable *bool
	IncludeDefaults   *bool

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r IndicesGetFieldMappingRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f IndicesGetFieldMapping) WithContext(v context.Context) func(*IndicesGetFieldMappingRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesGetFieldMapping) WithIndex(v ...string) func(*IndicesGetFieldMappingRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesGetFieldMapping) WithAllowNoIndices(v bool) func(*IndicesGetFieldMappingRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesGetFieldMapping) WithExpandWildcards(v ...string) func(*IndicesGetFieldMappingRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesGetFieldMapping) WithIgnoreUnavailable(v bool) func(*IndicesGetFieldMappingRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesGetFieldMapping) WithIncludeDefaults(v bool) func(*IndicesGetFieldMappingRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesGetFieldMapping) WithPretty() func(*IndicesGetFieldMappingRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesGetFieldMapping) WithHuman() func(*IndicesGetFieldMappingRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesGetFieldMapping) WithErrorTrace() func(*IndicesGetFieldMappingRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesGetFieldMapping) WithFilterPath(v ...string) func(*IndicesGetFieldMappingRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesGetFieldMapping) WithHeader(h map[string]string) func(*IndicesGetFieldMappingRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesGetFieldMapping) WithOpaqueID(s string) func(*IndicesGetFieldMappingRequest) {
	_ = "STUB: not implemented"
	return nil
}
