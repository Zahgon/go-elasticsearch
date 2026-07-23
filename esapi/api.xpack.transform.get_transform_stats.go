package esapi

import (
	"context"
	"net/http"
	"time"
)

func newTransformGetTransformStatsFunc(t Transport) TransformGetTransformStats {
	_ = "STUB: not implemented"
	return *new(TransformGetTransformStats)
}

type TransformGetTransformStats func(transform_id []string, o ...func(*TransformGetTransformStatsRequest)) (*Response, error)

type TransformGetTransformStatsRequest struct {
	TransformID []string

	AllowNoMatch *bool
	From         *int64
	Size         *int64
	Timeout      time.Duration

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r TransformGetTransformStatsRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f TransformGetTransformStats) WithContext(v context.Context) func(*TransformGetTransformStatsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f TransformGetTransformStats) WithAllowNoMatch(v bool) func(*TransformGetTransformStatsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f TransformGetTransformStats) WithFrom(v int64) func(*TransformGetTransformStatsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f TransformGetTransformStats) WithSize(v int64) func(*TransformGetTransformStatsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f TransformGetTransformStats) WithTimeout(v time.Duration) func(*TransformGetTransformStatsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f TransformGetTransformStats) WithPretty() func(*TransformGetTransformStatsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f TransformGetTransformStats) WithHuman() func(*TransformGetTransformStatsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f TransformGetTransformStats) WithErrorTrace() func(*TransformGetTransformStatsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f TransformGetTransformStats) WithFilterPath(v ...string) func(*TransformGetTransformStatsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f TransformGetTransformStats) WithHeader(h map[string]string) func(*TransformGetTransformStatsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f TransformGetTransformStats) WithOpaqueID(s string) func(*TransformGetTransformStatsRequest) {
	_ = "STUB: not implemented"
	return nil
}
