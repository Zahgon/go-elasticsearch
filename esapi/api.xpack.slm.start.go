package esapi

import (
	"context"
	"net/http"
	"time"
)

func newSlmStartFunc(t Transport) SlmStart { _ = "STUB: not implemented"; return *new(SlmStart) }

type SlmStart func(o ...func(*SlmStartRequest)) (*Response, error)

type SlmStartRequest struct {
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

func (r SlmStartRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f SlmStart) WithContext(v context.Context) func(*SlmStartRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SlmStart) WithMasterTimeout(v time.Duration) func(*SlmStartRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SlmStart) WithTimeout(v time.Duration) func(*SlmStartRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SlmStart) WithPretty() func(*SlmStartRequest) { _ = "STUB: not implemented"; return nil }

func (f SlmStart) WithHuman() func(*SlmStartRequest) { _ = "STUB: not implemented"; return nil }

func (f SlmStart) WithErrorTrace() func(*SlmStartRequest) { _ = "STUB: not implemented"; return nil }

func (f SlmStart) WithFilterPath(v ...string) func(*SlmStartRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SlmStart) WithHeader(h map[string]string) func(*SlmStartRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SlmStart) WithOpaqueID(s string) func(*SlmStartRequest) {
	_ = "STUB: not implemented"
	return nil
}
