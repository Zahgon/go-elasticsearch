package esapi

import (
	"context"
	"net/http"
	"time"
)

func newIndicesGetDataStreamOptionsFunc(t Transport) IndicesGetDataStreamOptions {
	_ = "STUB: not implemented"
	return *new(IndicesGetDataStreamOptions)
}

type IndicesGetDataStreamOptions func(name []string, o ...func(*IndicesGetDataStreamOptionsRequest)) (*Response, error)

type IndicesGetDataStreamOptionsRequest struct {
	Name []string

	ExpandWildcards []string
	MasterTimeout   time.Duration

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r IndicesGetDataStreamOptionsRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f IndicesGetDataStreamOptions) WithContext(v context.Context) func(*IndicesGetDataStreamOptionsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesGetDataStreamOptions) WithExpandWildcards(v ...string) func(*IndicesGetDataStreamOptionsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesGetDataStreamOptions) WithMasterTimeout(v time.Duration) func(*IndicesGetDataStreamOptionsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesGetDataStreamOptions) WithPretty() func(*IndicesGetDataStreamOptionsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesGetDataStreamOptions) WithHuman() func(*IndicesGetDataStreamOptionsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesGetDataStreamOptions) WithErrorTrace() func(*IndicesGetDataStreamOptionsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesGetDataStreamOptions) WithFilterPath(v ...string) func(*IndicesGetDataStreamOptionsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesGetDataStreamOptions) WithHeader(h map[string]string) func(*IndicesGetDataStreamOptionsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesGetDataStreamOptions) WithOpaqueID(s string) func(*IndicesGetDataStreamOptionsRequest) {
	_ = "STUB: not implemented"
	return nil
}
