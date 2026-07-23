package esapi

import (
	"context"
	"net/http"
)

func newProjectDeleteRoutingFunc(t Transport) ProjectDeleteRouting {
	_ = "STUB: not implemented"
	return *new(ProjectDeleteRouting)
}

type ProjectDeleteRouting func(name string, o ...func(*ProjectDeleteRoutingRequest)) (*Response, error)

type ProjectDeleteRoutingRequest struct {
	Name string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r ProjectDeleteRoutingRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f ProjectDeleteRouting) WithContext(v context.Context) func(*ProjectDeleteRoutingRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ProjectDeleteRouting) WithPretty() func(*ProjectDeleteRoutingRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ProjectDeleteRouting) WithHuman() func(*ProjectDeleteRoutingRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ProjectDeleteRouting) WithErrorTrace() func(*ProjectDeleteRoutingRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ProjectDeleteRouting) WithFilterPath(v ...string) func(*ProjectDeleteRoutingRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ProjectDeleteRouting) WithHeader(h map[string]string) func(*ProjectDeleteRoutingRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ProjectDeleteRouting) WithOpaqueID(s string) func(*ProjectDeleteRoutingRequest) {
	_ = "STUB: not implemented"
	return nil
}
