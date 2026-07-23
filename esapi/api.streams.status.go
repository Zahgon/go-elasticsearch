package esapi

import (
	"context"
	"net/http"
	"time"
)

func newStreamsStatusFunc(t Transport) StreamsStatus {
	_ = "STUB: not implemented"
	return *new(StreamsStatus)
}

type StreamsStatus func(o ...func(*StreamsStatusRequest)) (*Response, error)

type StreamsStatusRequest struct {
	MasterTimeout time.Duration

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r StreamsStatusRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f StreamsStatus) WithContext(v context.Context) func(*StreamsStatusRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f StreamsStatus) WithMasterTimeout(v time.Duration) func(*StreamsStatusRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f StreamsStatus) WithPretty() func(*StreamsStatusRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f StreamsStatus) WithHuman() func(*StreamsStatusRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f StreamsStatus) WithErrorTrace() func(*StreamsStatusRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f StreamsStatus) WithFilterPath(v ...string) func(*StreamsStatusRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f StreamsStatus) WithHeader(h map[string]string) func(*StreamsStatusRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f StreamsStatus) WithOpaqueID(s string) func(*StreamsStatusRequest) {
	_ = "STUB: not implemented"
	return nil
}
