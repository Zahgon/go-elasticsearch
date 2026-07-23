package esapi

import (
	"context"
	"io"
	"net/http"
	"time"
)

func newIndicesShrinkFunc(t Transport) IndicesShrink {
	_ = "STUB: not implemented"
	return *new(IndicesShrink)
}

type IndicesShrink func(index string, body io.Reader, target string, o ...func(*IndicesShrinkRequest)) (*Response, error)

type IndicesShrinkRequest struct {
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

func (r IndicesShrinkRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f IndicesShrink) WithContext(v context.Context) func(*IndicesShrinkRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesShrink) WithMasterTimeout(v time.Duration) func(*IndicesShrinkRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesShrink) WithTimeout(v time.Duration) func(*IndicesShrinkRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesShrink) WithWaitForActiveShards(v string) func(*IndicesShrinkRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesShrink) WithPretty() func(*IndicesShrinkRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesShrink) WithHuman() func(*IndicesShrinkRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesShrink) WithErrorTrace() func(*IndicesShrinkRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesShrink) WithFilterPath(v ...string) func(*IndicesShrinkRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesShrink) WithHeader(h map[string]string) func(*IndicesShrinkRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesShrink) WithOpaqueID(s string) func(*IndicesShrinkRequest) {
	_ = "STUB: not implemented"
	return nil
}
