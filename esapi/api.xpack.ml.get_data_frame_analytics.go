package esapi

import (
	"context"
	"net/http"
)

func newMLGetDataFrameAnalyticsFunc(t Transport) MLGetDataFrameAnalytics {
	_ = "STUB: not implemented"
	return *new(MLGetDataFrameAnalytics)
}

type MLGetDataFrameAnalytics func(o ...func(*MLGetDataFrameAnalyticsRequest)) (*Response, error)

type MLGetDataFrameAnalyticsRequest struct {
	ID string

	AllowNoMatch     *bool
	ExcludeGenerated *bool
	From             *int
	Size             *int

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r MLGetDataFrameAnalyticsRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f MLGetDataFrameAnalytics) WithContext(v context.Context) func(*MLGetDataFrameAnalyticsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLGetDataFrameAnalytics) WithID(v string) func(*MLGetDataFrameAnalyticsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLGetDataFrameAnalytics) WithAllowNoMatch(v bool) func(*MLGetDataFrameAnalyticsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLGetDataFrameAnalytics) WithExcludeGenerated(v bool) func(*MLGetDataFrameAnalyticsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLGetDataFrameAnalytics) WithFrom(v int) func(*MLGetDataFrameAnalyticsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLGetDataFrameAnalytics) WithSize(v int) func(*MLGetDataFrameAnalyticsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLGetDataFrameAnalytics) WithPretty() func(*MLGetDataFrameAnalyticsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLGetDataFrameAnalytics) WithHuman() func(*MLGetDataFrameAnalyticsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLGetDataFrameAnalytics) WithErrorTrace() func(*MLGetDataFrameAnalyticsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLGetDataFrameAnalytics) WithFilterPath(v ...string) func(*MLGetDataFrameAnalyticsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLGetDataFrameAnalytics) WithHeader(h map[string]string) func(*MLGetDataFrameAnalyticsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLGetDataFrameAnalytics) WithOpaqueID(s string) func(*MLGetDataFrameAnalyticsRequest) {
	_ = "STUB: not implemented"
	return nil
}
