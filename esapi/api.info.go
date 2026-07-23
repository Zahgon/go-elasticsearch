package esapi

import (
	"context"
	"net/http"
)

func newInfoFunc(t Transport) Info { _ = "STUB: not implemented"; return *new(Info) }

type Info func(o ...func(*InfoRequest)) (*Response, error)

type InfoRequest struct {
	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r InfoRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f Info) WithContext(v context.Context) func(*InfoRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f Info) WithHuman() func(*InfoRequest) { _ = "STUB: not implemented"; return nil }

func (f Info) WithErrorTrace() func(*InfoRequest) { _ = "STUB: not implemented"; return nil }

func (f Info) WithFilterPath(v ...string) func(*InfoRequest) { _ = "STUB: not implemented"; return nil }

func (f Info) WithHeader(h map[string]string) func(*InfoRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f Info) WithOpaqueID(s string) func(*InfoRequest) { _ = "STUB: not implemented"; return nil }
