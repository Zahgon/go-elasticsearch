package esapi

import (
	"context"
	"net/http"
)

func newIngestGeoIPStatsFunc(t Transport) IngestGeoIPStats {
	_ = "STUB: not implemented"
	return *new(IngestGeoIPStats)
}

type IngestGeoIPStats func(o ...func(*IngestGeoIPStatsRequest)) (*Response, error)

type IngestGeoIPStatsRequest struct {
	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r IngestGeoIPStatsRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f IngestGeoIPStats) WithContext(v context.Context) func(*IngestGeoIPStatsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IngestGeoIPStats) WithPretty() func(*IngestGeoIPStatsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IngestGeoIPStats) WithHuman() func(*IngestGeoIPStatsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IngestGeoIPStats) WithErrorTrace() func(*IngestGeoIPStatsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IngestGeoIPStats) WithFilterPath(v ...string) func(*IngestGeoIPStatsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IngestGeoIPStats) WithHeader(h map[string]string) func(*IngestGeoIPStatsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IngestGeoIPStats) WithOpaqueID(s string) func(*IngestGeoIPStatsRequest) {
	_ = "STUB: not implemented"
	return nil
}
