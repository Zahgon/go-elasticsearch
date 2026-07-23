package esapi

import (
	"context"
	"net/http"
	"time"
)

func newClusterStatsFunc(t Transport) ClusterStats {
	_ = "STUB: not implemented"
	return *new(ClusterStats)
}

type ClusterStats func(o ...func(*ClusterStatsRequest)) (*Response, error)

type ClusterStatsRequest struct {
	NodeID []string

	IncludeRemotes *bool
	Timeout        time.Duration

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r ClusterStatsRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f ClusterStats) WithContext(v context.Context) func(*ClusterStatsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ClusterStats) WithNodeID(v ...string) func(*ClusterStatsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ClusterStats) WithIncludeRemotes(v bool) func(*ClusterStatsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ClusterStats) WithTimeout(v time.Duration) func(*ClusterStatsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ClusterStats) WithPretty() func(*ClusterStatsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ClusterStats) WithHuman() func(*ClusterStatsRequest) { _ = "STUB: not implemented"; return nil }

func (f ClusterStats) WithErrorTrace() func(*ClusterStatsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ClusterStats) WithFilterPath(v ...string) func(*ClusterStatsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ClusterStats) WithHeader(h map[string]string) func(*ClusterStatsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f ClusterStats) WithOpaqueID(s string) func(*ClusterStatsRequest) {
	_ = "STUB: not implemented"
	return nil
}
