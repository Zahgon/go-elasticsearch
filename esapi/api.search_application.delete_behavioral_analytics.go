package esapi

import (
	"context"
	"net/http"
)

func newSearchApplicationDeleteBehavioralAnalyticsFunc(t Transport) SearchApplicationDeleteBehavioralAnalytics {
	_ = "STUB: not implemented"
	return *new(SearchApplicationDeleteBehavioralAnalytics)
}

type SearchApplicationDeleteBehavioralAnalytics func(name string, o ...func(*SearchApplicationDeleteBehavioralAnalyticsRequest)) (*Response, error)

type SearchApplicationDeleteBehavioralAnalyticsRequest struct {
	Name string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r SearchApplicationDeleteBehavioralAnalyticsRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f SearchApplicationDeleteBehavioralAnalytics) WithContext(v context.Context) func(*SearchApplicationDeleteBehavioralAnalyticsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SearchApplicationDeleteBehavioralAnalytics) WithPretty() func(*SearchApplicationDeleteBehavioralAnalyticsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SearchApplicationDeleteBehavioralAnalytics) WithHuman() func(*SearchApplicationDeleteBehavioralAnalyticsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SearchApplicationDeleteBehavioralAnalytics) WithErrorTrace() func(*SearchApplicationDeleteBehavioralAnalyticsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SearchApplicationDeleteBehavioralAnalytics) WithFilterPath(v ...string) func(*SearchApplicationDeleteBehavioralAnalyticsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SearchApplicationDeleteBehavioralAnalytics) WithHeader(h map[string]string) func(*SearchApplicationDeleteBehavioralAnalyticsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SearchApplicationDeleteBehavioralAnalytics) WithOpaqueID(s string) func(*SearchApplicationDeleteBehavioralAnalyticsRequest) {
	_ = "STUB: not implemented"
	return nil
}
