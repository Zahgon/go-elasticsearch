package esapi

import (
	"context"
	"net/http"
)

func newMLGetJobsFunc(t Transport) MLGetJobs { _ = "STUB: not implemented"; return *new(MLGetJobs) }

type MLGetJobs func(o ...func(*MLGetJobsRequest)) (*Response, error)

type MLGetJobsRequest struct {
	JobID []string

	AllowNoMatch     *bool
	ExcludeGenerated *bool

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r MLGetJobsRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f MLGetJobs) WithContext(v context.Context) func(*MLGetJobsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLGetJobs) WithJobID(v ...string) func(*MLGetJobsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLGetJobs) WithAllowNoMatch(v bool) func(*MLGetJobsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLGetJobs) WithExcludeGenerated(v bool) func(*MLGetJobsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLGetJobs) WithPretty() func(*MLGetJobsRequest) { _ = "STUB: not implemented"; return nil }

func (f MLGetJobs) WithHuman() func(*MLGetJobsRequest) { _ = "STUB: not implemented"; return nil }

func (f MLGetJobs) WithErrorTrace() func(*MLGetJobsRequest) { _ = "STUB: not implemented"; return nil }

func (f MLGetJobs) WithFilterPath(v ...string) func(*MLGetJobsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLGetJobs) WithHeader(h map[string]string) func(*MLGetJobsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLGetJobs) WithOpaqueID(s string) func(*MLGetJobsRequest) {
	_ = "STUB: not implemented"
	return nil
}
