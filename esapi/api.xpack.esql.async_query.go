package esapi

import (
	"context"
	"io"
	"net/http"
)

func newEsqlAsyncQueryFunc(t Transport) EsqlAsyncQuery {
	_ = "STUB: not implemented"
	return *new(EsqlAsyncQuery)
}

type EsqlAsyncQuery func(body io.Reader, o ...func(*EsqlAsyncQueryRequest)) (*Response, error)

type EsqlAsyncQueryRequest struct {
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

func (r EsqlAsyncQueryRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f EsqlAsyncQuery) WithContext(v context.Context) func(*EsqlAsyncQueryRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f EsqlAsyncQuery) WithAllowPartialResults(v bool) func(*EsqlAsyncQueryRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f EsqlAsyncQuery) WithDelimiter(v string) func(*EsqlAsyncQueryRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f EsqlAsyncQuery) WithDropNullColumns(v bool) func(*EsqlAsyncQueryRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f EsqlAsyncQuery) WithFormat(v string) func(*EsqlAsyncQueryRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f EsqlAsyncQuery) WithPretty() func(*EsqlAsyncQueryRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f EsqlAsyncQuery) WithHuman() func(*EsqlAsyncQueryRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f EsqlAsyncQuery) WithErrorTrace() func(*EsqlAsyncQueryRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f EsqlAsyncQuery) WithFilterPath(v ...string) func(*EsqlAsyncQueryRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f EsqlAsyncQuery) WithHeader(h map[string]string) func(*EsqlAsyncQueryRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f EsqlAsyncQuery) WithOpaqueID(s string) func(*EsqlAsyncQueryRequest) {
	_ = "STUB: not implemented"
	return nil
}
