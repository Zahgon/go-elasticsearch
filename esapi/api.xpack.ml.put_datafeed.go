package esapi

import (
	"context"
	"io"
	"net/http"
)

func newMLPutDatafeedFunc(t Transport) MLPutDatafeed {
	_ = "STUB: not implemented"
	return *new(MLPutDatafeed)
}

type MLPutDatafeed func(body io.Reader, datafeed_id string, o ...func(*MLPutDatafeedRequest)) (*Response, error)

type MLPutDatafeedRequest struct {
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

func (r MLPutDatafeedRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f MLPutDatafeed) WithContext(v context.Context) func(*MLPutDatafeedRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLPutDatafeed) WithAllowNoIndices(v bool) func(*MLPutDatafeedRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLPutDatafeed) WithExpandWildcards(v ...string) func(*MLPutDatafeedRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLPutDatafeed) WithIgnoreThrottled(v bool) func(*MLPutDatafeedRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLPutDatafeed) WithIgnoreUnavailable(v bool) func(*MLPutDatafeedRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLPutDatafeed) WithPretty() func(*MLPutDatafeedRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLPutDatafeed) WithHuman() func(*MLPutDatafeedRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLPutDatafeed) WithErrorTrace() func(*MLPutDatafeedRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLPutDatafeed) WithFilterPath(v ...string) func(*MLPutDatafeedRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLPutDatafeed) WithHeader(h map[string]string) func(*MLPutDatafeedRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLPutDatafeed) WithOpaqueID(s string) func(*MLPutDatafeedRequest) {
	_ = "STUB: not implemented"
	return nil
}
