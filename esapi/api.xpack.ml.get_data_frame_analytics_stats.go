package esapi

import (
	"context"
	"net/http"
)

func newMLGetDataFrameAnalyticsStatsFunc(t Transport) MLGetDataFrameAnalyticsStats {
	_ = "STUB: not implemented"
	return *new(MLGetDataFrameAnalyticsStats)
}

type MLGetDataFrameAnalyticsStats func(o ...func(*MLGetDataFrameAnalyticsStatsRequest)) (*Response, error)

type MLGetDataFrameAnalyticsStatsRequest struct {
	ID string

	AllowNoMatch *bool
	From         *int
	Size         *int
	Verbose      *bool

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r MLGetDataFrameAnalyticsStatsRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f MLGetDataFrameAnalyticsStats) WithContext(v context.Context) func(*MLGetDataFrameAnalyticsStatsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLGetDataFrameAnalyticsStats) WithID(v string) func(*MLGetDataFrameAnalyticsStatsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLGetDataFrameAnalyticsStats) WithAllowNoMatch(v bool) func(*MLGetDataFrameAnalyticsStatsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLGetDataFrameAnalyticsStats) WithFrom(v int) func(*MLGetDataFrameAnalyticsStatsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLGetDataFrameAnalyticsStats) WithSize(v int) func(*MLGetDataFrameAnalyticsStatsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLGetDataFrameAnalyticsStats) WithVerbose(v bool) func(*MLGetDataFrameAnalyticsStatsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLGetDataFrameAnalyticsStats) WithPretty() func(*MLGetDataFrameAnalyticsStatsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLGetDataFrameAnalyticsStats) WithHuman() func(*MLGetDataFrameAnalyticsStatsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLGetDataFrameAnalyticsStats) WithErrorTrace() func(*MLGetDataFrameAnalyticsStatsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLGetDataFrameAnalyticsStats) WithFilterPath(v ...string) func(*MLGetDataFrameAnalyticsStatsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLGetDataFrameAnalyticsStats) WithHeader(h map[string]string) func(*MLGetDataFrameAnalyticsStatsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLGetDataFrameAnalyticsStats) WithOpaqueID(s string) func(*MLGetDataFrameAnalyticsStatsRequest) {
	_ = "STUB: not implemented"
	return nil
}
