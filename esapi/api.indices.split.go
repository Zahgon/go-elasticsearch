package esapi

import (
	"context"
	"io"
	"net/http"
	"time"
)

func newIndicesSplitFunc(t Transport) IndicesSplit {
	_ = "STUB: not implemented"
	return *new(IndicesSplit)
}

type IndicesSplit func(index string, body io.Reader, target string, o ...func(*IndicesSplitRequest)) (*Response, error)

type IndicesSplitRequest struct {
	Index string

	Body io.Reader

	Target string

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

func (r IndicesSplitRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f IndicesSplit) WithContext(v context.Context) func(*IndicesSplitRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesSplit) WithMasterTimeout(v time.Duration) func(*IndicesSplitRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesSplit) WithTimeout(v time.Duration) func(*IndicesSplitRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesSplit) WithWaitForActiveShards(v string) func(*IndicesSplitRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesSplit) WithPretty() func(*IndicesSplitRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesSplit) WithHuman() func(*IndicesSplitRequest) { _ = "STUB: not implemented"; return nil }

func (f IndicesSplit) WithErrorTrace() func(*IndicesSplitRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesSplit) WithFilterPath(v ...string) func(*IndicesSplitRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesSplit) WithHeader(h map[string]string) func(*IndicesSplitRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesSplit) WithOpaqueID(s string) func(*IndicesSplitRequest) {
	_ = "STUB: not implemented"
	return nil
}
