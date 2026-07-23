package esapi

import (
	"context"
	"net/http"
)

func newReindexRethrottleFunc(t Transport) ReindexRethrottle {
	_ = "STUB: not implemented"
	return *new(ReindexRethrottle)
}

type ReindexRethrottle func(task_id string, requests_per_second *int, o ...func(*ReindexRethrottleRequest)) (*Response, error)

type ReindexRethrottleRequest struct {
	TaskID string

	RequestsPerSecond *int

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r ReindexRethrottleRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f ReindexRethrottle) WithContext(v context.Context) func(*ReindexRethrottleRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ReindexRethrottle) WithRequestsPerSecond(v int) func(*ReindexRethrottleRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ReindexRethrottle) WithPretty() func(*ReindexRethrottleRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ReindexRethrottle) WithHuman() func(*ReindexRethrottleRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ReindexRethrottle) WithErrorTrace() func(*ReindexRethrottleRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ReindexRethrottle) WithFilterPath(v ...string) func(*ReindexRethrottleRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ReindexRethrottle) WithHeader(h map[string]string) func(*ReindexRethrottleRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ReindexRethrottle) WithOpaqueID(s string) func(*ReindexRethrottleRequest) {
	_ = "STUB: not implemented"
	return nil
}
