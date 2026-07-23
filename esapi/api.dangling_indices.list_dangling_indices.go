package esapi

import (
	"context"
	"net/http"
)

func newDanglingIndicesListDanglingIndicesFunc(t Transport) DanglingIndicesListDanglingIndices {
	_ = "STUB: not implemented"
	return *new(DanglingIndicesListDanglingIndices)
}

type DanglingIndicesListDanglingIndices func(o ...func(*DanglingIndicesListDanglingIndicesRequest)) (*Response, error)

type DanglingIndicesListDanglingIndicesRequest struct {
	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r DanglingIndicesListDanglingIndicesRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f DanglingIndicesListDanglingIndices) WithContext(v context.Context) func(*DanglingIndicesListDanglingIndicesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f DanglingIndicesListDanglingIndices) WithPretty() func(*DanglingIndicesListDanglingIndicesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f DanglingIndicesListDanglingIndices) WithHuman() func(*DanglingIndicesListDanglingIndicesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f DanglingIndicesListDanglingIndices) WithErrorTrace() func(*DanglingIndicesListDanglingIndicesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f DanglingIndicesListDanglingIndices) WithFilterPath(v ...string) func(*DanglingIndicesListDanglingIndicesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f DanglingIndicesListDanglingIndices) WithHeader(h map[string]string) func(*DanglingIndicesListDanglingIndicesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f DanglingIndicesListDanglingIndices) WithOpaqueID(s string) func(*DanglingIndicesListDanglingIndicesRequest) {
	_ = "STUB: not implemented"
	return nil
}
