package esapi

import (
	"context"
	"net/http"
)

func newTransformGetNodeStatsFunc(t Transport) TransformGetNodeStats {
	_ = "STUB: not implemented"
	return *new(TransformGetNodeStats)
}

type TransformGetNodeStats func(o ...func(*TransformGetNodeStatsRequest)) (*Response, error)

type TransformGetNodeStatsRequest struct {
	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r TransformGetNodeStatsRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f TransformGetNodeStats) WithContext(v context.Context) func(*TransformGetNodeStatsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f TransformGetNodeStats) WithPretty() func(*TransformGetNodeStatsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f TransformGetNodeStats) WithHuman() func(*TransformGetNodeStatsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f TransformGetNodeStats) WithErrorTrace() func(*TransformGetNodeStatsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f TransformGetNodeStats) WithFilterPath(v ...string) func(*TransformGetNodeStatsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f TransformGetNodeStats) WithHeader(h map[string]string) func(*TransformGetNodeStatsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f TransformGetNodeStats) WithOpaqueID(s string) func(*TransformGetNodeStatsRequest) {
	_ = "STUB: not implemented"
	return nil
}
