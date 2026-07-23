package esapi

import (
	"context"
	"net/http"
)

func newWatcherAckWatchFunc(t Transport) WatcherAckWatch {
	_ = "STUB: not implemented"
	return *new(WatcherAckWatch)
}

type WatcherAckWatch func(watch_id string, o ...func(*WatcherAckWatchRequest)) (*Response, error)

type WatcherAckWatchRequest struct {
	ActionID []string
	WatchID  string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r WatcherAckWatchRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f WatcherAckWatch) WithContext(v context.Context) func(*WatcherAckWatchRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f WatcherAckWatch) WithActionID(v ...string) func(*WatcherAckWatchRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f WatcherAckWatch) WithPretty() func(*WatcherAckWatchRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f WatcherAckWatch) WithHuman() func(*WatcherAckWatchRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f WatcherAckWatch) WithErrorTrace() func(*WatcherAckWatchRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f WatcherAckWatch) WithFilterPath(v ...string) func(*WatcherAckWatchRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f WatcherAckWatch) WithHeader(h map[string]string) func(*WatcherAckWatchRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f WatcherAckWatch) WithOpaqueID(s string) func(*WatcherAckWatchRequest) {
	_ = "STUB: not implemented"
	return nil
}
