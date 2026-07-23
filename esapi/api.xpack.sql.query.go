package esapi

import (
	"context"
	"io"
	"net/http"
)

func newSQLQueryFunc(t Transport) SQLQuery { _ = "STUB: not implemented"; return *new(SQLQuery) }

type SQLQuery func(body io.Reader, o ...func(*SQLQueryRequest)) (*Response, error)

type SQLQueryRequest struct {
	Body io.Reader

	Format string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r SQLQueryRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f SQLQuery) WithContext(v context.Context) func(*SQLQueryRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SQLQuery) WithFormat(v string) func(*SQLQueryRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SQLQuery) WithPretty() func(*SQLQueryRequest) { _ = "STUB: not implemented"; return nil }

func (f SQLQuery) WithHuman() func(*SQLQueryRequest) { _ = "STUB: not implemented"; return nil }

func (f SQLQuery) WithErrorTrace() func(*SQLQueryRequest) { _ = "STUB: not implemented"; return nil }

func (f SQLQuery) WithFilterPath(v ...string) func(*SQLQueryRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SQLQuery) WithHeader(h map[string]string) func(*SQLQueryRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SQLQuery) WithOpaqueID(s string) func(*SQLQueryRequest) {
	_ = "STUB: not implemented"
	return nil
}
