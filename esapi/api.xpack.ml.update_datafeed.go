package esapi

import (
	"context"
	"io"
	"net/http"
)

func newMLUpdateDatafeedFunc(t Transport) MLUpdateDatafeed {
	_ = "STUB: not implemented"
	return *new(MLUpdateDatafeed)
}

type MLUpdateDatafeed func(body io.Reader, datafeed_id string, o ...func(*MLUpdateDatafeedRequest)) (*Response, error)

type MLUpdateDatafeedRequest struct {
	Body io.Reader

	DatafeedID string

	AllowNoIndices    *bool
	ExpandWildcards   []string
	IgnoreThrottled   *bool
	IgnoreUnavailable *bool

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r MLUpdateDatafeedRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f MLUpdateDatafeed) WithContext(v context.Context) func(*MLUpdateDatafeedRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLUpdateDatafeed) WithAllowNoIndices(v bool) func(*MLUpdateDatafeedRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLUpdateDatafeed) WithExpandWildcards(v ...string) func(*MLUpdateDatafeedRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLUpdateDatafeed) WithIgnoreThrottled(v bool) func(*MLUpdateDatafeedRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLUpdateDatafeed) WithIgnoreUnavailable(v bool) func(*MLUpdateDatafeedRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLUpdateDatafeed) WithPretty() func(*MLUpdateDatafeedRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLUpdateDatafeed) WithHuman() func(*MLUpdateDatafeedRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLUpdateDatafeed) WithErrorTrace() func(*MLUpdateDatafeedRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLUpdateDatafeed) WithFilterPath(v ...string) func(*MLUpdateDatafeedRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLUpdateDatafeed) WithHeader(h map[string]string) func(*MLUpdateDatafeedRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLUpdateDatafeed) WithOpaqueID(s string) func(*MLUpdateDatafeedRequest) {
	_ = "STUB: not implemented"
	return nil
}
