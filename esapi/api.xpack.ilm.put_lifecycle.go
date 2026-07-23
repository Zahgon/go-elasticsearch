package esapi

import (
	"context"
	"io"
	"net/http"
	"time"
)

func newILMPutLifecycleFunc(t Transport) ILMPutLifecycle {
	_ = "STUB: not implemented"
	return *new(ILMPutLifecycle)
}

type ILMPutLifecycle func(body io.Reader, policy string, o ...func(*ILMPutLifecycleRequest)) (*Response, error)

type ILMPutLifecycleRequest struct {
	Body io.Reader

	Policy string

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

func (r ILMPutLifecycleRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f ILMPutLifecycle) WithContext(v context.Context) func(*ILMPutLifecycleRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ILMPutLifecycle) WithMasterTimeout(v time.Duration) func(*ILMPutLifecycleRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ILMPutLifecycle) WithTimeout(v time.Duration) func(*ILMPutLifecycleRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ILMPutLifecycle) WithPretty() func(*ILMPutLifecycleRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ILMPutLifecycle) WithHuman() func(*ILMPutLifecycleRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ILMPutLifecycle) WithErrorTrace() func(*ILMPutLifecycleRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ILMPutLifecycle) WithFilterPath(v ...string) func(*ILMPutLifecycleRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ILMPutLifecycle) WithHeader(h map[string]string) func(*ILMPutLifecycleRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ILMPutLifecycle) WithOpaqueID(s string) func(*ILMPutLifecycleRequest) {
	_ = "STUB: not implemented"
	return nil
}
