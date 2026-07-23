package esapi

import (
	"context"
	"net/http"
	"time"
)

func newStreamsLogsEnableFunc(t Transport) StreamsLogsEnable {
	_ = "STUB: not implemented"
	return *new(StreamsLogsEnable)
}

type StreamsLogsEnable func(name string, o ...func(*StreamsLogsEnableRequest)) (*Response, error)

type StreamsLogsEnableRequest struct {
	Name string

	MasterTimeout time.Duration
	Timeout       time.Duration

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r StreamsLogsEnableRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f StreamsLogsEnable) WithContext(v context.Context) func(*StreamsLogsEnableRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f StreamsLogsEnable) WithMasterTimeout(v time.Duration) func(*StreamsLogsEnableRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f StreamsLogsEnable) WithTimeout(v time.Duration) func(*StreamsLogsEnableRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f StreamsLogsEnable) WithPretty() func(*StreamsLogsEnableRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f StreamsLogsEnable) WithHuman() func(*StreamsLogsEnableRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f StreamsLogsEnable) WithErrorTrace() func(*StreamsLogsEnableRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f StreamsLogsEnable) WithFilterPath(v ...string) func(*StreamsLogsEnableRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f StreamsLogsEnable) WithHeader(h map[string]string) func(*StreamsLogsEnableRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f StreamsLogsEnable) WithOpaqueID(s string) func(*StreamsLogsEnableRequest) {
	_ = "STUB: not implemented"
	return nil
}
