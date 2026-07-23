package esapi

import (
	"context"
	"net/http"
)

func newIndicesStatsFunc(t Transport) IndicesStats {
	_ = "STUB: not implemented"
	return *new(IndicesStats)
}

type IndicesStats func(o ...func(*IndicesStatsRequest)) (*Response, error)

type IndicesStatsRequest struct {
	Index []string

	Metric []string

	CompletionFields        []string
	ExpandWildcards         []string
	FielddataFields         []string
	Fields                  []string
	ForbidClosedIndices     *bool
	Groups                  []string
	IncludeSegmentFileSizes *bool
	IncludeUnloadedSegments *bool
	Level                   string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r IndicesStatsRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f IndicesStats) WithContext(v context.Context) func(*IndicesStatsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesStats) WithIndex(v ...string) func(*IndicesStatsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesStats) WithMetric(v ...string) func(*IndicesStatsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesStats) WithCompletionFields(v ...string) func(*IndicesStatsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesStats) WithExpandWildcards(v ...string) func(*IndicesStatsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesStats) WithFielddataFields(v ...string) func(*IndicesStatsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesStats) WithFields(v ...string) func(*IndicesStatsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesStats) WithForbidClosedIndices(v bool) func(*IndicesStatsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesStats) WithGroups(v ...string) func(*IndicesStatsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesStats) WithIncludeSegmentFileSizes(v bool) func(*IndicesStatsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesStats) WithIncludeUnloadedSegments(v bool) func(*IndicesStatsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesStats) WithLevel(v string) func(*IndicesStatsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesStats) WithPretty() func(*IndicesStatsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesStats) WithHuman() func(*IndicesStatsRequest) { _ = "STUB: not implemented"; return nil }

func (f IndicesStats) WithErrorTrace() func(*IndicesStatsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesStats) WithFilterPath(v ...string) func(*IndicesStatsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesStats) WithHeader(h map[string]string) func(*IndicesStatsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesStats) WithOpaqueID(s string) func(*IndicesStatsRequest) {
	_ = "STUB: not implemented"
	return nil
}
