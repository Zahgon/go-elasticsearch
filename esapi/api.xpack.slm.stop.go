package esapi

import (
	"context"
	"net/http"
	"time"
)

func newSlmStopFunc(t Transport) SlmStop { _ = "STUB: not implemented"; return *new(SlmStop) }

type SlmStop func(o ...func(*SlmStopRequest)) (*Response, error)

type SlmStopRequest struct {
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

func (r SlmStopRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f SlmStop) WithContext(v context.Context) func(*SlmStopRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SlmStop) WithMasterTimeout(v time.Duration) func(*SlmStopRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SlmStop) WithTimeout(v time.Duration) func(*SlmStopRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SlmStop) WithPretty() func(*SlmStopRequest) { _ = "STUB: not implemented"; return nil }

func (f SlmStop) WithHuman() func(*SlmStopRequest) { _ = "STUB: not implemented"; return nil }

func (f SlmStop) WithErrorTrace() func(*SlmStopRequest) { _ = "STUB: not implemented"; return nil }

func (f SlmStop) WithFilterPath(v ...string) func(*SlmStopRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SlmStop) WithHeader(h map[string]string) func(*SlmStopRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SlmStop) WithOpaqueID(s string) func(*SlmStopRequest) {
	_ = "STUB: not implemented"
	return nil
}
