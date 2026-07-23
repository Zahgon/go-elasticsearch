package esapi

import (
	"context"
	"net/http"
)

func newEsqlAsyncQueryStopFunc(t Transport) EsqlAsyncQueryStop {
	_ = "STUB: not implemented"
	return *new(EsqlAsyncQueryStop)
}

type EsqlAsyncQueryStop func(id string, o ...func(*EsqlAsyncQueryStopRequest)) (*Response, error)

type EsqlAsyncQueryStopRequest struct {
	DocumentID string

	DropNullColumns *bool

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r EsqlAsyncQueryStopRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f EsqlAsyncQueryStop) WithContext(v context.Context) func(*EsqlAsyncQueryStopRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f EsqlAsyncQueryStop) WithDropNullColumns(v bool) func(*EsqlAsyncQueryStopRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f EsqlAsyncQueryStop) WithPretty() func(*EsqlAsyncQueryStopRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f EsqlAsyncQueryStop) WithHuman() func(*EsqlAsyncQueryStopRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f EsqlAsyncQueryStop) WithErrorTrace() func(*EsqlAsyncQueryStopRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f EsqlAsyncQueryStop) WithFilterPath(v ...string) func(*EsqlAsyncQueryStopRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f EsqlAsyncQueryStop) WithHeader(h map[string]string) func(*EsqlAsyncQueryStopRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f EsqlAsyncQueryStop) WithOpaqueID(s string) func(*EsqlAsyncQueryStopRequest) {
	_ = "STUB: not implemented"
	return nil
}
