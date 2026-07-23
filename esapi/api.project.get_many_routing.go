package esapi

import (
	"context"
	"net/http"
)

func newProjectGetManyRoutingFunc(t Transport) ProjectGetManyRouting {
	_ = "STUB: not implemented"
	return *new(ProjectGetManyRouting)
}

type ProjectGetManyRouting func(o ...func(*ProjectGetManyRoutingRequest)) (*Response, error)

type ProjectGetManyRoutingRequest struct {
	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r ProjectGetManyRoutingRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f ProjectGetManyRouting) WithContext(v context.Context) func(*ProjectGetManyRoutingRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ProjectGetManyRouting) WithPretty() func(*ProjectGetManyRoutingRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ProjectGetManyRouting) WithHuman() func(*ProjectGetManyRoutingRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ProjectGetManyRouting) WithErrorTrace() func(*ProjectGetManyRoutingRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ProjectGetManyRouting) WithFilterPath(v ...string) func(*ProjectGetManyRoutingRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ProjectGetManyRouting) WithHeader(h map[string]string) func(*ProjectGetManyRoutingRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ProjectGetManyRouting) WithOpaqueID(s string) func(*ProjectGetManyRoutingRequest) {
	_ = "STUB: not implemented"
	return nil
}
