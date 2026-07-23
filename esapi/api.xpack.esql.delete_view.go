package esapi

import (
	"context"
	"net/http"
)

func newEsqlDeleteViewFunc(t Transport) EsqlDeleteView {
	_ = "STUB: not implemented"
	return *new(EsqlDeleteView)
}

type EsqlDeleteView func(name string, o ...func(*EsqlDeleteViewRequest)) (*Response, error)

type EsqlDeleteViewRequest struct {
	Name string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r EsqlDeleteViewRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f EsqlDeleteView) WithContext(v context.Context) func(*EsqlDeleteViewRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f EsqlDeleteView) WithPretty() func(*EsqlDeleteViewRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f EsqlDeleteView) WithHuman() func(*EsqlDeleteViewRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f EsqlDeleteView) WithErrorTrace() func(*EsqlDeleteViewRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f EsqlDeleteView) WithFilterPath(v ...string) func(*EsqlDeleteViewRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f EsqlDeleteView) WithHeader(h map[string]string) func(*EsqlDeleteViewRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f EsqlDeleteView) WithOpaqueID(s string) func(*EsqlDeleteViewRequest) {
	_ = "STUB: not implemented"
	return nil
}
