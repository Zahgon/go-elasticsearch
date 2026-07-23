package esapi

import (
	"context"
	"io"
	"net/http"
	"time"
)

func newMLOpenJobFunc(t Transport) MLOpenJob { _ = "STUB: not implemented"; return *new(MLOpenJob) }

type MLOpenJob func(job_id string, o ...func(*MLOpenJobRequest)) (*Response, error)

type MLOpenJobRequest struct {
	Body io.Reader

	JobID string

	Timeout time.Duration

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r MLOpenJobRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f MLOpenJob) WithContext(v context.Context) func(*MLOpenJobRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLOpenJob) WithBody(v io.Reader) func(*MLOpenJobRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLOpenJob) WithTimeout(v time.Duration) func(*MLOpenJobRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLOpenJob) WithPretty() func(*MLOpenJobRequest) { _ = "STUB: not implemented"; return nil }

func (f MLOpenJob) WithHuman() func(*MLOpenJobRequest) { _ = "STUB: not implemented"; return nil }

func (f MLOpenJob) WithErrorTrace() func(*MLOpenJobRequest) { _ = "STUB: not implemented"; return nil }

func (f MLOpenJob) WithFilterPath(v ...string) func(*MLOpenJobRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLOpenJob) WithHeader(h map[string]string) func(*MLOpenJobRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLOpenJob) WithOpaqueID(s string) func(*MLOpenJobRequest) {
	_ = "STUB: not implemented"
	return nil
}
