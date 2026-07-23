package esapi

import (
	"context"
	"io"
	"net/http"
	"time"
)

func newMLStopDatafeedFunc(t Transport) MLStopDatafeed {
	_ = "STUB: not implemented"
	return *new(MLStopDatafeed)
}

type MLStopDatafeed func(datafeed_id string, o ...func(*MLStopDatafeedRequest)) (*Response, error)

type MLStopDatafeedRequest struct {
	Body io.Reader

	DatafeedID string

	AllowNoMatch *bool
	CloseJob     *bool
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

func (r MLStopDatafeedRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f MLStopDatafeed) WithContext(v context.Context) func(*MLStopDatafeedRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLStopDatafeed) WithBody(v io.Reader) func(*MLStopDatafeedRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLStopDatafeed) WithAllowNoMatch(v bool) func(*MLStopDatafeedRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLStopDatafeed) WithCloseJob(v bool) func(*MLStopDatafeedRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLStopDatafeed) WithForce(v bool) func(*MLStopDatafeedRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLStopDatafeed) WithTimeout(v time.Duration) func(*MLStopDatafeedRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLStopDatafeed) WithPretty() func(*MLStopDatafeedRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLStopDatafeed) WithHuman() func(*MLStopDatafeedRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLStopDatafeed) WithErrorTrace() func(*MLStopDatafeedRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLStopDatafeed) WithFilterPath(v ...string) func(*MLStopDatafeedRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLStopDatafeed) WithHeader(h map[string]string) func(*MLStopDatafeedRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLStopDatafeed) WithOpaqueID(s string) func(*MLStopDatafeedRequest) {
	_ = "STUB: not implemented"
	return nil
}
