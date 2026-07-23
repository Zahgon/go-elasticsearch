package esapi

import (
	"context"
	"io"
	"net/http"
)

func newProfilingFlamegraphFunc(t Transport) ProfilingFlamegraph {
	_ = "STUB: not implemented"
	return *new(ProfilingFlamegraph)
}

type ProfilingFlamegraph func(body io.Reader, o ...func(*ProfilingFlamegraphRequest)) (*Response, error)

type ProfilingFlamegraphRequest struct {
	Body io.Reader

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r ProfilingFlamegraphRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f ProfilingFlamegraph) WithContext(v context.Context) func(*ProfilingFlamegraphRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ProfilingFlamegraph) WithPretty() func(*ProfilingFlamegraphRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ProfilingFlamegraph) WithHuman() func(*ProfilingFlamegraphRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ProfilingFlamegraph) WithErrorTrace() func(*ProfilingFlamegraphRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ProfilingFlamegraph) WithFilterPath(v ...string) func(*ProfilingFlamegraphRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ProfilingFlamegraph) WithHeader(h map[string]string) func(*ProfilingFlamegraphRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ProfilingFlamegraph) WithOpaqueID(s string) func(*ProfilingFlamegraphRequest) {
	_ = "STUB: not implemented"
	return nil
}
