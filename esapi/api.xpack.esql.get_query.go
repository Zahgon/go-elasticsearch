package esapi

import (
	"context"
	"net/http"
)

func newEsqlGetQueryFunc(t Transport) EsqlGetQuery {
	_ = "STUB: not implemented"
	return *new(EsqlGetQuery)
}

type EsqlGetQuery func(id string, o ...func(*EsqlGetQueryRequest)) (*Response, error)

type EsqlGetQueryRequest struct {
	DocumentID string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r EsqlGetQueryRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f EsqlGetQuery) WithContext(v context.Context) func(*EsqlGetQueryRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f EsqlGetQuery) WithPretty() func(*EsqlGetQueryRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f EsqlGetQuery) WithHuman() func(*EsqlGetQueryRequest) { _ = "STUB: not implemented"; return nil }

func (f EsqlGetQuery) WithErrorTrace() func(*EsqlGetQueryRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f EsqlGetQuery) WithFilterPath(v ...string) func(*EsqlGetQueryRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f EsqlGetQuery) WithHeader(h map[string]string) func(*EsqlGetQueryRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f EsqlGetQuery) WithOpaqueID(s string) func(*EsqlGetQueryRequest) {
	_ = "STUB: not implemented"
	return nil
}
