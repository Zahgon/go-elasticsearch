package esapi

import (
	"context"
	"net/http"
)

func newProjectGetRoutingFunc(t Transport) ProjectGetRouting {
	_ = "STUB: not implemented"
	return *new(ProjectGetRouting)
}

type ProjectGetRouting func(name string, o ...func(*ProjectGetRoutingRequest)) (*Response, error)

type ProjectGetRoutingRequest struct {
	Name string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r ProjectGetRoutingRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f ProjectGetRouting) WithContext(v context.Context) func(*ProjectGetRoutingRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ProjectGetRouting) WithPretty() func(*ProjectGetRoutingRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ProjectGetRouting) WithHuman() func(*ProjectGetRoutingRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ProjectGetRouting) WithErrorTrace() func(*ProjectGetRoutingRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ProjectGetRouting) WithFilterPath(v ...string) func(*ProjectGetRoutingRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ProjectGetRouting) WithHeader(h map[string]string) func(*ProjectGetRoutingRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ProjectGetRouting) WithOpaqueID(s string) func(*ProjectGetRoutingRequest) {
	_ = "STUB: not implemented"
	return nil
}
