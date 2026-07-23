package esapi

import (
	"context"
	"net/http"
)

func newMLGetJobStatsFunc(t Transport) MLGetJobStats {
	_ = "STUB: not implemented"
	return *new(MLGetJobStats)
}

type MLGetJobStats func(o ...func(*MLGetJobStatsRequest)) (*Response, error)

type MLGetJobStatsRequest struct {
	JobID string

	AllowNoMatch *bool

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r MLGetJobStatsRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f MLGetJobStats) WithContext(v context.Context) func(*MLGetJobStatsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLGetJobStats) WithJobID(v string) func(*MLGetJobStatsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLGetJobStats) WithAllowNoMatch(v bool) func(*MLGetJobStatsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLGetJobStats) WithPretty() func(*MLGetJobStatsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLGetJobStats) WithHuman() func(*MLGetJobStatsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLGetJobStats) WithErrorTrace() func(*MLGetJobStatsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLGetJobStats) WithFilterPath(v ...string) func(*MLGetJobStatsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLGetJobStats) WithHeader(h map[string]string) func(*MLGetJobStatsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLGetJobStats) WithOpaqueID(s string) func(*MLGetJobStatsRequest) {
	_ = "STUB: not implemented"
	return nil
}
