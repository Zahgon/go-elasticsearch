package esapi

import (
	"context"
	"net/http"
	"time"
)

func newWatcherStartFunc(t Transport) WatcherStart {
	_ = "STUB: not implemented"
	return *new(WatcherStart)
}

type WatcherStart func(o ...func(*WatcherStartRequest)) (*Response, error)

type WatcherStartRequest struct {
	MasterTimeout time.Duration

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r WatcherStartRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f WatcherStart) WithContext(v context.Context) func(*WatcherStartRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f WatcherStart) WithMasterTimeout(v time.Duration) func(*WatcherStartRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f WatcherStart) WithPretty() func(*WatcherStartRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f WatcherStart) WithHuman() func(*WatcherStartRequest) { _ = "STUB: not implemented"; return nil }

func (f WatcherStart) WithErrorTrace() func(*WatcherStartRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f WatcherStart) WithFilterPath(v ...string) func(*WatcherStartRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f WatcherStart) WithHeader(h map[string]string) func(*WatcherStartRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f WatcherStart) WithOpaqueID(s string) func(*WatcherStartRequest) {
	_ = "STUB: not implemented"
	return nil
}
