package esapi

import (
	"context"
	"net/http"
	"time"
)

func newAsyncSearchGetFunc(t Transport) AsyncSearchGet {
	_ = "STUB: not implemented"
	return *new(AsyncSearchGet)
}

type AsyncSearchGet func(id string, o ...func(*AsyncSearchGetRequest)) (*Response, error)

type AsyncSearchGetRequest struct {
	DocumentID string

	KeepAlive                 time.Duration
	ReturnIntermediateResults *bool
	TypedKeys                 *bool
	WaitForCompletionTimeout  time.Duration

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r AsyncSearchGetRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f AsyncSearchGet) WithContext(v context.Context) func(*AsyncSearchGetRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f AsyncSearchGet) WithKeepAlive(v time.Duration) func(*AsyncSearchGetRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f AsyncSearchGet) WithReturnIntermediateResults(v bool) func(*AsyncSearchGetRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f AsyncSearchGet) WithTypedKeys(v bool) func(*AsyncSearchGetRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f AsyncSearchGet) WithWaitForCompletionTimeout(v time.Duration) func(*AsyncSearchGetRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f AsyncSearchGet) WithPretty() func(*AsyncSearchGetRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f AsyncSearchGet) WithHuman() func(*AsyncSearchGetRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f AsyncSearchGet) WithErrorTrace() func(*AsyncSearchGetRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f AsyncSearchGet) WithFilterPath(v ...string) func(*AsyncSearchGetRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f AsyncSearchGet) WithHeader(h map[string]string) func(*AsyncSearchGetRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f AsyncSearchGet) WithOpaqueID(s string) func(*AsyncSearchGetRequest) {
	_ = "STUB: not implemented"
	return nil
}
