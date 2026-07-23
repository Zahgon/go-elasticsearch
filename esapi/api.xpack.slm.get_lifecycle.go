package esapi

import (
	"context"
	"net/http"
	"time"
)

func newSlmGetLifecycleFunc(t Transport) SlmGetLifecycle {
	_ = "STUB: not implemented"
	return *new(SlmGetLifecycle)
}

type SlmGetLifecycle func(o ...func(*SlmGetLifecycleRequest)) (*Response, error)

type SlmGetLifecycleRequest struct {
	PolicyID []string

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

func (r SlmGetLifecycleRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f SlmGetLifecycle) WithContext(v context.Context) func(*SlmGetLifecycleRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SlmGetLifecycle) WithPolicyID(v ...string) func(*SlmGetLifecycleRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SlmGetLifecycle) WithMasterTimeout(v time.Duration) func(*SlmGetLifecycleRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SlmGetLifecycle) WithTimeout(v time.Duration) func(*SlmGetLifecycleRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SlmGetLifecycle) WithPretty() func(*SlmGetLifecycleRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SlmGetLifecycle) WithHuman() func(*SlmGetLifecycleRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SlmGetLifecycle) WithErrorTrace() func(*SlmGetLifecycleRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SlmGetLifecycle) WithFilterPath(v ...string) func(*SlmGetLifecycleRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SlmGetLifecycle) WithHeader(h map[string]string) func(*SlmGetLifecycleRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SlmGetLifecycle) WithOpaqueID(s string) func(*SlmGetLifecycleRequest) {
	_ = "STUB: not implemented"
	return nil
}
