package esapi

import (
	"context"
	"net/http"
	"time"
)

func newMLDeleteDataFrameAnalyticsFunc(t Transport) MLDeleteDataFrameAnalytics {
	_ = "STUB: not implemented"
	return *new(MLDeleteDataFrameAnalytics)
}

type MLDeleteDataFrameAnalytics func(id string, o ...func(*MLDeleteDataFrameAnalyticsRequest)) (*Response, error)

type MLDeleteDataFrameAnalyticsRequest struct {
	ID string

	Force   *bool
	Timeout time.Duration

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r MLDeleteDataFrameAnalyticsRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f MLDeleteDataFrameAnalytics) WithContext(v context.Context) func(*MLDeleteDataFrameAnalyticsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLDeleteDataFrameAnalytics) WithForce(v bool) func(*MLDeleteDataFrameAnalyticsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLDeleteDataFrameAnalytics) WithTimeout(v time.Duration) func(*MLDeleteDataFrameAnalyticsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLDeleteDataFrameAnalytics) WithPretty() func(*MLDeleteDataFrameAnalyticsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLDeleteDataFrameAnalytics) WithHuman() func(*MLDeleteDataFrameAnalyticsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLDeleteDataFrameAnalytics) WithErrorTrace() func(*MLDeleteDataFrameAnalyticsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLDeleteDataFrameAnalytics) WithFilterPath(v ...string) func(*MLDeleteDataFrameAnalyticsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLDeleteDataFrameAnalytics) WithHeader(h map[string]string) func(*MLDeleteDataFrameAnalyticsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLDeleteDataFrameAnalytics) WithOpaqueID(s string) func(*MLDeleteDataFrameAnalyticsRequest) {
	_ = "STUB: not implemented"
	return nil
}
