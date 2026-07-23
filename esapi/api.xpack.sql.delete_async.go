package esapi

import (
	"context"
	"net/http"
)

func newSQLDeleteAsyncFunc(t Transport) SQLDeleteAsync {
	_ = "STUB: not implemented"
	return *new(SQLDeleteAsync)
}

type SQLDeleteAsync func(id string, o ...func(*SQLDeleteAsyncRequest)) (*Response, error)

type SQLDeleteAsyncRequest struct {
	DocumentID string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r SQLDeleteAsyncRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f SQLDeleteAsync) WithContext(v context.Context) func(*SQLDeleteAsyncRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SQLDeleteAsync) WithPretty() func(*SQLDeleteAsyncRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SQLDeleteAsync) WithHuman() func(*SQLDeleteAsyncRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SQLDeleteAsync) WithErrorTrace() func(*SQLDeleteAsyncRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SQLDeleteAsync) WithFilterPath(v ...string) func(*SQLDeleteAsyncRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SQLDeleteAsync) WithHeader(h map[string]string) func(*SQLDeleteAsyncRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SQLDeleteAsync) WithOpaqueID(s string) func(*SQLDeleteAsyncRequest) {
	_ = "STUB: not implemented"
	return nil
}
