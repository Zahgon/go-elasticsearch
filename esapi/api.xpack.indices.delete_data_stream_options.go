package esapi

import (
	"context"
	"net/http"
	"time"
)

func newIndicesDeleteDataStreamOptionsFunc(t Transport) IndicesDeleteDataStreamOptions {
	_ = "STUB: not implemented"
	return *new(IndicesDeleteDataStreamOptions)
}

type IndicesDeleteDataStreamOptions func(name []string, o ...func(*IndicesDeleteDataStreamOptionsRequest)) (*Response, error)

type IndicesDeleteDataStreamOptionsRequest struct {
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

func (r IndicesDeleteDataStreamOptionsRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f IndicesDeleteDataStreamOptions) WithContext(v context.Context) func(*IndicesDeleteDataStreamOptionsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesDeleteDataStreamOptions) WithExpandWildcards(v ...string) func(*IndicesDeleteDataStreamOptionsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesDeleteDataStreamOptions) WithMasterTimeout(v time.Duration) func(*IndicesDeleteDataStreamOptionsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesDeleteDataStreamOptions) WithTimeout(v time.Duration) func(*IndicesDeleteDataStreamOptionsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesDeleteDataStreamOptions) WithPretty() func(*IndicesDeleteDataStreamOptionsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesDeleteDataStreamOptions) WithHuman() func(*IndicesDeleteDataStreamOptionsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesDeleteDataStreamOptions) WithErrorTrace() func(*IndicesDeleteDataStreamOptionsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesDeleteDataStreamOptions) WithFilterPath(v ...string) func(*IndicesDeleteDataStreamOptionsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesDeleteDataStreamOptions) WithHeader(h map[string]string) func(*IndicesDeleteDataStreamOptionsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesDeleteDataStreamOptions) WithOpaqueID(s string) func(*IndicesDeleteDataStreamOptionsRequest) {
	_ = "STUB: not implemented"
	return nil
}
