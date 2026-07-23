package esapi

import (
	"context"
	"net/http"
	"time"
)

func newEqlGetFunc(t Transport) EqlGet { _ = "STUB: not implemented"; return *new(EqlGet) }

type EqlGet func(id string, o ...func(*EqlGetRequest)) (*Response, error)

type EqlGetRequest struct {
	DocumentID string

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

func (r EqlGetRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f EqlGet) WithContext(v context.Context) func(*EqlGetRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f EqlGet) WithKeepAlive(v time.Duration) func(*EqlGetRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f EqlGet) WithWaitForCompletionTimeout(v time.Duration) func(*EqlGetRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f EqlGet) WithPretty() func(*EqlGetRequest) { _ = "STUB: not implemented"; return nil }

func (f EqlGet) WithHuman() func(*EqlGetRequest) { _ = "STUB: not implemented"; return nil }

func (f EqlGet) WithErrorTrace() func(*EqlGetRequest) { _ = "STUB: not implemented"; return nil }

func (f EqlGet) WithFilterPath(v ...string) func(*EqlGetRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f EqlGet) WithHeader(h map[string]string) func(*EqlGetRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f EqlGet) WithOpaqueID(s string) func(*EqlGetRequest) { _ = "STUB: not implemented"; return nil }
