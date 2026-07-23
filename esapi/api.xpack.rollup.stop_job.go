package esapi

import (
	"context"
	"net/http"
	"time"
)

func newRollupStopJobFunc(t Transport) RollupStopJob {
	_ = "STUB: not implemented"
	return *new(RollupStopJob)
}

type RollupStopJob func(id string, o ...func(*RollupStopJobRequest)) (*Response, error)

type RollupStopJobRequest struct {
	JobID string

	Timeout           time.Duration
	WaitForCompletion *bool

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r RollupStopJobRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f RollupStopJob) WithContext(v context.Context) func(*RollupStopJobRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f RollupStopJob) WithTimeout(v time.Duration) func(*RollupStopJobRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f RollupStopJob) WithWaitForCompletion(v bool) func(*RollupStopJobRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f RollupStopJob) WithPretty() func(*RollupStopJobRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f RollupStopJob) WithHuman() func(*RollupStopJobRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f RollupStopJob) WithErrorTrace() func(*RollupStopJobRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f RollupStopJob) WithFilterPath(v ...string) func(*RollupStopJobRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f RollupStopJob) WithHeader(h map[string]string) func(*RollupStopJobRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f RollupStopJob) WithOpaqueID(s string) func(*RollupStopJobRequest) {
	_ = "STUB: not implemented"
	return nil
}
