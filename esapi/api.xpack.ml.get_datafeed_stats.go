package esapi

import (
	"context"
	"net/http"
)

func newMLGetDatafeedStatsFunc(t Transport) MLGetDatafeedStats {
	_ = "STUB: not implemented"
	return *new(MLGetDatafeedStats)
}

type MLGetDatafeedStats func(o ...func(*MLGetDatafeedStatsRequest)) (*Response, error)

type MLGetDatafeedStatsRequest struct {
	DatafeedID []string

	AllowNoMatch *bool

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r MLGetDatafeedStatsRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f MLGetDatafeedStats) WithContext(v context.Context) func(*MLGetDatafeedStatsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLGetDatafeedStats) WithDatafeedID(v ...string) func(*MLGetDatafeedStatsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLGetDatafeedStats) WithAllowNoMatch(v bool) func(*MLGetDatafeedStatsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLGetDatafeedStats) WithPretty() func(*MLGetDatafeedStatsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLGetDatafeedStats) WithHuman() func(*MLGetDatafeedStatsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLGetDatafeedStats) WithErrorTrace() func(*MLGetDatafeedStatsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLGetDatafeedStats) WithFilterPath(v ...string) func(*MLGetDatafeedStatsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLGetDatafeedStats) WithHeader(h map[string]string) func(*MLGetDatafeedStatsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLGetDatafeedStats) WithOpaqueID(s string) func(*MLGetDatafeedStatsRequest) {
	_ = "STUB: not implemented"
	return nil
}
