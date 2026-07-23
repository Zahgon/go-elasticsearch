package esapi

import (
	"context"
	"io"
	"net/http"
)

func newEsqlQueryFunc(t Transport) EsqlQuery { _ = "STUB: not implemented"; return *new(EsqlQuery) }

type EsqlQuery func(body io.Reader, o ...func(*EsqlQueryRequest)) (*Response, error)

type EsqlQueryRequest struct {
	Body io.Reader

	AllowPartialResults *bool
	Delimiter           string
	DropNullColumns     *bool
	Format              string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r EsqlQueryRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f EsqlQuery) WithContext(v context.Context) func(*EsqlQueryRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f EsqlQuery) WithAllowPartialResults(v bool) func(*EsqlQueryRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f EsqlQuery) WithDelimiter(v string) func(*EsqlQueryRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f EsqlQuery) WithDropNullColumns(v bool) func(*EsqlQueryRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f EsqlQuery) WithFormat(v string) func(*EsqlQueryRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f EsqlQuery) WithPretty() func(*EsqlQueryRequest) { _ = "STUB: not implemented"; return nil }

func (f EsqlQuery) WithHuman() func(*EsqlQueryRequest) { _ = "STUB: not implemented"; return nil }

func (f EsqlQuery) WithErrorTrace() func(*EsqlQueryRequest) { _ = "STUB: not implemented"; return nil }

func (f EsqlQuery) WithFilterPath(v ...string) func(*EsqlQueryRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f EsqlQuery) WithHeader(h map[string]string) func(*EsqlQueryRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f EsqlQuery) WithOpaqueID(s string) func(*EsqlQueryRequest) {
	_ = "STUB: not implemented"
	return nil
}
