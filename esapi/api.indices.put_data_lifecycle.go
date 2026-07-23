package esapi

import (
	"context"
	"io"
	"net/http"
	"time"
)

func newIndicesPutDataLifecycleFunc(t Transport) IndicesPutDataLifecycle {
	_ = "STUB: not implemented"
	return *new(IndicesPutDataLifecycle)
}

type IndicesPutDataLifecycle func(name []string, body io.Reader, o ...func(*IndicesPutDataLifecycleRequest)) (*Response, error)

type IndicesPutDataLifecycleRequest struct {
	Body io.Reader

	Name []string

	ExpandWildcards []string
	MasterTimeout   time.Duration
	Timeout         time.Duration

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r IndicesPutDataLifecycleRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f IndicesPutDataLifecycle) WithContext(v context.Context) func(*IndicesPutDataLifecycleRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesPutDataLifecycle) WithExpandWildcards(v ...string) func(*IndicesPutDataLifecycleRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesPutDataLifecycle) WithMasterTimeout(v time.Duration) func(*IndicesPutDataLifecycleRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesPutDataLifecycle) WithTimeout(v time.Duration) func(*IndicesPutDataLifecycleRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesPutDataLifecycle) WithPretty() func(*IndicesPutDataLifecycleRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesPutDataLifecycle) WithHuman() func(*IndicesPutDataLifecycleRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesPutDataLifecycle) WithErrorTrace() func(*IndicesPutDataLifecycleRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesPutDataLifecycle) WithFilterPath(v ...string) func(*IndicesPutDataLifecycleRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesPutDataLifecycle) WithHeader(h map[string]string) func(*IndicesPutDataLifecycleRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesPutDataLifecycle) WithOpaqueID(s string) func(*IndicesPutDataLifecycleRequest) {
	_ = "STUB: not implemented"
	return nil
}
