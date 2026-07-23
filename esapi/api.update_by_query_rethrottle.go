package esapi

import (
	"context"
	"net/http"
)

func newUpdateByQueryRethrottleFunc(t Transport) UpdateByQueryRethrottle {
	_ = "STUB: not implemented"
	return *new(UpdateByQueryRethrottle)
}

type UpdateByQueryRethrottle func(task_id string, requests_per_second *int, o ...func(*UpdateByQueryRethrottleRequest)) (*Response, error)

type UpdateByQueryRethrottleRequest struct {
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

func (r UpdateByQueryRethrottleRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f UpdateByQueryRethrottle) WithContext(v context.Context) func(*UpdateByQueryRethrottleRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f UpdateByQueryRethrottle) WithRequestsPerSecond(v int) func(*UpdateByQueryRethrottleRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f UpdateByQueryRethrottle) WithPretty() func(*UpdateByQueryRethrottleRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f UpdateByQueryRethrottle) WithHuman() func(*UpdateByQueryRethrottleRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f UpdateByQueryRethrottle) WithErrorTrace() func(*UpdateByQueryRethrottleRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f UpdateByQueryRethrottle) WithFilterPath(v ...string) func(*UpdateByQueryRethrottleRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f UpdateByQueryRethrottle) WithHeader(h map[string]string) func(*UpdateByQueryRethrottleRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f UpdateByQueryRethrottle) WithOpaqueID(s string) func(*UpdateByQueryRethrottleRequest) {
	_ = "STUB: not implemented"
	return nil
}
