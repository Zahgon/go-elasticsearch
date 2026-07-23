package esapi

import (
	"context"
	"io"
	"net/http"
)

func newProfilingTopnFunctionsFunc(t Transport) ProfilingTopnFunctions {
	_ = "STUB: not implemented"
	return *new(ProfilingTopnFunctions)
}

type ProfilingTopnFunctions func(body io.Reader, o ...func(*ProfilingTopnFunctionsRequest)) (*Response, error)

type ProfilingTopnFunctionsRequest struct {
	Body io.Reader

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r ProfilingTopnFunctionsRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f ProfilingTopnFunctions) WithContext(v context.Context) func(*ProfilingTopnFunctionsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ProfilingTopnFunctions) WithPretty() func(*ProfilingTopnFunctionsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ProfilingTopnFunctions) WithHuman() func(*ProfilingTopnFunctionsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ProfilingTopnFunctions) WithErrorTrace() func(*ProfilingTopnFunctionsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ProfilingTopnFunctions) WithFilterPath(v ...string) func(*ProfilingTopnFunctionsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ProfilingTopnFunctions) WithHeader(h map[string]string) func(*ProfilingTopnFunctionsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ProfilingTopnFunctions) WithOpaqueID(s string) func(*ProfilingTopnFunctionsRequest) {
	_ = "STUB: not implemented"
	return nil
}
