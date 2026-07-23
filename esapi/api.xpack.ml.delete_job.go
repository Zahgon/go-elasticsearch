package esapi

import (
	"context"
	"net/http"
)

func newMLDeleteJobFunc(t Transport) MLDeleteJob {
	_ = "STUB: not implemented"
	return *new(MLDeleteJob)
}

type MLDeleteJob func(job_id string, o ...func(*MLDeleteJobRequest)) (*Response, error)

type MLDeleteJobRequest struct {
	JobID string

	DeleteUserAnnotations *bool
	Force                 *bool
	WaitForCompletion     *bool

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r MLDeleteJobRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f MLDeleteJob) WithContext(v context.Context) func(*MLDeleteJobRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLDeleteJob) WithDeleteUserAnnotations(v bool) func(*MLDeleteJobRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLDeleteJob) WithForce(v bool) func(*MLDeleteJobRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLDeleteJob) WithWaitForCompletion(v bool) func(*MLDeleteJobRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLDeleteJob) WithPretty() func(*MLDeleteJobRequest) { _ = "STUB: not implemented"; return nil }

func (f MLDeleteJob) WithHuman() func(*MLDeleteJobRequest) { _ = "STUB: not implemented"; return nil }

func (f MLDeleteJob) WithErrorTrace() func(*MLDeleteJobRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLDeleteJob) WithFilterPath(v ...string) func(*MLDeleteJobRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLDeleteJob) WithHeader(h map[string]string) func(*MLDeleteJobRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLDeleteJob) WithOpaqueID(s string) func(*MLDeleteJobRequest) {
	_ = "STUB: not implemented"
	return nil
}
