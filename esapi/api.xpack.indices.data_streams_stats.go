package esapi

import (
	"context"
	"net/http"
)

func newIndicesDataStreamsStatsFunc(t Transport) IndicesDataStreamsStats {
	_ = "STUB: not implemented"
	return *new(IndicesDataStreamsStats)
}

type IndicesDataStreamsStats func(o ...func(*IndicesDataStreamsStatsRequest)) (*Response, error)

type IndicesDataStreamsStatsRequest struct {
	Name []string

	ExpandWildcards []string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r IndicesDataStreamsStatsRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f IndicesDataStreamsStats) WithContext(v context.Context) func(*IndicesDataStreamsStatsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesDataStreamsStats) WithName(v ...string) func(*IndicesDataStreamsStatsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesDataStreamsStats) WithExpandWildcards(v ...string) func(*IndicesDataStreamsStatsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesDataStreamsStats) WithPretty() func(*IndicesDataStreamsStatsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesDataStreamsStats) WithHuman() func(*IndicesDataStreamsStatsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesDataStreamsStats) WithErrorTrace() func(*IndicesDataStreamsStatsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesDataStreamsStats) WithFilterPath(v ...string) func(*IndicesDataStreamsStatsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesDataStreamsStats) WithHeader(h map[string]string) func(*IndicesDataStreamsStatsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f IndicesDataStreamsStats) WithOpaqueID(s string) func(*IndicesDataStreamsStatsRequest) {
	_ = "STUB: not implemented"
	return nil
}
