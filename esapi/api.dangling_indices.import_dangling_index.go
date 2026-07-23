package esapi

import (
	"context"
	"net/http"
	"time"
)

func newDanglingIndicesImportDanglingIndexFunc(t Transport) DanglingIndicesImportDanglingIndex {
	_ = "STUB: not implemented"
	return *new(DanglingIndicesImportDanglingIndex)
}

type DanglingIndicesImportDanglingIndex func(index_uuid string, o ...func(*DanglingIndicesImportDanglingIndexRequest)) (*Response, error)

type DanglingIndicesImportDanglingIndexRequest struct {
	IndexUUID string

	AcceptDataLoss *bool
	MasterTimeout  time.Duration
	Timeout        time.Duration

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r DanglingIndicesImportDanglingIndexRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f DanglingIndicesImportDanglingIndex) WithContext(v context.Context) func(*DanglingIndicesImportDanglingIndexRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f DanglingIndicesImportDanglingIndex) WithAcceptDataLoss(v bool) func(*DanglingIndicesImportDanglingIndexRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f DanglingIndicesImportDanglingIndex) WithMasterTimeout(v time.Duration) func(*DanglingIndicesImportDanglingIndexRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f DanglingIndicesImportDanglingIndex) WithTimeout(v time.Duration) func(*DanglingIndicesImportDanglingIndexRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f DanglingIndicesImportDanglingIndex) WithPretty() func(*DanglingIndicesImportDanglingIndexRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f DanglingIndicesImportDanglingIndex) WithHuman() func(*DanglingIndicesImportDanglingIndexRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f DanglingIndicesImportDanglingIndex) WithErrorTrace() func(*DanglingIndicesImportDanglingIndexRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f DanglingIndicesImportDanglingIndex) WithFilterPath(v ...string) func(*DanglingIndicesImportDanglingIndexRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f DanglingIndicesImportDanglingIndex) WithHeader(h map[string]string) func(*DanglingIndicesImportDanglingIndexRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f DanglingIndicesImportDanglingIndex) WithOpaqueID(s string) func(*DanglingIndicesImportDanglingIndexRequest) {
	_ = "STUB: not implemented"
	return nil
}
