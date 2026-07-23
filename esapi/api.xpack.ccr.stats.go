package esapi

import (
	"context"
	"net/http"
	"time"
)

func newCCRStatsFunc(t Transport) CCRStats { _ = "STUB: not implemented"; return *new(CCRStats) }

type CCRStats func(o ...func(*CCRStatsRequest)) (*Response, error)

type CCRStatsRequest struct {
	MasterTimeout time.Duration
	Timeout       time.Duration

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r CCRStatsRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f CCRStats) WithContext(v context.Context) func(*CCRStatsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CCRStats) WithMasterTimeout(v time.Duration) func(*CCRStatsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CCRStats) WithTimeout(v time.Duration) func(*CCRStatsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CCRStats) WithPretty() func(*CCRStatsRequest) { _ = "STUB: not implemented"; return nil }

func (f CCRStats) WithHuman() func(*CCRStatsRequest) { _ = "STUB: not implemented"; return nil }

func (f CCRStats) WithErrorTrace() func(*CCRStatsRequest) { _ = "STUB: not implemented"; return nil }

func (f CCRStats) WithFilterPath(v ...string) func(*CCRStatsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CCRStats) WithHeader(h map[string]string) func(*CCRStatsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CCRStats) WithOpaqueID(s string) func(*CCRStatsRequest) {
	_ = "STUB: not implemented"
	return nil
}
