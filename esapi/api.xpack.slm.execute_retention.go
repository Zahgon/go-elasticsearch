package esapi

import (
	"context"
	"net/http"
	"time"
)

func newSlmExecuteRetentionFunc(t Transport) SlmExecuteRetention {
	_ = "STUB: not implemented"
	return *new(SlmExecuteRetention)
}

type SlmExecuteRetention func(o ...func(*SlmExecuteRetentionRequest)) (*Response, error)

type SlmExecuteRetentionRequest struct {
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

func (r SlmExecuteRetentionRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f SlmExecuteRetention) WithContext(v context.Context) func(*SlmExecuteRetentionRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SlmExecuteRetention) WithMasterTimeout(v time.Duration) func(*SlmExecuteRetentionRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SlmExecuteRetention) WithTimeout(v time.Duration) func(*SlmExecuteRetentionRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SlmExecuteRetention) WithPretty() func(*SlmExecuteRetentionRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SlmExecuteRetention) WithHuman() func(*SlmExecuteRetentionRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SlmExecuteRetention) WithErrorTrace() func(*SlmExecuteRetentionRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SlmExecuteRetention) WithFilterPath(v ...string) func(*SlmExecuteRetentionRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SlmExecuteRetention) WithHeader(h map[string]string) func(*SlmExecuteRetentionRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SlmExecuteRetention) WithOpaqueID(s string) func(*SlmExecuteRetentionRequest) {
	_ = "STUB: not implemented"
	return nil
}
