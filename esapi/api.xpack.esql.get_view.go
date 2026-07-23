package esapi

import (
	"context"
	"net/http"
)

func newEsqlGetViewFunc(t Transport) EsqlGetView {
	_ = "STUB: not implemented"
	return *new(EsqlGetView)
}

type EsqlGetView func(o ...func(*EsqlGetViewRequest)) (*Response, error)

type EsqlGetViewRequest struct {
	Name []string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r EsqlGetViewRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f EsqlGetView) WithContext(v context.Context) func(*EsqlGetViewRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f EsqlGetView) WithName(v ...string) func(*EsqlGetViewRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f EsqlGetView) WithPretty() func(*EsqlGetViewRequest) { _ = "STUB: not implemented"; return nil }

func (f EsqlGetView) WithHuman() func(*EsqlGetViewRequest) { _ = "STUB: not implemented"; return nil }

func (f EsqlGetView) WithErrorTrace() func(*EsqlGetViewRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f EsqlGetView) WithFilterPath(v ...string) func(*EsqlGetViewRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f EsqlGetView) WithHeader(h map[string]string) func(*EsqlGetViewRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f EsqlGetView) WithOpaqueID(s string) func(*EsqlGetViewRequest) {
	_ = "STUB: not implemented"
	return nil
}
