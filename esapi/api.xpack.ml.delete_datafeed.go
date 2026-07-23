package esapi

import (
	"context"
	"net/http"
)

func newMLDeleteDatafeedFunc(t Transport) MLDeleteDatafeed {
	_ = "STUB: not implemented"
	return *new(MLDeleteDatafeed)
}

type MLDeleteDatafeed func(datafeed_id string, o ...func(*MLDeleteDatafeedRequest)) (*Response, error)

type MLDeleteDatafeedRequest struct {
	DatafeedID string

	Force *bool

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r MLDeleteDatafeedRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f MLDeleteDatafeed) WithContext(v context.Context) func(*MLDeleteDatafeedRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLDeleteDatafeed) WithForce(v bool) func(*MLDeleteDatafeedRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLDeleteDatafeed) WithPretty() func(*MLDeleteDatafeedRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLDeleteDatafeed) WithHuman() func(*MLDeleteDatafeedRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLDeleteDatafeed) WithErrorTrace() func(*MLDeleteDatafeedRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLDeleteDatafeed) WithFilterPath(v ...string) func(*MLDeleteDatafeedRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLDeleteDatafeed) WithHeader(h map[string]string) func(*MLDeleteDatafeedRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLDeleteDatafeed) WithOpaqueID(s string) func(*MLDeleteDatafeedRequest) {
	_ = "STUB: not implemented"
	return nil
}
