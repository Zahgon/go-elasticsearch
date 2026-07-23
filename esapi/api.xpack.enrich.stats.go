package esapi

import (
	"context"
	"net/http"
	"time"
)

func newEnrichStatsFunc(t Transport) EnrichStats {
	_ = "STUB: not implemented"
	return *new(EnrichStats)
}

type EnrichStats func(o ...func(*EnrichStatsRequest)) (*Response, error)

type EnrichStatsRequest struct {
	MasterTimeout time.Duration

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r EnrichStatsRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f EnrichStats) WithContext(v context.Context) func(*EnrichStatsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f EnrichStats) WithMasterTimeout(v time.Duration) func(*EnrichStatsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f EnrichStats) WithPretty() func(*EnrichStatsRequest) { _ = "STUB: not implemented"; return nil }

func (f EnrichStats) WithHuman() func(*EnrichStatsRequest) { _ = "STUB: not implemented"; return nil }

func (f EnrichStats) WithErrorTrace() func(*EnrichStatsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f EnrichStats) WithFilterPath(v ...string) func(*EnrichStatsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f EnrichStats) WithHeader(h map[string]string) func(*EnrichStatsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f EnrichStats) WithOpaqueID(s string) func(*EnrichStatsRequest) {
	_ = "STUB: not implemented"
	return nil
}
