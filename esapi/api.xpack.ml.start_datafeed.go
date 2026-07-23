package esapi

import (
	"context"
	"io"
	"net/http"
	"time"
)

func newMLStartDatafeedFunc(t Transport) MLStartDatafeed {
	_ = "STUB: not implemented"
	return *new(MLStartDatafeed)
}

type MLStartDatafeed func(datafeed_id string, o ...func(*MLStartDatafeedRequest)) (*Response, error)

type MLStartDatafeedRequest struct {
	Body io.Reader

	DatafeedID string

	End     string
	Start   string
	Timeout time.Duration

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r MLStartDatafeedRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f MLStartDatafeed) WithContext(v context.Context) func(*MLStartDatafeedRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLStartDatafeed) WithBody(v io.Reader) func(*MLStartDatafeedRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLStartDatafeed) WithEnd(v string) func(*MLStartDatafeedRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLStartDatafeed) WithStart(v string) func(*MLStartDatafeedRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLStartDatafeed) WithTimeout(v time.Duration) func(*MLStartDatafeedRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLStartDatafeed) WithPretty() func(*MLStartDatafeedRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLStartDatafeed) WithHuman() func(*MLStartDatafeedRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLStartDatafeed) WithErrorTrace() func(*MLStartDatafeedRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLStartDatafeed) WithFilterPath(v ...string) func(*MLStartDatafeedRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLStartDatafeed) WithHeader(h map[string]string) func(*MLStartDatafeedRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLStartDatafeed) WithOpaqueID(s string) func(*MLStartDatafeedRequest) {
	_ = "STUB: not implemented"
	return nil
}
