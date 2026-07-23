package esapi

import (
	"context"
	"io"
	"net/http"
)

func newProjectCreateRoutingFunc(t Transport) ProjectCreateRouting {
	_ = "STUB: not implemented"
	return *new(ProjectCreateRouting)
}

type ProjectCreateRouting func(name string, body io.Reader, o ...func(*ProjectCreateRoutingRequest)) (*Response, error)

type ProjectCreateRoutingRequest struct {
	Body io.Reader

	Name string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r ProjectCreateRoutingRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f ProjectCreateRouting) WithContext(v context.Context) func(*ProjectCreateRoutingRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ProjectCreateRouting) WithPretty() func(*ProjectCreateRoutingRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ProjectCreateRouting) WithHuman() func(*ProjectCreateRoutingRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ProjectCreateRouting) WithErrorTrace() func(*ProjectCreateRoutingRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ProjectCreateRouting) WithFilterPath(v ...string) func(*ProjectCreateRoutingRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ProjectCreateRouting) WithHeader(h map[string]string) func(*ProjectCreateRoutingRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ProjectCreateRouting) WithOpaqueID(s string) func(*ProjectCreateRoutingRequest) {
	_ = "STUB: not implemented"
	return nil
}
