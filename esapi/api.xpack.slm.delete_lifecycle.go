package esapi

import (
	"context"
	"net/http"
	"time"
)

func newSlmDeleteLifecycleFunc(t Transport) SlmDeleteLifecycle {
	_ = "STUB: not implemented"
	return *new(SlmDeleteLifecycle)
}

type SlmDeleteLifecycle func(policy_id string, o ...func(*SlmDeleteLifecycleRequest)) (*Response, error)

type SlmDeleteLifecycleRequest struct {
	PolicyID string

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

func (r SlmDeleteLifecycleRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f SlmDeleteLifecycle) WithContext(v context.Context) func(*SlmDeleteLifecycleRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SlmDeleteLifecycle) WithMasterTimeout(v time.Duration) func(*SlmDeleteLifecycleRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SlmDeleteLifecycle) WithTimeout(v time.Duration) func(*SlmDeleteLifecycleRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SlmDeleteLifecycle) WithPretty() func(*SlmDeleteLifecycleRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SlmDeleteLifecycle) WithHuman() func(*SlmDeleteLifecycleRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SlmDeleteLifecycle) WithErrorTrace() func(*SlmDeleteLifecycleRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SlmDeleteLifecycle) WithFilterPath(v ...string) func(*SlmDeleteLifecycleRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SlmDeleteLifecycle) WithHeader(h map[string]string) func(*SlmDeleteLifecycleRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SlmDeleteLifecycle) WithOpaqueID(s string) func(*SlmDeleteLifecycleRequest) {
	_ = "STUB: not implemented"
	return nil
}
