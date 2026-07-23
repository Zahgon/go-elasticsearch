package esapi

import (
	"context"
	"net/http"
)

func newWatcherActivateWatchFunc(t Transport) WatcherActivateWatch {
	_ = "STUB: not implemented"
	return *new(WatcherActivateWatch)
}

type WatcherActivateWatch func(watch_id string, o ...func(*WatcherActivateWatchRequest)) (*Response, error)

type WatcherActivateWatchRequest struct {
	WatchID string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r WatcherActivateWatchRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f WatcherActivateWatch) WithContext(v context.Context) func(*WatcherActivateWatchRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f WatcherActivateWatch) WithPretty() func(*WatcherActivateWatchRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f WatcherActivateWatch) WithHuman() func(*WatcherActivateWatchRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f WatcherActivateWatch) WithErrorTrace() func(*WatcherActivateWatchRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f WatcherActivateWatch) WithFilterPath(v ...string) func(*WatcherActivateWatchRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f WatcherActivateWatch) WithHeader(h map[string]string) func(*WatcherActivateWatchRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f WatcherActivateWatch) WithOpaqueID(s string) func(*WatcherActivateWatchRequest) {
	_ = "STUB: not implemented"
	return nil
}
