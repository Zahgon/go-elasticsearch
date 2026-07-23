package esapi

import (
	"context"
	"net/http"
	"time"
)

func newEsqlAsyncQueryGetFunc(t Transport) EsqlAsyncQueryGet {
	_ = "STUB: not implemented"
	return *new(EsqlAsyncQueryGet)
}

type EsqlAsyncQueryGet func(id string, o ...func(*EsqlAsyncQueryGetRequest)) (*Response, error)

type EsqlAsyncQueryGetRequest struct {
	DocumentID string

	DropNullColumns          *bool
	Format                   string
	KeepAlive                time.Duration
	WaitForCompletionTimeout time.Duration

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r EsqlAsyncQueryGetRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f EsqlAsyncQueryGet) WithContext(v context.Context) func(*EsqlAsyncQueryGetRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f EsqlAsyncQueryGet) WithDropNullColumns(v bool) func(*EsqlAsyncQueryGetRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f EsqlAsyncQueryGet) WithFormat(v string) func(*EsqlAsyncQueryGetRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f EsqlAsyncQueryGet) WithKeepAlive(v time.Duration) func(*EsqlAsyncQueryGetRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f EsqlAsyncQueryGet) WithWaitForCompletionTimeout(v time.Duration) func(*EsqlAsyncQueryGetRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f EsqlAsyncQueryGet) WithPretty() func(*EsqlAsyncQueryGetRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f EsqlAsyncQueryGet) WithHuman() func(*EsqlAsyncQueryGetRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f EsqlAsyncQueryGet) WithErrorTrace() func(*EsqlAsyncQueryGetRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f EsqlAsyncQueryGet) WithFilterPath(v ...string) func(*EsqlAsyncQueryGetRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f EsqlAsyncQueryGet) WithHeader(h map[string]string) func(*EsqlAsyncQueryGetRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f EsqlAsyncQueryGet) WithOpaqueID(s string) func(*EsqlAsyncQueryGetRequest) {
	_ = "STUB: not implemented"
	return nil
}
