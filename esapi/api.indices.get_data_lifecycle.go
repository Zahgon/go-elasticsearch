package esapi

import (
	"context"
	"net/http"
	"time"
)

func newIndicesGetDataLifecycleFunc(t Transport) IndicesGetDataLifecycle {
	_ = "STUB: not implemented"
	return *new(IndicesGetDataLifecycle)
}

type IndicesGetDataLifecycle func(name []string, o ...func(*IndicesGetDataLifecycleRequest)) (*Response, error)

type IndicesGetDataLifecycleRequest struct {
	Name []string

	ExpandWildcards []string
	IncludeDefaults *bool
	MasterTimeout   time.Duration

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r IndicesGetDataLifecycleRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f IndicesGetDataLifecycle) WithContext(v context.Context) func(*IndicesGetDataLifecycleRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesGetDataLifecycle) WithExpandWildcards(v ...string) func(*IndicesGetDataLifecycleRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesGetDataLifecycle) WithIncludeDefaults(v bool) func(*IndicesGetDataLifecycleRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesGetDataLifecycle) WithMasterTimeout(v time.Duration) func(*IndicesGetDataLifecycleRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesGetDataLifecycle) WithPretty() func(*IndicesGetDataLifecycleRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesGetDataLifecycle) WithHuman() func(*IndicesGetDataLifecycleRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesGetDataLifecycle) WithErrorTrace() func(*IndicesGetDataLifecycleRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesGetDataLifecycle) WithFilterPath(v ...string) func(*IndicesGetDataLifecycleRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesGetDataLifecycle) WithHeader(h map[string]string) func(*IndicesGetDataLifecycleRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesGetDataLifecycle) WithOpaqueID(s string) func(*IndicesGetDataLifecycleRequest) {
	_ = "STUB: not implemented"
	return nil
}
