package esapi

import (
	"context"
	"net/http"
)

func newRollupDeleteJobFunc(t Transport) RollupDeleteJob {
	_ = "STUB: not implemented"
	return *new(RollupDeleteJob)
}

type RollupDeleteJob func(id string, o ...func(*RollupDeleteJobRequest)) (*Response, error)

type RollupDeleteJobRequest struct {
	JobID string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r RollupDeleteJobRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f RollupDeleteJob) WithContext(v context.Context) func(*RollupDeleteJobRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f RollupDeleteJob) WithPretty() func(*RollupDeleteJobRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f RollupDeleteJob) WithHuman() func(*RollupDeleteJobRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f RollupDeleteJob) WithErrorTrace() func(*RollupDeleteJobRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f RollupDeleteJob) WithFilterPath(v ...string) func(*RollupDeleteJobRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f RollupDeleteJob) WithHeader(h map[string]string) func(*RollupDeleteJobRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f RollupDeleteJob) WithOpaqueID(s string) func(*RollupDeleteJobRequest) {
	_ = "STUB: not implemented"
	return nil
}
