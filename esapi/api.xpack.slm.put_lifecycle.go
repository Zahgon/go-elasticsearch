package esapi

import (
	"context"
	"io"
	"net/http"
	"time"
)

func newSlmPutLifecycleFunc(t Transport) SlmPutLifecycle {
	_ = "STUB: not implemented"
	return *new(SlmPutLifecycle)
}

type SlmPutLifecycle func(body io.Reader, policy_id string, o ...func(*SlmPutLifecycleRequest)) (*Response, error)

type SlmPutLifecycleRequest struct {
	Body io.Reader

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

func (r SlmPutLifecycleRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f SlmPutLifecycle) WithContext(v context.Context) func(*SlmPutLifecycleRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SlmPutLifecycle) WithMasterTimeout(v time.Duration) func(*SlmPutLifecycleRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SlmPutLifecycle) WithTimeout(v time.Duration) func(*SlmPutLifecycleRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SlmPutLifecycle) WithPretty() func(*SlmPutLifecycleRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SlmPutLifecycle) WithHuman() func(*SlmPutLifecycleRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SlmPutLifecycle) WithErrorTrace() func(*SlmPutLifecycleRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SlmPutLifecycle) WithFilterPath(v ...string) func(*SlmPutLifecycleRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SlmPutLifecycle) WithHeader(h map[string]string) func(*SlmPutLifecycleRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SlmPutLifecycle) WithOpaqueID(s string) func(*SlmPutLifecycleRequest) {
	_ = "STUB: not implemented"
	return nil
}
