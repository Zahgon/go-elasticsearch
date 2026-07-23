package esapi

import (
	"context"
	"net/http"
)

func newDeleteByQueryRethrottleFunc(t Transport) DeleteByQueryRethrottle {
	_ = "STUB: not implemented"
	return *new(DeleteByQueryRethrottle)
}

type DeleteByQueryRethrottle func(task_id string, requests_per_second *int, o ...func(*DeleteByQueryRethrottleRequest)) (*Response, error)

type DeleteByQueryRethrottleRequest struct {
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

func (r DeleteByQueryRethrottleRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f DeleteByQueryRethrottle) WithContext(v context.Context) func(*DeleteByQueryRethrottleRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f DeleteByQueryRethrottle) WithRequestsPerSecond(v int) func(*DeleteByQueryRethrottleRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f DeleteByQueryRethrottle) WithPretty() func(*DeleteByQueryRethrottleRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f DeleteByQueryRethrottle) WithHuman() func(*DeleteByQueryRethrottleRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f DeleteByQueryRethrottle) WithErrorTrace() func(*DeleteByQueryRethrottleRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f DeleteByQueryRethrottle) WithFilterPath(v ...string) func(*DeleteByQueryRethrottleRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f DeleteByQueryRethrottle) WithHeader(h map[string]string) func(*DeleteByQueryRethrottleRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f DeleteByQueryRethrottle) WithOpaqueID(s string) func(*DeleteByQueryRethrottleRequest) {
	_ = "STUB: not implemented"
	return nil
}
