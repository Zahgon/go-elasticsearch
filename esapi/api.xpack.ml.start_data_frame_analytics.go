package esapi

import (
	"context"
	"io"
	"net/http"
	"time"
)

func newMLStartDataFrameAnalyticsFunc(t Transport) MLStartDataFrameAnalytics {
	_ = "STUB: not implemented"
	return *new(MLStartDataFrameAnalytics)
}

type MLStartDataFrameAnalytics func(id string, o ...func(*MLStartDataFrameAnalyticsRequest)) (*Response, error)

type MLStartDataFrameAnalyticsRequest struct {
	ID string

	Body io.Reader

	Timeout time.Duration

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r MLStartDataFrameAnalyticsRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f MLStartDataFrameAnalytics) WithContext(v context.Context) func(*MLStartDataFrameAnalyticsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLStartDataFrameAnalytics) WithBody(v io.Reader) func(*MLStartDataFrameAnalyticsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLStartDataFrameAnalytics) WithTimeout(v time.Duration) func(*MLStartDataFrameAnalyticsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLStartDataFrameAnalytics) WithPretty() func(*MLStartDataFrameAnalyticsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLStartDataFrameAnalytics) WithHuman() func(*MLStartDataFrameAnalyticsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLStartDataFrameAnalytics) WithErrorTrace() func(*MLStartDataFrameAnalyticsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLStartDataFrameAnalytics) WithFilterPath(v ...string) func(*MLStartDataFrameAnalyticsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLStartDataFrameAnalytics) WithHeader(h map[string]string) func(*MLStartDataFrameAnalyticsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLStartDataFrameAnalytics) WithOpaqueID(s string) func(*MLStartDataFrameAnalyticsRequest) {
	_ = "STUB: not implemented"
	return nil
}
