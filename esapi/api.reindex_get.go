package esapi

import (
	"context"
	"net/http"
	"time"
)

func newReindexGetFunc(t Transport) ReindexGet { _ = "STUB: not implemented"; return *new(ReindexGet) }

type ReindexGet func(task_id string, o ...func(*ReindexGetRequest)) (*Response, error)

type ReindexGetRequest struct {
	TaskID string

	Timeout           time.Duration
	WaitForCompletion *bool

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r ReindexGetRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f ReindexGet) WithContext(v context.Context) func(*ReindexGetRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ReindexGet) WithTimeout(v time.Duration) func(*ReindexGetRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ReindexGet) WithWaitForCompletion(v bool) func(*ReindexGetRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ReindexGet) WithPretty() func(*ReindexGetRequest) { _ = "STUB: not implemented"; return nil }

func (f ReindexGet) WithHuman() func(*ReindexGetRequest) { _ = "STUB: not implemented"; return nil }

func (f ReindexGet) WithErrorTrace() func(*ReindexGetRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ReindexGet) WithFilterPath(v ...string) func(*ReindexGetRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ReindexGet) WithHeader(h map[string]string) func(*ReindexGetRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ReindexGet) WithOpaqueID(s string) func(*ReindexGetRequest) {
	_ = "STUB: not implemented"
	return nil
}
