package esapi

import (
	"context"
	"net/http"
	"time"
)

func newMLGetMemoryStatsFunc(t Transport) MLGetMemoryStats {
	_ = "STUB: not implemented"
	return *new(MLGetMemoryStats)
}

type MLGetMemoryStats func(o ...func(*MLGetMemoryStatsRequest)) (*Response, error)

type MLGetMemoryStatsRequest struct {
	NodeID string

	MasterTimeout time.Duration
	Timeout       time.Duration

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r MLGetMemoryStatsRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f MLGetMemoryStats) WithContext(v context.Context) func(*MLGetMemoryStatsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLGetMemoryStats) WithNodeID(v string) func(*MLGetMemoryStatsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLGetMemoryStats) WithMasterTimeout(v time.Duration) func(*MLGetMemoryStatsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLGetMemoryStats) WithTimeout(v time.Duration) func(*MLGetMemoryStatsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLGetMemoryStats) WithPretty() func(*MLGetMemoryStatsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLGetMemoryStats) WithHuman() func(*MLGetMemoryStatsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLGetMemoryStats) WithErrorTrace() func(*MLGetMemoryStatsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLGetMemoryStats) WithFilterPath(v ...string) func(*MLGetMemoryStatsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLGetMemoryStats) WithHeader(h map[string]string) func(*MLGetMemoryStatsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLGetMemoryStats) WithOpaqueID(s string) func(*MLGetMemoryStatsRequest) {
	_ = "STUB: not implemented"
	return nil
}
