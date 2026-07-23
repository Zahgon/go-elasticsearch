package esapi

import (
	"context"
	"net/http"
)

func newIndicesFieldUsageStatsFunc(t Transport) IndicesFieldUsageStats {
	_ = "STUB: not implemented"
	return *new(IndicesFieldUsageStats)
}

type IndicesFieldUsageStats func(index []string, o ...func(*IndicesFieldUsageStatsRequest)) (*Response, error)

type IndicesFieldUsageStatsRequest struct {
	Index []string

	AllowNoIndices    *bool
	ExpandWildcards   []string
	Fields            []string
	IgnoreUnavailable *bool

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r IndicesFieldUsageStatsRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f IndicesFieldUsageStats) WithContext(v context.Context) func(*IndicesFieldUsageStatsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesFieldUsageStats) WithAllowNoIndices(v bool) func(*IndicesFieldUsageStatsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesFieldUsageStats) WithExpandWildcards(v ...string) func(*IndicesFieldUsageStatsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesFieldUsageStats) WithFields(v ...string) func(*IndicesFieldUsageStatsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesFieldUsageStats) WithIgnoreUnavailable(v bool) func(*IndicesFieldUsageStatsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesFieldUsageStats) WithPretty() func(*IndicesFieldUsageStatsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesFieldUsageStats) WithHuman() func(*IndicesFieldUsageStatsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesFieldUsageStats) WithErrorTrace() func(*IndicesFieldUsageStatsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesFieldUsageStats) WithFilterPath(v ...string) func(*IndicesFieldUsageStatsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesFieldUsageStats) WithHeader(h map[string]string) func(*IndicesFieldUsageStatsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesFieldUsageStats) WithOpaqueID(s string) func(*IndicesFieldUsageStatsRequest) {
	_ = "STUB: not implemented"
	return nil
}
