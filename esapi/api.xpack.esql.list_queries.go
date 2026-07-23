package esapi

import (
	"context"
	"net/http"
)

func newEsqlListQueriesFunc(t Transport) EsqlListQueries {
	_ = "STUB: not implemented"
	return *new(EsqlListQueries)
}

type EsqlListQueries func(o ...func(*EsqlListQueriesRequest)) (*Response, error)

type EsqlListQueriesRequest struct {
	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r EsqlListQueriesRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f EsqlListQueries) WithContext(v context.Context) func(*EsqlListQueriesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f EsqlListQueries) WithPretty() func(*EsqlListQueriesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f EsqlListQueries) WithHuman() func(*EsqlListQueriesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f EsqlListQueries) WithErrorTrace() func(*EsqlListQueriesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f EsqlListQueries) WithFilterPath(v ...string) func(*EsqlListQueriesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f EsqlListQueries) WithHeader(h map[string]string) func(*EsqlListQueriesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f EsqlListQueries) WithOpaqueID(s string) func(*EsqlListQueriesRequest) {
	_ = "STUB: not implemented"
	return nil
}
