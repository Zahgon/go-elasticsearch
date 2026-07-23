package esapi

import (
	"context"
	"net/http"
	"time"
)

func newSlmGetStatsFunc(t Transport) SlmGetStats {
	_ = "STUB: not implemented"
	return *new(SlmGetStats)
}

type SlmGetStats func(o ...func(*SlmGetStatsRequest)) (*Response, error)

type SlmGetStatsRequest struct {
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

func (r SlmGetStatsRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f SlmGetStats) WithContext(v context.Context) func(*SlmGetStatsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SlmGetStats) WithMasterTimeout(v time.Duration) func(*SlmGetStatsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SlmGetStats) WithTimeout(v time.Duration) func(*SlmGetStatsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SlmGetStats) WithPretty() func(*SlmGetStatsRequest) { _ = "STUB: not implemented"; return nil }

func (f SlmGetStats) WithHuman() func(*SlmGetStatsRequest) { _ = "STUB: not implemented"; return nil }

func (f SlmGetStats) WithErrorTrace() func(*SlmGetStatsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SlmGetStats) WithFilterPath(v ...string) func(*SlmGetStatsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SlmGetStats) WithHeader(h map[string]string) func(*SlmGetStatsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SlmGetStats) WithOpaqueID(s string) func(*SlmGetStatsRequest) {
	_ = "STUB: not implemented"
	return nil
}
