package esapi

import (
	"context"
	"net/http"
)

func newEqlDeleteFunc(t Transport) EqlDelete { _ = "STUB: not implemented"; return *new(EqlDelete) }

type EqlDelete func(id string, o ...func(*EqlDeleteRequest)) (*Response, error)

type EqlDeleteRequest struct {
	DocumentID string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r EqlDeleteRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f EqlDelete) WithContext(v context.Context) func(*EqlDeleteRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f EqlDelete) WithPretty() func(*EqlDeleteRequest) { _ = "STUB: not implemented"; return nil }

func (f EqlDelete) WithHuman() func(*EqlDeleteRequest) { _ = "STUB: not implemented"; return nil }

func (f EqlDelete) WithErrorTrace() func(*EqlDeleteRequest) { _ = "STUB: not implemented"; return nil }

func (f EqlDelete) WithFilterPath(v ...string) func(*EqlDeleteRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f EqlDelete) WithHeader(h map[string]string) func(*EqlDeleteRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f EqlDelete) WithOpaqueID(s string) func(*EqlDeleteRequest) {
	_ = "STUB: not implemented"
	return nil
}
