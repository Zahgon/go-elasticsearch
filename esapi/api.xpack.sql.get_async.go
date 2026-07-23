package esapi

import (
	"context"
	"net/http"
	"time"
)

func newSQLGetAsyncFunc(t Transport) SQLGetAsync {
	_ = "STUB: not implemented"
	return *new(SQLGetAsync)
}

type SQLGetAsync func(id string, o ...func(*SQLGetAsyncRequest)) (*Response, error)

type SQLGetAsyncRequest struct {
	DocumentID string

	Delimiter                string
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

func (r SQLGetAsyncRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f SQLGetAsync) WithContext(v context.Context) func(*SQLGetAsyncRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SQLGetAsync) WithDelimiter(v string) func(*SQLGetAsyncRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SQLGetAsync) WithFormat(v string) func(*SQLGetAsyncRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SQLGetAsync) WithKeepAlive(v time.Duration) func(*SQLGetAsyncRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SQLGetAsync) WithWaitForCompletionTimeout(v time.Duration) func(*SQLGetAsyncRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SQLGetAsync) WithPretty() func(*SQLGetAsyncRequest) { _ = "STUB: not implemented"; return nil }

func (f SQLGetAsync) WithHuman() func(*SQLGetAsyncRequest) { _ = "STUB: not implemented"; return nil }

func (f SQLGetAsync) WithErrorTrace() func(*SQLGetAsyncRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SQLGetAsync) WithFilterPath(v ...string) func(*SQLGetAsyncRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SQLGetAsync) WithHeader(h map[string]string) func(*SQLGetAsyncRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SQLGetAsync) WithOpaqueID(s string) func(*SQLGetAsyncRequest) {
	_ = "STUB: not implemented"
	return nil
}
