package esapi

import (
	"context"
	"net/http"
)

func newIndicesGetDataLifecycleStatsFunc(t Transport) IndicesGetDataLifecycleStats {
	_ = "STUB: not implemented"
	return *new(IndicesGetDataLifecycleStats)
}

type IndicesGetDataLifecycleStats func(o ...func(*IndicesGetDataLifecycleStatsRequest)) (*Response, error)

type IndicesGetDataLifecycleStatsRequest struct {
	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r IndicesGetDataLifecycleStatsRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f IndicesGetDataLifecycleStats) WithContext(v context.Context) func(*IndicesGetDataLifecycleStatsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesGetDataLifecycleStats) WithPretty() func(*IndicesGetDataLifecycleStatsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesGetDataLifecycleStats) WithHuman() func(*IndicesGetDataLifecycleStatsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesGetDataLifecycleStats) WithErrorTrace() func(*IndicesGetDataLifecycleStatsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesGetDataLifecycleStats) WithFilterPath(v ...string) func(*IndicesGetDataLifecycleStatsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesGetDataLifecycleStats) WithHeader(h map[string]string) func(*IndicesGetDataLifecycleStatsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesGetDataLifecycleStats) WithOpaqueID(s string) func(*IndicesGetDataLifecycleStatsRequest) {
	_ = "STUB: not implemented"
	return nil
}
