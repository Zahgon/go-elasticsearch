package esapi

import (
	"context"
	"io"
	"net/http"
)

func newMLExplainDataFrameAnalyticsFunc(t Transport) MLExplainDataFrameAnalytics {
	_ = "STUB: not implemented"
	return *new(MLExplainDataFrameAnalytics)
}

type MLExplainDataFrameAnalytics func(o ...func(*MLExplainDataFrameAnalyticsRequest)) (*Response, error)

type MLExplainDataFrameAnalyticsRequest struct {
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

func (r MLExplainDataFrameAnalyticsRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f MLExplainDataFrameAnalytics) WithContext(v context.Context) func(*MLExplainDataFrameAnalyticsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLExplainDataFrameAnalytics) WithBody(v io.Reader) func(*MLExplainDataFrameAnalyticsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLExplainDataFrameAnalytics) WithDocumentID(v string) func(*MLExplainDataFrameAnalyticsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLExplainDataFrameAnalytics) WithPretty() func(*MLExplainDataFrameAnalyticsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLExplainDataFrameAnalytics) WithHuman() func(*MLExplainDataFrameAnalyticsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLExplainDataFrameAnalytics) WithErrorTrace() func(*MLExplainDataFrameAnalyticsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLExplainDataFrameAnalytics) WithFilterPath(v ...string) func(*MLExplainDataFrameAnalyticsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLExplainDataFrameAnalytics) WithHeader(h map[string]string) func(*MLExplainDataFrameAnalyticsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLExplainDataFrameAnalytics) WithOpaqueID(s string) func(*MLExplainDataFrameAnalyticsRequest) {
	_ = "STUB: not implemented"
	return nil
}
