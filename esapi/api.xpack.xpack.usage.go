package esapi

import (
	"context"
	"net/http"
	"time"
)

func newXPackUsageFunc(t Transport) XPackUsage { _ = "STUB: not implemented"; return *new(XPackUsage) }

type XPackUsage func(o ...func(*XPackUsageRequest)) (*Response, error)

type XPackUsageRequest struct {
	MasterTimeout time.Duration

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r XPackUsageRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f XPackUsage) WithContext(v context.Context) func(*XPackUsageRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f XPackUsage) WithMasterTimeout(v time.Duration) func(*XPackUsageRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f XPackUsage) WithPretty() func(*XPackUsageRequest) { _ = "STUB: not implemented"; return nil }

func (f XPackUsage) WithHuman() func(*XPackUsageRequest) { _ = "STUB: not implemented"; return nil }

func (f XPackUsage) WithErrorTrace() func(*XPackUsageRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f XPackUsage) WithFilterPath(v ...string) func(*XPackUsageRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f XPackUsage) WithHeader(h map[string]string) func(*XPackUsageRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f XPackUsage) WithOpaqueID(s string) func(*XPackUsageRequest) {
	_ = "STUB: not implemented"
	return nil
}
