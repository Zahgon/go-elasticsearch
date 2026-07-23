package esapi

import (
	"context"
	"net/http"
)

func newSecurityGetStatsFunc(t Transport) SecurityGetStats {
	_ = "STUB: not implemented"
	return *new(SecurityGetStats)
}

type SecurityGetStats func(o ...func(*SecurityGetStatsRequest)) (*Response, error)

type SecurityGetStatsRequest struct {
	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r SecurityGetStatsRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f SecurityGetStats) WithContext(v context.Context) func(*SecurityGetStatsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityGetStats) WithPretty() func(*SecurityGetStatsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityGetStats) WithHuman() func(*SecurityGetStatsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityGetStats) WithErrorTrace() func(*SecurityGetStatsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityGetStats) WithFilterPath(v ...string) func(*SecurityGetStatsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityGetStats) WithHeader(h map[string]string) func(*SecurityGetStatsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SecurityGetStats) WithOpaqueID(s string) func(*SecurityGetStatsRequest) {
	_ = "STUB: not implemented"
	return nil
}
