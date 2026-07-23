package esapi

import (
	"context"
	"net/http"
)

func newWatcherGetWatchFunc(t Transport) WatcherGetWatch {
	_ = "STUB: not implemented"
	return *new(WatcherGetWatch)
}

type WatcherGetWatch func(id string, o ...func(*WatcherGetWatchRequest)) (*Response, error)

type WatcherGetWatchRequest struct {
	WatchID string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r WatcherGetWatchRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f WatcherGetWatch) WithContext(v context.Context) func(*WatcherGetWatchRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f WatcherGetWatch) WithPretty() func(*WatcherGetWatchRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f WatcherGetWatch) WithHuman() func(*WatcherGetWatchRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f WatcherGetWatch) WithErrorTrace() func(*WatcherGetWatchRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f WatcherGetWatch) WithFilterPath(v ...string) func(*WatcherGetWatchRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f WatcherGetWatch) WithHeader(h map[string]string) func(*WatcherGetWatchRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f WatcherGetWatch) WithOpaqueID(s string) func(*WatcherGetWatchRequest) {
	_ = "STUB: not implemented"
	return nil
}
