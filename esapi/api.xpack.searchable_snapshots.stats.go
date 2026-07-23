package esapi

import (
	"context"
	"net/http"
)

func newSearchableSnapshotsStatsFunc(t Transport) SearchableSnapshotsStats {
	_ = "STUB: not implemented"
	return *new(SearchableSnapshotsStats)
}

type SearchableSnapshotsStats func(o ...func(*SearchableSnapshotsStatsRequest)) (*Response, error)

type SearchableSnapshotsStatsRequest struct {
	Index []string

	Level string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r SearchableSnapshotsStatsRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f SearchableSnapshotsStats) WithContext(v context.Context) func(*SearchableSnapshotsStatsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SearchableSnapshotsStats) WithIndex(v ...string) func(*SearchableSnapshotsStatsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SearchableSnapshotsStats) WithLevel(v string) func(*SearchableSnapshotsStatsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SearchableSnapshotsStats) WithPretty() func(*SearchableSnapshotsStatsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SearchableSnapshotsStats) WithHuman() func(*SearchableSnapshotsStatsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SearchableSnapshotsStats) WithErrorTrace() func(*SearchableSnapshotsStatsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SearchableSnapshotsStats) WithFilterPath(v ...string) func(*SearchableSnapshotsStatsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SearchableSnapshotsStats) WithHeader(h map[string]string) func(*SearchableSnapshotsStatsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SearchableSnapshotsStats) WithOpaqueID(s string) func(*SearchableSnapshotsStatsRequest) {
	_ = "STUB: not implemented"
	return nil
}
