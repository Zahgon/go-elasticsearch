package esapi

import (
	"context"
	"net/http"
)

func newSearchableSnapshotsCacheStatsFunc(t Transport) SearchableSnapshotsCacheStats {
	_ = "STUB: not implemented"
	return *new(SearchableSnapshotsCacheStats)
}

type SearchableSnapshotsCacheStats func(o ...func(*SearchableSnapshotsCacheStatsRequest)) (*Response, error)

type SearchableSnapshotsCacheStatsRequest struct {
	NodeID []string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r SearchableSnapshotsCacheStatsRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f SearchableSnapshotsCacheStats) WithContext(v context.Context) func(*SearchableSnapshotsCacheStatsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SearchableSnapshotsCacheStats) WithNodeID(v ...string) func(*SearchableSnapshotsCacheStatsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SearchableSnapshotsCacheStats) WithPretty() func(*SearchableSnapshotsCacheStatsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SearchableSnapshotsCacheStats) WithHuman() func(*SearchableSnapshotsCacheStatsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SearchableSnapshotsCacheStats) WithErrorTrace() func(*SearchableSnapshotsCacheStatsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SearchableSnapshotsCacheStats) WithFilterPath(v ...string) func(*SearchableSnapshotsCacheStatsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SearchableSnapshotsCacheStats) WithHeader(h map[string]string) func(*SearchableSnapshotsCacheStatsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SearchableSnapshotsCacheStats) WithOpaqueID(s string) func(*SearchableSnapshotsCacheStatsRequest) {
	_ = "STUB: not implemented"
	return nil
}
