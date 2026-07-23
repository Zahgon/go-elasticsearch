package esapi

import (
	"context"
	"io"
	"net/http"
	"time"
)

func newIndicesPutMappingFunc(t Transport) IndicesPutMapping {
	_ = "STUB: not implemented"
	return *new(IndicesPutMapping)
}

type IndicesPutMapping func(index []string, body io.Reader, o ...func(*IndicesPutMappingRequest)) (*Response, error)

type IndicesPutMappingRequest struct {
	Index []string

	Body io.Reader

	AllowNoIndices    *bool
	ExpandWildcards   []string
	IgnoreUnavailable *bool
	MasterTimeout     time.Duration
	Timeout           time.Duration
	WriteIndexOnly    *bool

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r IndicesPutMappingRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f IndicesPutMapping) WithContext(v context.Context) func(*IndicesPutMappingRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesPutMapping) WithAllowNoIndices(v bool) func(*IndicesPutMappingRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesPutMapping) WithExpandWildcards(v ...string) func(*IndicesPutMappingRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesPutMapping) WithIgnoreUnavailable(v bool) func(*IndicesPutMappingRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesPutMapping) WithMasterTimeout(v time.Duration) func(*IndicesPutMappingRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesPutMapping) WithTimeout(v time.Duration) func(*IndicesPutMappingRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesPutMapping) WithWriteIndexOnly(v bool) func(*IndicesPutMappingRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesPutMapping) WithPretty() func(*IndicesPutMappingRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesPutMapping) WithHuman() func(*IndicesPutMappingRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesPutMapping) WithErrorTrace() func(*IndicesPutMappingRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesPutMapping) WithFilterPath(v ...string) func(*IndicesPutMappingRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesPutMapping) WithHeader(h map[string]string) func(*IndicesPutMappingRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesPutMapping) WithOpaqueID(s string) func(*IndicesPutMappingRequest) {
	_ = "STUB: not implemented"
	return nil
}
