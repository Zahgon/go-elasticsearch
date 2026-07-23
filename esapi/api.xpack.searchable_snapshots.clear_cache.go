package esapi

import (
	"context"
	"net/http"
)

func newSearchableSnapshotsClearCacheFunc(t Transport) SearchableSnapshotsClearCache {
	_ = "STUB: not implemented"
	return *new(SearchableSnapshotsClearCache)
}

type SearchableSnapshotsClearCache func(o ...func(*SearchableSnapshotsClearCacheRequest)) (*Response, error)

type SearchableSnapshotsClearCacheRequest struct {
	Index []string

	AllowNoIndices    *bool
	ExpandWildcards   []string
	IgnoreUnavailable *bool

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r SearchableSnapshotsClearCacheRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f SearchableSnapshotsClearCache) WithContext(v context.Context) func(*SearchableSnapshotsClearCacheRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SearchableSnapshotsClearCache) WithIndex(v ...string) func(*SearchableSnapshotsClearCacheRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SearchableSnapshotsClearCache) WithAllowNoIndices(v bool) func(*SearchableSnapshotsClearCacheRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SearchableSnapshotsClearCache) WithExpandWildcards(v ...string) func(*SearchableSnapshotsClearCacheRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SearchableSnapshotsClearCache) WithIgnoreUnavailable(v bool) func(*SearchableSnapshotsClearCacheRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SearchableSnapshotsClearCache) WithPretty() func(*SearchableSnapshotsClearCacheRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SearchableSnapshotsClearCache) WithHuman() func(*SearchableSnapshotsClearCacheRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SearchableSnapshotsClearCache) WithErrorTrace() func(*SearchableSnapshotsClearCacheRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SearchableSnapshotsClearCache) WithFilterPath(v ...string) func(*SearchableSnapshotsClearCacheRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SearchableSnapshotsClearCache) WithHeader(h map[string]string) func(*SearchableSnapshotsClearCacheRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SearchableSnapshotsClearCache) WithOpaqueID(s string) func(*SearchableSnapshotsClearCacheRequest) {
	_ = "STUB: not implemented"
	return nil
}
