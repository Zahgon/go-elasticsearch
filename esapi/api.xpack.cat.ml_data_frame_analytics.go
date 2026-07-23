package esapi

import (
	"context"
	"net/http"
)

func newCatMLDataFrameAnalyticsFunc(t Transport) CatMLDataFrameAnalytics {
	_ = "STUB: not implemented"
	return *new(CatMLDataFrameAnalytics)
}

type CatMLDataFrameAnalytics func(o ...func(*CatMLDataFrameAnalyticsRequest)) (*Response, error)

type CatMLDataFrameAnalyticsRequest struct {
	DocumentID string

	AllowNoMatch *bool
	Bytes        string
	Format       string
	H            []string
	Help         *bool
	S            []string
	Time         string
	V            *bool

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r CatMLDataFrameAnalyticsRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f CatMLDataFrameAnalytics) WithContext(v context.Context) func(*CatMLDataFrameAnalyticsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatMLDataFrameAnalytics) WithDocumentID(v string) func(*CatMLDataFrameAnalyticsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatMLDataFrameAnalytics) WithAllowNoMatch(v bool) func(*CatMLDataFrameAnalyticsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatMLDataFrameAnalytics) WithBytes(v string) func(*CatMLDataFrameAnalyticsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatMLDataFrameAnalytics) WithFormat(v string) func(*CatMLDataFrameAnalyticsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatMLDataFrameAnalytics) WithH(v ...string) func(*CatMLDataFrameAnalyticsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatMLDataFrameAnalytics) WithHelp(v bool) func(*CatMLDataFrameAnalyticsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatMLDataFrameAnalytics) WithS(v ...string) func(*CatMLDataFrameAnalyticsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatMLDataFrameAnalytics) WithTime(v string) func(*CatMLDataFrameAnalyticsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatMLDataFrameAnalytics) WithV(v bool) func(*CatMLDataFrameAnalyticsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatMLDataFrameAnalytics) WithPretty() func(*CatMLDataFrameAnalyticsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatMLDataFrameAnalytics) WithHuman() func(*CatMLDataFrameAnalyticsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatMLDataFrameAnalytics) WithErrorTrace() func(*CatMLDataFrameAnalyticsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatMLDataFrameAnalytics) WithFilterPath(v ...string) func(*CatMLDataFrameAnalyticsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatMLDataFrameAnalytics) WithHeader(h map[string]string) func(*CatMLDataFrameAnalyticsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatMLDataFrameAnalytics) WithOpaqueID(s string) func(*CatMLDataFrameAnalyticsRequest) {
	_ = "STUB: not implemented"
	return nil
}
