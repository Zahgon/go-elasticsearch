package esapi

import (
	"context"
	"net/http"
	"time"
)

func newWatcherStopFunc(t Transport) WatcherStop {
	_ = "STUB: not implemented"
	return *new(WatcherStop)
}

type WatcherStop func(o ...func(*WatcherStopRequest)) (*Response, error)

type WatcherStopRequest struct {
	MasterTimeout time.Duration

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r WatcherStopRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f WatcherStop) WithContext(v context.Context) func(*WatcherStopRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f WatcherStop) WithMasterTimeout(v time.Duration) func(*WatcherStopRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f WatcherStop) WithPretty() func(*WatcherStopRequest) { _ = "STUB: not implemented"; return nil }

func (f WatcherStop) WithHuman() func(*WatcherStopRequest) { _ = "STUB: not implemented"; return nil }

func (f WatcherStop) WithErrorTrace() func(*WatcherStopRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f WatcherStop) WithFilterPath(v ...string) func(*WatcherStopRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f WatcherStop) WithHeader(h map[string]string) func(*WatcherStopRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f WatcherStop) WithOpaqueID(s string) func(*WatcherStopRequest) {
	_ = "STUB: not implemented"
	return nil
}
