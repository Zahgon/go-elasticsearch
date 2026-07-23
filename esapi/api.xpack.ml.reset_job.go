package esapi

import (
	"context"
	"net/http"
)

func newMLResetJobFunc(t Transport) MLResetJob { _ = "STUB: not implemented"; return *new(MLResetJob) }

type MLResetJob func(job_id string, o ...func(*MLResetJobRequest)) (*Response, error)

type MLResetJobRequest struct {
	JobID string

	DeleteUserAnnotations *bool
	WaitForCompletion     *bool

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r MLResetJobRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f MLResetJob) WithContext(v context.Context) func(*MLResetJobRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLResetJob) WithDeleteUserAnnotations(v bool) func(*MLResetJobRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLResetJob) WithWaitForCompletion(v bool) func(*MLResetJobRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLResetJob) WithPretty() func(*MLResetJobRequest) { _ = "STUB: not implemented"; return nil }

func (f MLResetJob) WithHuman() func(*MLResetJobRequest) { _ = "STUB: not implemented"; return nil }

func (f MLResetJob) WithErrorTrace() func(*MLResetJobRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLResetJob) WithFilterPath(v ...string) func(*MLResetJobRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLResetJob) WithHeader(h map[string]string) func(*MLResetJobRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLResetJob) WithOpaqueID(s string) func(*MLResetJobRequest) {
	_ = "STUB: not implemented"
	return nil
}
