package esapi

import (
	"context"
	"io"
	"net/http"
	"time"
)

func newFleetSearchFunc(t Transport) FleetSearch {
	_ = "STUB: not implemented"
	return *new(FleetSearch)
}

type FleetSearch func(index string, body io.Reader, o ...func(*FleetSearchRequest)) (*Response, error)

type FleetSearchRequest struct {
	Index string

	Body io.Reader

	AllowPartialSearchResults *bool
	WaitForCheckpoints        []string
	WaitForCheckpointsTimeout time.Duration

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r FleetSearchRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f FleetSearch) WithContext(v context.Context) func(*FleetSearchRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f FleetSearch) WithAllowPartialSearchResults(v bool) func(*FleetSearchRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f FleetSearch) WithWaitForCheckpoints(v ...string) func(*FleetSearchRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f FleetSearch) WithWaitForCheckpointsTimeout(v time.Duration) func(*FleetSearchRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f FleetSearch) WithPretty() func(*FleetSearchRequest) { _ = "STUB: not implemented"; return nil }

func (f FleetSearch) WithHuman() func(*FleetSearchRequest) { _ = "STUB: not implemented"; return nil }

func (f FleetSearch) WithErrorTrace() func(*FleetSearchRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f FleetSearch) WithFilterPath(v ...string) func(*FleetSearchRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f FleetSearch) WithHeader(h map[string]string) func(*FleetSearchRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f FleetSearch) WithOpaqueID(s string) func(*FleetSearchRequest) {
	_ = "STUB: not implemented"
	return nil
}
