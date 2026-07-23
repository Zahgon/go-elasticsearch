package esapi

import (
	"context"
	"net/http"
)

func newRollupStartJobFunc(t Transport) RollupStartJob {
	_ = "STUB: not implemented"
	return *new(RollupStartJob)
}

type RollupStartJob func(id string, o ...func(*RollupStartJobRequest)) (*Response, error)

type RollupStartJobRequest struct {
	JobID string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r RollupStartJobRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f RollupStartJob) WithContext(v context.Context) func(*RollupStartJobRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f RollupStartJob) WithPretty() func(*RollupStartJobRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f RollupStartJob) WithHuman() func(*RollupStartJobRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f RollupStartJob) WithErrorTrace() func(*RollupStartJobRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f RollupStartJob) WithFilterPath(v ...string) func(*RollupStartJobRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f RollupStartJob) WithHeader(h map[string]string) func(*RollupStartJobRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f RollupStartJob) WithOpaqueID(s string) func(*RollupStartJobRequest) {
	_ = "STUB: not implemented"
	return nil
}
