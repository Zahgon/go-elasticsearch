package esapi

import (
	"context"
	"io"
	"net/http"
)

func newMLPutDataFrameAnalyticsFunc(t Transport) MLPutDataFrameAnalytics {
	_ = "STUB: not implemented"
	return *new(MLPutDataFrameAnalytics)
}

type MLPutDataFrameAnalytics func(id string, body io.Reader, o ...func(*MLPutDataFrameAnalyticsRequest)) (*Response, error)

type MLPutDataFrameAnalyticsRequest struct {
	ID string

	Body io.Reader

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r MLPutDataFrameAnalyticsRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f MLPutDataFrameAnalytics) WithContext(v context.Context) func(*MLPutDataFrameAnalyticsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLPutDataFrameAnalytics) WithPretty() func(*MLPutDataFrameAnalyticsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLPutDataFrameAnalytics) WithHuman() func(*MLPutDataFrameAnalyticsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLPutDataFrameAnalytics) WithErrorTrace() func(*MLPutDataFrameAnalyticsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLPutDataFrameAnalytics) WithFilterPath(v ...string) func(*MLPutDataFrameAnalyticsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLPutDataFrameAnalytics) WithHeader(h map[string]string) func(*MLPutDataFrameAnalyticsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLPutDataFrameAnalytics) WithOpaqueID(s string) func(*MLPutDataFrameAnalyticsRequest) {
	_ = "STUB: not implemented"
	return nil
}
