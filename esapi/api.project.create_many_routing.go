package esapi

import (
	"context"
	"io"
	"net/http"
)

func newProjectCreateManyRoutingFunc(t Transport) ProjectCreateManyRouting {
	_ = "STUB: not implemented"
	return *new(ProjectCreateManyRouting)
}

type ProjectCreateManyRouting func(body io.Reader, o ...func(*ProjectCreateManyRoutingRequest)) (*Response, error)

type ProjectCreateManyRoutingRequest struct {
	Body io.Reader

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r ProjectCreateManyRoutingRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f ProjectCreateManyRouting) WithContext(v context.Context) func(*ProjectCreateManyRoutingRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ProjectCreateManyRouting) WithPretty() func(*ProjectCreateManyRoutingRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ProjectCreateManyRouting) WithHuman() func(*ProjectCreateManyRoutingRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ProjectCreateManyRouting) WithErrorTrace() func(*ProjectCreateManyRoutingRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ProjectCreateManyRouting) WithFilterPath(v ...string) func(*ProjectCreateManyRoutingRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ProjectCreateManyRouting) WithHeader(h map[string]string) func(*ProjectCreateManyRoutingRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ProjectCreateManyRouting) WithOpaqueID(s string) func(*ProjectCreateManyRoutingRequest) {
	_ = "STUB: not implemented"
	return nil
}
