package esapi

import (
	"context"
	"net/http"
)

func newAsyncSearchDeleteFunc(t Transport) AsyncSearchDelete {
	_ = "STUB: not implemented"
	return *new(AsyncSearchDelete)
}

type AsyncSearchDelete func(id string, o ...func(*AsyncSearchDeleteRequest)) (*Response, error)

type AsyncSearchDeleteRequest struct {
	DocumentID string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r AsyncSearchDeleteRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f AsyncSearchDelete) WithContext(v context.Context) func(*AsyncSearchDeleteRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f AsyncSearchDelete) WithPretty() func(*AsyncSearchDeleteRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f AsyncSearchDelete) WithHuman() func(*AsyncSearchDeleteRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f AsyncSearchDelete) WithErrorTrace() func(*AsyncSearchDeleteRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f AsyncSearchDelete) WithFilterPath(v ...string) func(*AsyncSearchDeleteRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f AsyncSearchDelete) WithHeader(h map[string]string) func(*AsyncSearchDeleteRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f AsyncSearchDelete) WithOpaqueID(s string) func(*AsyncSearchDeleteRequest) {
	_ = "STUB: not implemented"
	return nil
}
