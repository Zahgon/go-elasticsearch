package esapi

import (
	"context"
	"net/http"
	"time"
)

func newStreamsLogsDisableFunc(t Transport) StreamsLogsDisable {
	_ = "STUB: not implemented"
	return *new(StreamsLogsDisable)
}

type StreamsLogsDisable func(name string, o ...func(*StreamsLogsDisableRequest)) (*Response, error)

type StreamsLogsDisableRequest struct {
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

func (r StreamsLogsDisableRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f StreamsLogsDisable) WithContext(v context.Context) func(*StreamsLogsDisableRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f StreamsLogsDisable) WithMasterTimeout(v time.Duration) func(*StreamsLogsDisableRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f StreamsLogsDisable) WithTimeout(v time.Duration) func(*StreamsLogsDisableRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f StreamsLogsDisable) WithPretty() func(*StreamsLogsDisableRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f StreamsLogsDisable) WithHuman() func(*StreamsLogsDisableRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f StreamsLogsDisable) WithErrorTrace() func(*StreamsLogsDisableRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f StreamsLogsDisable) WithFilterPath(v ...string) func(*StreamsLogsDisableRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f StreamsLogsDisable) WithHeader(h map[string]string) func(*StreamsLogsDisableRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f StreamsLogsDisable) WithOpaqueID(s string) func(*StreamsLogsDisableRequest) {
	_ = "STUB: not implemented"
	return nil
}
