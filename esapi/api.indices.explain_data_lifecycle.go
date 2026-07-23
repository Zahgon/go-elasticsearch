package esapi

import (
	"context"
	"net/http"
	"time"
)

func newIndicesExplainDataLifecycleFunc(t Transport) IndicesExplainDataLifecycle {
	_ = "STUB: not implemented"
	return *new(IndicesExplainDataLifecycle)
}

type IndicesExplainDataLifecycle func(index []string, o ...func(*IndicesExplainDataLifecycleRequest)) (*Response, error)

type IndicesExplainDataLifecycleRequest struct {
	Index []string

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

func (r IndicesExplainDataLifecycleRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f IndicesExplainDataLifecycle) WithContext(v context.Context) func(*IndicesExplainDataLifecycleRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesExplainDataLifecycle) WithIncludeDefaults(v bool) func(*IndicesExplainDataLifecycleRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesExplainDataLifecycle) WithMasterTimeout(v time.Duration) func(*IndicesExplainDataLifecycleRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesExplainDataLifecycle) WithPretty() func(*IndicesExplainDataLifecycleRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesExplainDataLifecycle) WithHuman() func(*IndicesExplainDataLifecycleRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesExplainDataLifecycle) WithErrorTrace() func(*IndicesExplainDataLifecycleRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesExplainDataLifecycle) WithFilterPath(v ...string) func(*IndicesExplainDataLifecycleRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesExplainDataLifecycle) WithHeader(h map[string]string) func(*IndicesExplainDataLifecycleRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesExplainDataLifecycle) WithOpaqueID(s string) func(*IndicesExplainDataLifecycleRequest) {
	_ = "STUB: not implemented"
	return nil
}
