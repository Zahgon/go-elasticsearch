package esapi

import (
	"context"
	"io"
	"net/http"
)

func newEsqlPutViewFunc(t Transport) EsqlPutView {
	_ = "STUB: not implemented"
	return *new(EsqlPutView)
}

type EsqlPutView func(name string, body io.Reader, o ...func(*EsqlPutViewRequest)) (*Response, error)

type EsqlPutViewRequest struct {
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

func (r EsqlPutViewRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f EsqlPutView) WithContext(v context.Context) func(*EsqlPutViewRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f EsqlPutView) WithPretty() func(*EsqlPutViewRequest) { _ = "STUB: not implemented"; return nil }

func (f EsqlPutView) WithHuman() func(*EsqlPutViewRequest) { _ = "STUB: not implemented"; return nil }

func (f EsqlPutView) WithErrorTrace() func(*EsqlPutViewRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f EsqlPutView) WithFilterPath(v ...string) func(*EsqlPutViewRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f EsqlPutView) WithHeader(h map[string]string) func(*EsqlPutViewRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f EsqlPutView) WithOpaqueID(s string) func(*EsqlPutViewRequest) {
	_ = "STUB: not implemented"
	return nil
}
