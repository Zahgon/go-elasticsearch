package esapi

import (
	"context"
	"io"
	"net/http"
)

func newProfilingStacktracesFunc(t Transport) ProfilingStacktraces {
	_ = "STUB: not implemented"
	return *new(ProfilingStacktraces)
}

type ProfilingStacktraces func(body io.Reader, o ...func(*ProfilingStacktracesRequest)) (*Response, error)

type ProfilingStacktracesRequest struct {
	Body io.Reader

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r ProfilingStacktracesRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f ProfilingStacktraces) WithContext(v context.Context) func(*ProfilingStacktracesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ProfilingStacktraces) WithPretty() func(*ProfilingStacktracesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ProfilingStacktraces) WithHuman() func(*ProfilingStacktracesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ProfilingStacktraces) WithErrorTrace() func(*ProfilingStacktracesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ProfilingStacktraces) WithFilterPath(v ...string) func(*ProfilingStacktracesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ProfilingStacktraces) WithHeader(h map[string]string) func(*ProfilingStacktracesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ProfilingStacktraces) WithOpaqueID(s string) func(*ProfilingStacktracesRequest) {
	_ = "STUB: not implemented"
	return nil
}
