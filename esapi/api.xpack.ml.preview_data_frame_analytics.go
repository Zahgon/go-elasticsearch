package esapi

import (
	"context"
	"io"
	"net/http"
)

func newMLPreviewDataFrameAnalyticsFunc(t Transport) MLPreviewDataFrameAnalytics {
	_ = "STUB: not implemented"
	return *new(MLPreviewDataFrameAnalytics)
}

type MLPreviewDataFrameAnalytics func(o ...func(*MLPreviewDataFrameAnalyticsRequest)) (*Response, error)

type MLPreviewDataFrameAnalyticsRequest struct {
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

func (r MLPreviewDataFrameAnalyticsRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f MLPreviewDataFrameAnalytics) WithContext(v context.Context) func(*MLPreviewDataFrameAnalyticsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLPreviewDataFrameAnalytics) WithBody(v io.Reader) func(*MLPreviewDataFrameAnalyticsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLPreviewDataFrameAnalytics) WithDocumentID(v string) func(*MLPreviewDataFrameAnalyticsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLPreviewDataFrameAnalytics) WithPretty() func(*MLPreviewDataFrameAnalyticsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLPreviewDataFrameAnalytics) WithHuman() func(*MLPreviewDataFrameAnalyticsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLPreviewDataFrameAnalytics) WithErrorTrace() func(*MLPreviewDataFrameAnalyticsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLPreviewDataFrameAnalytics) WithFilterPath(v ...string) func(*MLPreviewDataFrameAnalyticsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLPreviewDataFrameAnalytics) WithHeader(h map[string]string) func(*MLPreviewDataFrameAnalyticsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLPreviewDataFrameAnalytics) WithOpaqueID(s string) func(*MLPreviewDataFrameAnalyticsRequest) {
	_ = "STUB: not implemented"
	return nil
}
