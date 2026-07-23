package esapi

import (
	"context"
	"net/http"
	"time"
)

func newHealthReportFunc(t Transport) HealthReport {
	_ = "STUB: not implemented"
	return *new(HealthReport)
}

type HealthReport func(o ...func(*HealthReportRequest)) (*Response, error)

type HealthReportRequest struct {
	Feature []string

	Size    *int
	Timeout time.Duration
	Verbose *bool

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r HealthReportRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f HealthReport) WithContext(v context.Context) func(*HealthReportRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f HealthReport) WithFeature(v ...string) func(*HealthReportRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f HealthReport) WithSize(v int) func(*HealthReportRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f HealthReport) WithTimeout(v time.Duration) func(*HealthReportRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f HealthReport) WithVerbose(v bool) func(*HealthReportRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f HealthReport) WithPretty() func(*HealthReportRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f HealthReport) WithHuman() func(*HealthReportRequest) { _ = "STUB: not implemented"; return nil }

func (f HealthReport) WithErrorTrace() func(*HealthReportRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f HealthReport) WithFilterPath(v ...string) func(*HealthReportRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f HealthReport) WithHeader(h map[string]string) func(*HealthReportRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f HealthReport) WithOpaqueID(s string) func(*HealthReportRequest) {
	_ = "STUB: not implemented"
	return nil
}
