package esapi

import (
	"context"
	"io"
	"net/http"
)

func newMLUpdateJobFunc(t Transport) MLUpdateJob {
	_ = "STUB: not implemented"
	return *new(MLUpdateJob)
}

type MLUpdateJob func(job_id string, body io.Reader, o ...func(*MLUpdateJobRequest)) (*Response, error)

type MLUpdateJobRequest struct {
	Body io.Reader

	JobID string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r MLUpdateJobRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f MLUpdateJob) WithContext(v context.Context) func(*MLUpdateJobRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLUpdateJob) WithPretty() func(*MLUpdateJobRequest) { _ = "STUB: not implemented"; return nil }

func (f MLUpdateJob) WithHuman() func(*MLUpdateJobRequest) { _ = "STUB: not implemented"; return nil }

func (f MLUpdateJob) WithErrorTrace() func(*MLUpdateJobRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLUpdateJob) WithFilterPath(v ...string) func(*MLUpdateJobRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLUpdateJob) WithHeader(h map[string]string) func(*MLUpdateJobRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLUpdateJob) WithOpaqueID(s string) func(*MLUpdateJobRequest) {
	_ = "STUB: not implemented"
	return nil
}
