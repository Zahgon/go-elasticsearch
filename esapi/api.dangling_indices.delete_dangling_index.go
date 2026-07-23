package esapi

import (
	"context"
	"net/http"
	"time"
)

func newDanglingIndicesDeleteDanglingIndexFunc(t Transport) DanglingIndicesDeleteDanglingIndex {
	_ = "STUB: not implemented"
	return *new(DanglingIndicesDeleteDanglingIndex)
}

type DanglingIndicesDeleteDanglingIndex func(index_uuid string, o ...func(*DanglingIndicesDeleteDanglingIndexRequest)) (*Response, error)

type DanglingIndicesDeleteDanglingIndexRequest struct {
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

func (r DanglingIndicesDeleteDanglingIndexRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f DanglingIndicesDeleteDanglingIndex) WithContext(v context.Context) func(*DanglingIndicesDeleteDanglingIndexRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f DanglingIndicesDeleteDanglingIndex) WithAcceptDataLoss(v bool) func(*DanglingIndicesDeleteDanglingIndexRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f DanglingIndicesDeleteDanglingIndex) WithMasterTimeout(v time.Duration) func(*DanglingIndicesDeleteDanglingIndexRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f DanglingIndicesDeleteDanglingIndex) WithTimeout(v time.Duration) func(*DanglingIndicesDeleteDanglingIndexRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f DanglingIndicesDeleteDanglingIndex) WithPretty() func(*DanglingIndicesDeleteDanglingIndexRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f DanglingIndicesDeleteDanglingIndex) WithHuman() func(*DanglingIndicesDeleteDanglingIndexRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f DanglingIndicesDeleteDanglingIndex) WithErrorTrace() func(*DanglingIndicesDeleteDanglingIndexRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f DanglingIndicesDeleteDanglingIndex) WithFilterPath(v ...string) func(*DanglingIndicesDeleteDanglingIndexRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f DanglingIndicesDeleteDanglingIndex) WithHeader(h map[string]string) func(*DanglingIndicesDeleteDanglingIndexRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f DanglingIndicesDeleteDanglingIndex) WithOpaqueID(s string) func(*DanglingIndicesDeleteDanglingIndexRequest) {
	_ = "STUB: not implemented"
	return nil
}
