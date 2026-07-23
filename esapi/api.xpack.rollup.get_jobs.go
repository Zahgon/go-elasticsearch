package esapi

import (
	"context"
	"net/http"
)

func newRollupGetJobsFunc(t Transport) RollupGetJobs {
	_ = "STUB: not implemented"
	return *new(RollupGetJobs)
}

type RollupGetJobs func(o ...func(*RollupGetJobsRequest)) (*Response, error)

type RollupGetJobsRequest struct {
	JobID string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r RollupGetJobsRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f RollupGetJobs) WithContext(v context.Context) func(*RollupGetJobsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f RollupGetJobs) WithJobID(v string) func(*RollupGetJobsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f RollupGetJobs) WithPretty() func(*RollupGetJobsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f RollupGetJobs) WithHuman() func(*RollupGetJobsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f RollupGetJobs) WithErrorTrace() func(*RollupGetJobsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f RollupGetJobs) WithFilterPath(v ...string) func(*RollupGetJobsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f RollupGetJobs) WithHeader(h map[string]string) func(*RollupGetJobsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f RollupGetJobs) WithOpaqueID(s string) func(*RollupGetJobsRequest) {
	_ = "STUB: not implemented"
	return nil
}
