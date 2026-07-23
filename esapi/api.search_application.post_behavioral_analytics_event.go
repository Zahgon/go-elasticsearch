package esapi

import (
	"context"
	"io"
	"net/http"
)

func newSearchApplicationPostBehavioralAnalyticsEventFunc(t Transport) SearchApplicationPostBehavioralAnalyticsEvent {
	_ = "STUB: not implemented"
	return *new(SearchApplicationPostBehavioralAnalyticsEvent)
}

type SearchApplicationPostBehavioralAnalyticsEvent func(body io.Reader, collection_name string, event_type string, o ...func(*SearchApplicationPostBehavioralAnalyticsEventRequest)) (*Response, error)

type SearchApplicationPostBehavioralAnalyticsEventRequest struct {
	Body io.Reader

	CollectionName string
	EventType      string

	Debug *bool

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r SearchApplicationPostBehavioralAnalyticsEventRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f SearchApplicationPostBehavioralAnalyticsEvent) WithContext(v context.Context) func(*SearchApplicationPostBehavioralAnalyticsEventRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SearchApplicationPostBehavioralAnalyticsEvent) WithDebug(v bool) func(*SearchApplicationPostBehavioralAnalyticsEventRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SearchApplicationPostBehavioralAnalyticsEvent) WithPretty() func(*SearchApplicationPostBehavioralAnalyticsEventRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SearchApplicationPostBehavioralAnalyticsEvent) WithHuman() func(*SearchApplicationPostBehavioralAnalyticsEventRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SearchApplicationPostBehavioralAnalyticsEvent) WithErrorTrace() func(*SearchApplicationPostBehavioralAnalyticsEventRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SearchApplicationPostBehavioralAnalyticsEvent) WithFilterPath(v ...string) func(*SearchApplicationPostBehavioralAnalyticsEventRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SearchApplicationPostBehavioralAnalyticsEvent) WithHeader(h map[string]string) func(*SearchApplicationPostBehavioralAnalyticsEventRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SearchApplicationPostBehavioralAnalyticsEvent) WithOpaqueID(s string) func(*SearchApplicationPostBehavioralAnalyticsEventRequest) {
	_ = "STUB: not implemented"
	return nil
}
