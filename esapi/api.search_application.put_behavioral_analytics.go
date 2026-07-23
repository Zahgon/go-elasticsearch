package esapi

import (
	"context"
	"net/http"
)

func newSearchApplicationPutBehavioralAnalyticsFunc(t Transport) SearchApplicationPutBehavioralAnalytics {
	_ = "STUB: not implemented"
	return *new(SearchApplicationPutBehavioralAnalytics)
}

type SearchApplicationPutBehavioralAnalytics func(name string, o ...func(*SearchApplicationPutBehavioralAnalyticsRequest)) (*Response, error)

type SearchApplicationPutBehavioralAnalyticsRequest struct {
	Name string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r SearchApplicationPutBehavioralAnalyticsRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f SearchApplicationPutBehavioralAnalytics) WithContext(v context.Context) func(*SearchApplicationPutBehavioralAnalyticsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SearchApplicationPutBehavioralAnalytics) WithPretty() func(*SearchApplicationPutBehavioralAnalyticsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SearchApplicationPutBehavioralAnalytics) WithHuman() func(*SearchApplicationPutBehavioralAnalyticsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SearchApplicationPutBehavioralAnalytics) WithErrorTrace() func(*SearchApplicationPutBehavioralAnalyticsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SearchApplicationPutBehavioralAnalytics) WithFilterPath(v ...string) func(*SearchApplicationPutBehavioralAnalyticsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SearchApplicationPutBehavioralAnalytics) WithHeader(h map[string]string) func(*SearchApplicationPutBehavioralAnalyticsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SearchApplicationPutBehavioralAnalytics) WithOpaqueID(s string) func(*SearchApplicationPutBehavioralAnalyticsRequest) {
	_ = "STUB: not implemented"
	return nil
}
