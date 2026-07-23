package esapi

import (
	"context"
	"net/http"
)

func newWatcherDeactivateWatchFunc(t Transport) WatcherDeactivateWatch {
	_ = "STUB: not implemented"
	return *new(WatcherDeactivateWatch)
}

type WatcherDeactivateWatch func(watch_id string, o ...func(*WatcherDeactivateWatchRequest)) (*Response, error)

type WatcherDeactivateWatchRequest struct {
	WatchID string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r WatcherDeactivateWatchRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f WatcherDeactivateWatch) WithContext(v context.Context) func(*WatcherDeactivateWatchRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f WatcherDeactivateWatch) WithPretty() func(*WatcherDeactivateWatchRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f WatcherDeactivateWatch) WithHuman() func(*WatcherDeactivateWatchRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f WatcherDeactivateWatch) WithErrorTrace() func(*WatcherDeactivateWatchRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f WatcherDeactivateWatch) WithFilterPath(v ...string) func(*WatcherDeactivateWatchRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f WatcherDeactivateWatch) WithHeader(h map[string]string) func(*WatcherDeactivateWatchRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f WatcherDeactivateWatch) WithOpaqueID(s string) func(*WatcherDeactivateWatchRequest) {
	_ = "STUB: not implemented"
	return nil
}
