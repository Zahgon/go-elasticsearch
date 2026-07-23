package esapi

import (
	"context"
	"net/http"
	"time"
)

func newNodesStatsFunc(t Transport) NodesStats { _ = "STUB: not implemented"; return *new(NodesStats) }

type NodesStats func(o ...func(*NodesStatsRequest)) (*Response, error)

type NodesStatsRequest struct {
	IndexMetric []string
	Metric      []string
	NodeID      []string

	CompletionFields        []string
	FielddataFields         []string
	Fields                  []string
	Groups                  *bool
	IncludeSegmentFileSizes *bool
	IncludeUnloadedSegments *bool
	Level                   string
	Timeout                 time.Duration
	Types                   []string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r NodesStatsRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f NodesStats) WithContext(v context.Context) func(*NodesStatsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f NodesStats) WithIndexMetric(v ...string) func(*NodesStatsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f NodesStats) WithMetric(v ...string) func(*NodesStatsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f NodesStats) WithNodeID(v ...string) func(*NodesStatsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f NodesStats) WithCompletionFields(v ...string) func(*NodesStatsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f NodesStats) WithFielddataFields(v ...string) func(*NodesStatsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f NodesStats) WithFields(v ...string) func(*NodesStatsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f NodesStats) WithGroups(v bool) func(*NodesStatsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f NodesStats) WithIncludeSegmentFileSizes(v bool) func(*NodesStatsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f NodesStats) WithIncludeUnloadedSegments(v bool) func(*NodesStatsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f NodesStats) WithLevel(v string) func(*NodesStatsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f NodesStats) WithTimeout(v time.Duration) func(*NodesStatsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f NodesStats) WithTypes(v ...string) func(*NodesStatsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f NodesStats) WithPretty() func(*NodesStatsRequest) { _ = "STUB: not implemented"; return nil }

func (f NodesStats) WithHuman() func(*NodesStatsRequest) { _ = "STUB: not implemented"; return nil }

func (f NodesStats) WithErrorTrace() func(*NodesStatsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f NodesStats) WithFilterPath(v ...string) func(*NodesStatsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f NodesStats) WithHeader(h map[string]string) func(*NodesStatsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f NodesStats) WithOpaqueID(s string) func(*NodesStatsRequest) {
	_ = "STUB: not implemented"
	return nil
}
