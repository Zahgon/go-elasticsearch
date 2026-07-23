package esapi

import (
	"context"
	"io"
	"net/http"
)

func newWatcherQueryWatchesFunc(t Transport) WatcherQueryWatches {
	_ = "STUB: not implemented"
	return *new(WatcherQueryWatches)
}

type WatcherQueryWatches func(o ...func(*WatcherQueryWatchesRequest)) (*Response, error)

type WatcherQueryWatchesRequest struct {
	Body io.Reader

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r WatcherQueryWatchesRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f WatcherQueryWatches) WithContext(v context.Context) func(*WatcherQueryWatchesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f WatcherQueryWatches) WithBody(v io.Reader) func(*WatcherQueryWatchesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f WatcherQueryWatches) WithPretty() func(*WatcherQueryWatchesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f WatcherQueryWatches) WithHuman() func(*WatcherQueryWatchesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f WatcherQueryWatches) WithErrorTrace() func(*WatcherQueryWatchesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f WatcherQueryWatches) WithFilterPath(v ...string) func(*WatcherQueryWatchesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f WatcherQueryWatches) WithHeader(h map[string]string) func(*WatcherQueryWatchesRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f WatcherQueryWatches) WithOpaqueID(s string) func(*WatcherQueryWatchesRequest) {
	_ = "STUB: not implemented"
	return nil
}
