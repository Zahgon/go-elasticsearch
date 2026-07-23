package esapi

import (
	"context"
	"net/http"
	"time"
)

func newSearchShardsFunc(t Transport) SearchShards {
	_ = "STUB: not implemented"
	return *new(SearchShards)
}

type SearchShards func(o ...func(*SearchShardsRequest)) (*Response, error)

type SearchShardsRequest struct {
	Index []string

	AllowNoIndices    *bool
	ExpandWildcards   []string
	IgnoreUnavailable *bool
	Local             *bool
	MasterTimeout     time.Duration
	Preference        string
	Routing           []string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r SearchShardsRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f SearchShards) WithContext(v context.Context) func(*SearchShardsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SearchShards) WithIndex(v ...string) func(*SearchShardsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SearchShards) WithAllowNoIndices(v bool) func(*SearchShardsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SearchShards) WithExpandWildcards(v ...string) func(*SearchShardsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SearchShards) WithIgnoreUnavailable(v bool) func(*SearchShardsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SearchShards) WithLocal(v bool) func(*SearchShardsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SearchShards) WithMasterTimeout(v time.Duration) func(*SearchShardsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SearchShards) WithPreference(v string) func(*SearchShardsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SearchShards) WithRouting(v ...string) func(*SearchShardsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SearchShards) WithPretty() func(*SearchShardsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SearchShards) WithHuman() func(*SearchShardsRequest) { _ = "STUB: not implemented"; return nil }

func (f SearchShards) WithErrorTrace() func(*SearchShardsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SearchShards) WithFilterPath(v ...string) func(*SearchShardsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SearchShards) WithHeader(h map[string]string) func(*SearchShardsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f SearchShards) WithOpaqueID(s string) func(*SearchShardsRequest) {
	_ = "STUB: not implemented"
	return nil
}
