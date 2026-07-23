package esapi

import (
	"context"
	"net/http"
	"time"
)

func newProfilingStatusFunc(t Transport) ProfilingStatus {
	_ = "STUB: not implemented"
	return *new(ProfilingStatus)
}

type ProfilingStatus func(o ...func(*ProfilingStatusRequest)) (*Response, error)

type ProfilingStatusRequest struct {
	MasterTimeout           time.Duration
	Timeout                 time.Duration
	WaitForResourcesCreated *bool

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r ProfilingStatusRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f ProfilingStatus) WithContext(v context.Context) func(*ProfilingStatusRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ProfilingStatus) WithMasterTimeout(v time.Duration) func(*ProfilingStatusRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ProfilingStatus) WithTimeout(v time.Duration) func(*ProfilingStatusRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ProfilingStatus) WithWaitForResourcesCreated(v bool) func(*ProfilingStatusRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ProfilingStatus) WithPretty() func(*ProfilingStatusRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ProfilingStatus) WithHuman() func(*ProfilingStatusRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ProfilingStatus) WithErrorTrace() func(*ProfilingStatusRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ProfilingStatus) WithFilterPath(v ...string) func(*ProfilingStatusRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ProfilingStatus) WithHeader(h map[string]string) func(*ProfilingStatusRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ProfilingStatus) WithOpaqueID(s string) func(*ProfilingStatusRequest) {
	_ = "STUB: not implemented"
	return nil
}
