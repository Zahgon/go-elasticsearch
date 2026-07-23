package esapi

import (
	"context"
	"io"
	"net/http"
	"time"
)

func newMLStopDataFrameAnalyticsFunc(t Transport) MLStopDataFrameAnalytics {
	_ = "STUB: not implemented"
	return *new(MLStopDataFrameAnalytics)
}

type MLStopDataFrameAnalytics func(id string, o ...func(*MLStopDataFrameAnalyticsRequest)) (*Response, error)

type MLStopDataFrameAnalyticsRequest struct {
	ID string

	Body io.Reader

	AllowNoMatch *bool
	Force        *bool
	Timeout      time.Duration

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r MLStopDataFrameAnalyticsRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f MLStopDataFrameAnalytics) WithContext(v context.Context) func(*MLStopDataFrameAnalyticsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLStopDataFrameAnalytics) WithBody(v io.Reader) func(*MLStopDataFrameAnalyticsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLStopDataFrameAnalytics) WithAllowNoMatch(v bool) func(*MLStopDataFrameAnalyticsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLStopDataFrameAnalytics) WithForce(v bool) func(*MLStopDataFrameAnalyticsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLStopDataFrameAnalytics) WithTimeout(v time.Duration) func(*MLStopDataFrameAnalyticsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLStopDataFrameAnalytics) WithPretty() func(*MLStopDataFrameAnalyticsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLStopDataFrameAnalytics) WithHuman() func(*MLStopDataFrameAnalyticsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLStopDataFrameAnalytics) WithErrorTrace() func(*MLStopDataFrameAnalyticsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLStopDataFrameAnalytics) WithFilterPath(v ...string) func(*MLStopDataFrameAnalyticsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLStopDataFrameAnalytics) WithHeader(h map[string]string) func(*MLStopDataFrameAnalyticsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLStopDataFrameAnalytics) WithOpaqueID(s string) func(*MLStopDataFrameAnalyticsRequest) {
	_ = "STUB: not implemented"
	return nil
}
