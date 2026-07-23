package esapi

import (
	"context"
	"net/http"
)

func newCatMLJobsFunc(t Transport) CatMLJobs { _ = "STUB: not implemented"; return *new(CatMLJobs) }

type CatMLJobs func(o ...func(*CatMLJobsRequest)) (*Response, error)

type CatMLJobsRequest struct {
	JobID string

	AllowNoMatch *bool
	Bytes        string
	Format       string
	H            []string
	Help         *bool
	S            []string
	Time         string
	V            *bool

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r CatMLJobsRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f CatMLJobs) WithContext(v context.Context) func(*CatMLJobsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatMLJobs) WithJobID(v string) func(*CatMLJobsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatMLJobs) WithAllowNoMatch(v bool) func(*CatMLJobsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatMLJobs) WithBytes(v string) func(*CatMLJobsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatMLJobs) WithFormat(v string) func(*CatMLJobsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatMLJobs) WithH(v ...string) func(*CatMLJobsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatMLJobs) WithHelp(v bool) func(*CatMLJobsRequest) { _ = "STUB: not implemented"; return nil }

func (f CatMLJobs) WithS(v ...string) func(*CatMLJobsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatMLJobs) WithTime(v string) func(*CatMLJobsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatMLJobs) WithV(v bool) func(*CatMLJobsRequest) { _ = "STUB: not implemented"; return nil }

func (f CatMLJobs) WithPretty() func(*CatMLJobsRequest) { _ = "STUB: not implemented"; return nil }

func (f CatMLJobs) WithHuman() func(*CatMLJobsRequest) { _ = "STUB: not implemented"; return nil }

func (f CatMLJobs) WithErrorTrace() func(*CatMLJobsRequest) { _ = "STUB: not implemented"; return nil }

func (f CatMLJobs) WithFilterPath(v ...string) func(*CatMLJobsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatMLJobs) WithHeader(h map[string]string) func(*CatMLJobsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f CatMLJobs) WithOpaqueID(s string) func(*CatMLJobsRequest) {
	_ = "STUB: not implemented"
	return nil
}
