package esapi

import (
	"context"
	"net/http"
)

func newEsqlAsyncQueryDeleteFunc(t Transport) EsqlAsyncQueryDelete {
	_ = "STUB: not implemented"
	return *new(EsqlAsyncQueryDelete)
}

type EsqlAsyncQueryDelete func(id string, o ...func(*EsqlAsyncQueryDeleteRequest)) (*Response, error)

type EsqlAsyncQueryDeleteRequest struct {
	DocumentID string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r EsqlAsyncQueryDeleteRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f EsqlAsyncQueryDelete) WithContext(v context.Context) func(*EsqlAsyncQueryDeleteRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f EsqlAsyncQueryDelete) WithPretty() func(*EsqlAsyncQueryDeleteRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f EsqlAsyncQueryDelete) WithHuman() func(*EsqlAsyncQueryDeleteRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f EsqlAsyncQueryDelete) WithErrorTrace() func(*EsqlAsyncQueryDeleteRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f EsqlAsyncQueryDelete) WithFilterPath(v ...string) func(*EsqlAsyncQueryDeleteRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f EsqlAsyncQueryDelete) WithHeader(h map[string]string) func(*EsqlAsyncQueryDeleteRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f EsqlAsyncQueryDelete) WithOpaqueID(s string) func(*EsqlAsyncQueryDeleteRequest) {
	_ = "STUB: not implemented"
	return nil
}
