package esapi

import (
	"context"
	"net/http"
)

func newWatcherStatsFunc(t Transport) WatcherStats {
	_ = "STUB: not implemented"
	return *new(WatcherStats)
}

type WatcherStats func(o ...func(*WatcherStatsRequest)) (*Response, error)

type WatcherStatsRequest struct {
	Metric []string

	EmitStacktraces *bool

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r WatcherStatsRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f WatcherStats) WithContext(v context.Context) func(*WatcherStatsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f WatcherStats) WithMetric(v ...string) func(*WatcherStatsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f WatcherStats) WithEmitStacktraces(v bool) func(*WatcherStatsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f WatcherStats) WithPretty() func(*WatcherStatsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f WatcherStats) WithHuman() func(*WatcherStatsRequest) { _ = "STUB: not implemented"; return nil }

func (f WatcherStats) WithErrorTrace() func(*WatcherStatsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f WatcherStats) WithFilterPath(v ...string) func(*WatcherStatsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f WatcherStats) WithHeader(h map[string]string) func(*WatcherStatsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f WatcherStats) WithOpaqueID(s string) func(*WatcherStatsRequest) {
	_ = "STUB: not implemented"
	return nil
}
