package esapi

import (
	"context"
	"net/http"
	"time"
)

func newIndicesDeleteDataLifecycleFunc(t Transport) IndicesDeleteDataLifecycle {
	_ = "STUB: not implemented"
	return *new(IndicesDeleteDataLifecycle)
}

type IndicesDeleteDataLifecycle func(name []string, o ...func(*IndicesDeleteDataLifecycleRequest)) (*Response, error)

type IndicesDeleteDataLifecycleRequest struct {
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

func (r IndicesDeleteDataLifecycleRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f IndicesDeleteDataLifecycle) WithContext(v context.Context) func(*IndicesDeleteDataLifecycleRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesDeleteDataLifecycle) WithExpandWildcards(v ...string) func(*IndicesDeleteDataLifecycleRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesDeleteDataLifecycle) WithMasterTimeout(v time.Duration) func(*IndicesDeleteDataLifecycleRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesDeleteDataLifecycle) WithTimeout(v time.Duration) func(*IndicesDeleteDataLifecycleRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesDeleteDataLifecycle) WithPretty() func(*IndicesDeleteDataLifecycleRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesDeleteDataLifecycle) WithHuman() func(*IndicesDeleteDataLifecycleRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesDeleteDataLifecycle) WithErrorTrace() func(*IndicesDeleteDataLifecycleRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesDeleteDataLifecycle) WithFilterPath(v ...string) func(*IndicesDeleteDataLifecycleRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesDeleteDataLifecycle) WithHeader(h map[string]string) func(*IndicesDeleteDataLifecycleRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesDeleteDataLifecycle) WithOpaqueID(s string) func(*IndicesDeleteDataLifecycleRequest) {
	_ = "STUB: not implemented"
	return nil
}
