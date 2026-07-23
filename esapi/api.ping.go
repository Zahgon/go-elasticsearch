package esapi

import (
	"context"
	"net/http"
)

func newPingFunc(t Transport) Ping { _ = "STUB: not implemented"; return *new(Ping) }

type Ping func(o ...func(*PingRequest)) (*Response, error)

type PingRequest struct {
	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r PingRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f Ping) WithContext(v context.Context) func(*PingRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f Ping) WithPretty() func(*PingRequest) { _ = "STUB: not implemented"; return nil }

func (f Ping) WithHuman() func(*PingRequest) { _ = "STUB: not implemented"; return nil }

func (f Ping) WithErrorTrace() func(*PingRequest) { _ = "STUB: not implemented"; return nil }

func (f Ping) WithFilterPath(v ...string) func(*PingRequest) { _ = "STUB: not implemented"; return nil }

func (f Ping) WithHeader(h map[string]string) func(*PingRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f Ping) WithOpaqueID(s string) func(*PingRequest) { _ = "STUB: not implemented"; return nil }
