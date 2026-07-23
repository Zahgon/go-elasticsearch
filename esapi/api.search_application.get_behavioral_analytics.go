package esapi

import (
	"context"
	"net/http"
)

func newSearchApplicationGetBehavioralAnalyticsFunc(t Transport) SearchApplicationGetBehavioralAnalytics {
	_ = "STUB: not implemented"
	return *new(SearchApplicationGetBehavioralAnalytics)
}

type SearchApplicationGetBehavioralAnalytics func(o ...func(*SearchApplicationGetBehavioralAnalyticsRequest)) (*Response, error)

type SearchApplicationGetBehavioralAnalyticsRequest struct {
	Name []string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r SearchApplicationGetBehavioralAnalyticsRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f SearchApplicationGetBehavioralAnalytics) WithContext(v context.Context) func(*SearchApplicationGetBehavioralAnalyticsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SearchApplicationGetBehavioralAnalytics) WithName(v ...string) func(*SearchApplicationGetBehavioralAnalyticsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SearchApplicationGetBehavioralAnalytics) WithPretty() func(*SearchApplicationGetBehavioralAnalyticsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SearchApplicationGetBehavioralAnalytics) WithHuman() func(*SearchApplicationGetBehavioralAnalyticsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SearchApplicationGetBehavioralAnalytics) WithErrorTrace() func(*SearchApplicationGetBehavioralAnalyticsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SearchApplicationGetBehavioralAnalytics) WithFilterPath(v ...string) func(*SearchApplicationGetBehavioralAnalyticsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SearchApplicationGetBehavioralAnalytics) WithHeader(h map[string]string) func(*SearchApplicationGetBehavioralAnalyticsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SearchApplicationGetBehavioralAnalytics) WithOpaqueID(s string) func(*SearchApplicationGetBehavioralAnalyticsRequest) {
	_ = "STUB: not implemented"
	return nil
}
