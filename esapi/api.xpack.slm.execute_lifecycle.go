package esapi

import (
	"context"
	"net/http"
	"time"
)

func newSlmExecuteLifecycleFunc(t Transport) SlmExecuteLifecycle {
	_ = "STUB: not implemented"
	return *new(SlmExecuteLifecycle)
}

type SlmExecuteLifecycle func(policy_id string, o ...func(*SlmExecuteLifecycleRequest)) (*Response, error)

type SlmExecuteLifecycleRequest struct {
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

func (r SlmExecuteLifecycleRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f SlmExecuteLifecycle) WithContext(v context.Context) func(*SlmExecuteLifecycleRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SlmExecuteLifecycle) WithMasterTimeout(v time.Duration) func(*SlmExecuteLifecycleRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SlmExecuteLifecycle) WithTimeout(v time.Duration) func(*SlmExecuteLifecycleRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SlmExecuteLifecycle) WithPretty() func(*SlmExecuteLifecycleRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SlmExecuteLifecycle) WithHuman() func(*SlmExecuteLifecycleRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SlmExecuteLifecycle) WithErrorTrace() func(*SlmExecuteLifecycleRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SlmExecuteLifecycle) WithFilterPath(v ...string) func(*SlmExecuteLifecycleRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SlmExecuteLifecycle) WithHeader(h map[string]string) func(*SlmExecuteLifecycleRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SlmExecuteLifecycle) WithOpaqueID(s string) func(*SlmExecuteLifecycleRequest) {
	_ = "STUB: not implemented"
	return nil
}
