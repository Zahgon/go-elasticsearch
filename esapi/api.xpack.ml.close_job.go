package esapi

import (
	"context"
	"io"
	"net/http"
	"time"
)

func newMLCloseJobFunc(t Transport) MLCloseJob { _ = "STUB: not implemented"; return *new(MLCloseJob) }

type MLCloseJob func(job_id string, o ...func(*MLCloseJobRequest)) (*Response, error)

type MLCloseJobRequest struct {
	Body io.Reader

	JobID string

	AllowNoMatch *bool
	Force        *bool
	Timeout      time.Duration

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r MLCloseJobRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f MLCloseJob) WithContext(v context.Context) func(*MLCloseJobRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLCloseJob) WithBody(v io.Reader) func(*MLCloseJobRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLCloseJob) WithAllowNoMatch(v bool) func(*MLCloseJobRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLCloseJob) WithForce(v bool) func(*MLCloseJobRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLCloseJob) WithTimeout(v time.Duration) func(*MLCloseJobRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLCloseJob) WithPretty() func(*MLCloseJobRequest) { _ = "STUB: not implemented"; return nil }

func (f MLCloseJob) WithHuman() func(*MLCloseJobRequest) { _ = "STUB: not implemented"; return nil }

func (f MLCloseJob) WithErrorTrace() func(*MLCloseJobRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLCloseJob) WithFilterPath(v ...string) func(*MLCloseJobRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLCloseJob) WithHeader(h map[string]string) func(*MLCloseJobRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLCloseJob) WithOpaqueID(s string) func(*MLCloseJobRequest) {
	_ = "STUB: not implemented"
	return nil
}
