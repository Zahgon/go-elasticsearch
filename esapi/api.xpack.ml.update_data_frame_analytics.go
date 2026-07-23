package esapi

import (
	"context"
	"io"
	"net/http"
)

func newMLUpdateDataFrameAnalyticsFunc(t Transport) MLUpdateDataFrameAnalytics {
	_ = "STUB: not implemented"
	return *new(MLUpdateDataFrameAnalytics)
}

type MLUpdateDataFrameAnalytics func(id string, body io.Reader, o ...func(*MLUpdateDataFrameAnalyticsRequest)) (*Response, error)

type MLUpdateDataFrameAnalyticsRequest struct {
	DocumentID string

	Body io.Reader

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r MLUpdateDataFrameAnalyticsRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f MLUpdateDataFrameAnalytics) WithContext(v context.Context) func(*MLUpdateDataFrameAnalyticsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLUpdateDataFrameAnalytics) WithPretty() func(*MLUpdateDataFrameAnalyticsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLUpdateDataFrameAnalytics) WithHuman() func(*MLUpdateDataFrameAnalyticsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLUpdateDataFrameAnalytics) WithErrorTrace() func(*MLUpdateDataFrameAnalyticsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLUpdateDataFrameAnalytics) WithFilterPath(v ...string) func(*MLUpdateDataFrameAnalyticsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLUpdateDataFrameAnalytics) WithHeader(h map[string]string) func(*MLUpdateDataFrameAnalyticsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLUpdateDataFrameAnalytics) WithOpaqueID(s string) func(*MLUpdateDataFrameAnalyticsRequest) {
	_ = "STUB: not implemented"
	return nil
}
