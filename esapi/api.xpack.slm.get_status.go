package esapi

import (
	"context"
	"net/http"
	"time"
)

func newSlmGetStatusFunc(t Transport) SlmGetStatus {
	_ = "STUB: not implemented"
	return *new(SlmGetStatus)
}

type SlmGetStatus func(o ...func(*SlmGetStatusRequest)) (*Response, error)

type SlmGetStatusRequest struct {
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

func (r SlmGetStatusRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f SlmGetStatus) WithContext(v context.Context) func(*SlmGetStatusRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SlmGetStatus) WithMasterTimeout(v time.Duration) func(*SlmGetStatusRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SlmGetStatus) WithTimeout(v time.Duration) func(*SlmGetStatusRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SlmGetStatus) WithPretty() func(*SlmGetStatusRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SlmGetStatus) WithHuman() func(*SlmGetStatusRequest) { _ = "STUB: not implemented"; return nil }

func (f SlmGetStatus) WithErrorTrace() func(*SlmGetStatusRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SlmGetStatus) WithFilterPath(v ...string) func(*SlmGetStatusRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SlmGetStatus) WithHeader(h map[string]string) func(*SlmGetStatusRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SlmGetStatus) WithOpaqueID(s string) func(*SlmGetStatusRequest) {
	_ = "STUB: not implemented"
	return nil
}
