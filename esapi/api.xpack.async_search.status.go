package esapi

import (
	"context"
	"net/http"
	"time"
)

func newAsyncSearchStatusFunc(t Transport) AsyncSearchStatus {
	_ = "STUB: not implemented"
	return *new(AsyncSearchStatus)
}

type AsyncSearchStatus func(id string, o ...func(*AsyncSearchStatusRequest)) (*Response, error)

type AsyncSearchStatusRequest struct {
	DocumentID string

	KeepAlive time.Duration

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r AsyncSearchStatusRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f AsyncSearchStatus) WithContext(v context.Context) func(*AsyncSearchStatusRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f AsyncSearchStatus) WithKeepAlive(v time.Duration) func(*AsyncSearchStatusRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f AsyncSearchStatus) WithPretty() func(*AsyncSearchStatusRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f AsyncSearchStatus) WithHuman() func(*AsyncSearchStatusRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f AsyncSearchStatus) WithErrorTrace() func(*AsyncSearchStatusRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f AsyncSearchStatus) WithFilterPath(v ...string) func(*AsyncSearchStatusRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f AsyncSearchStatus) WithHeader(h map[string]string) func(*AsyncSearchStatusRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f AsyncSearchStatus) WithOpaqueID(s string) func(*AsyncSearchStatusRequest) {
	_ = "STUB: not implemented"
	return nil
}
