package esapi

import (
	"context"
	"io"
	"net/http"
)

func newWatcherExecuteWatchFunc(t Transport) WatcherExecuteWatch {
	_ = "STUB: not implemented"
	return *new(WatcherExecuteWatch)
}

type WatcherExecuteWatch func(o ...func(*WatcherExecuteWatchRequest)) (*Response, error)

type WatcherExecuteWatchRequest struct {
	WatchID string

	Body io.Reader

	Debug *bool

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r WatcherExecuteWatchRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f WatcherExecuteWatch) WithContext(v context.Context) func(*WatcherExecuteWatchRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f WatcherExecuteWatch) WithBody(v io.Reader) func(*WatcherExecuteWatchRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f WatcherExecuteWatch) WithWatchID(v string) func(*WatcherExecuteWatchRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f WatcherExecuteWatch) WithDebug(v bool) func(*WatcherExecuteWatchRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f WatcherExecuteWatch) WithPretty() func(*WatcherExecuteWatchRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f WatcherExecuteWatch) WithHuman() func(*WatcherExecuteWatchRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f WatcherExecuteWatch) WithErrorTrace() func(*WatcherExecuteWatchRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f WatcherExecuteWatch) WithFilterPath(v ...string) func(*WatcherExecuteWatchRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f WatcherExecuteWatch) WithHeader(h map[string]string) func(*WatcherExecuteWatchRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f WatcherExecuteWatch) WithOpaqueID(s string) func(*WatcherExecuteWatchRequest) {
	_ = "STUB: not implemented"
	return nil
}
