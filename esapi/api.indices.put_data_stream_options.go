package esapi

import (
	"context"
	"io"
	"net/http"
	"time"
)

func newIndicesPutDataStreamOptionsFunc(t Transport) IndicesPutDataStreamOptions {
	_ = "STUB: not implemented"
	return *new(IndicesPutDataStreamOptions)
}

type IndicesPutDataStreamOptions func(name []string, body io.Reader, o ...func(*IndicesPutDataStreamOptionsRequest)) (*Response, error)

type IndicesPutDataStreamOptionsRequest struct {
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

func (r IndicesPutDataStreamOptionsRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f IndicesPutDataStreamOptions) WithContext(v context.Context) func(*IndicesPutDataStreamOptionsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesPutDataStreamOptions) WithExpandWildcards(v ...string) func(*IndicesPutDataStreamOptionsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesPutDataStreamOptions) WithMasterTimeout(v time.Duration) func(*IndicesPutDataStreamOptionsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesPutDataStreamOptions) WithTimeout(v time.Duration) func(*IndicesPutDataStreamOptionsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesPutDataStreamOptions) WithPretty() func(*IndicesPutDataStreamOptionsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesPutDataStreamOptions) WithHuman() func(*IndicesPutDataStreamOptionsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesPutDataStreamOptions) WithErrorTrace() func(*IndicesPutDataStreamOptionsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesPutDataStreamOptions) WithFilterPath(v ...string) func(*IndicesPutDataStreamOptionsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesPutDataStreamOptions) WithHeader(h map[string]string) func(*IndicesPutDataStreamOptionsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesPutDataStreamOptions) WithOpaqueID(s string) func(*IndicesPutDataStreamOptionsRequest) {
	_ = "STUB: not implemented"
	return nil
}
