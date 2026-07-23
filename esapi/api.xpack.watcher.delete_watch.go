package esapi

import (
	"context"
	"net/http"
)

func newWatcherDeleteWatchFunc(t Transport) WatcherDeleteWatch {
	_ = "STUB: not implemented"
	return *new(WatcherDeleteWatch)
}

type WatcherDeleteWatch func(id string, o ...func(*WatcherDeleteWatchRequest)) (*Response, error)

type WatcherDeleteWatchRequest struct {
	WatchID string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r WatcherDeleteWatchRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f WatcherDeleteWatch) WithContext(v context.Context) func(*WatcherDeleteWatchRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f WatcherDeleteWatch) WithPretty() func(*WatcherDeleteWatchRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f WatcherDeleteWatch) WithHuman() func(*WatcherDeleteWatchRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f WatcherDeleteWatch) WithErrorTrace() func(*WatcherDeleteWatchRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f WatcherDeleteWatch) WithFilterPath(v ...string) func(*WatcherDeleteWatchRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f WatcherDeleteWatch) WithHeader(h map[string]string) func(*WatcherDeleteWatchRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f WatcherDeleteWatch) WithOpaqueID(s string) func(*WatcherDeleteWatchRequest) {
	_ = "STUB: not implemented"
	return nil
}
