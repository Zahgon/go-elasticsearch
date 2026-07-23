package esapi

import (
	"context"
	"io"
	"net/http"
)

func newFleetMsearchFunc(t Transport) FleetMsearch {
	_ = "STUB: not implemented"
	return *new(FleetMsearch)
}

type FleetMsearch func(body io.Reader, o ...func(*FleetMsearchRequest)) (*Response, error)

type FleetMsearchRequest struct {
	Index string

	Body io.Reader

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r FleetMsearchRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f FleetMsearch) WithContext(v context.Context) func(*FleetMsearchRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f FleetMsearch) WithIndex(v string) func(*FleetMsearchRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f FleetMsearch) WithPretty() func(*FleetMsearchRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f FleetMsearch) WithHuman() func(*FleetMsearchRequest) { _ = "STUB: not implemented"; return nil }

func (f FleetMsearch) WithErrorTrace() func(*FleetMsearchRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f FleetMsearch) WithFilterPath(v ...string) func(*FleetMsearchRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f FleetMsearch) WithHeader(h map[string]string) func(*FleetMsearchRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f FleetMsearch) WithOpaqueID(s string) func(*FleetMsearchRequest) {
	_ = "STUB: not implemented"
	return nil
}
