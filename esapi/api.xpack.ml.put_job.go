package esapi

import (
	"context"
	"io"
	"net/http"
)

func newMLPutJobFunc(t Transport) MLPutJob { _ = "STUB: not implemented"; return *new(MLPutJob) }

type MLPutJob func(job_id string, body io.Reader, o ...func(*MLPutJobRequest)) (*Response, error)

type MLPutJobRequest struct {
	Body io.Reader

	JobID string

	AllowNoIndices    *bool
	ExpandWildcards   []string
	IgnoreThrottled   *bool
	IgnoreUnavailable *bool

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r MLPutJobRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f MLPutJob) WithContext(v context.Context) func(*MLPutJobRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLPutJob) WithAllowNoIndices(v bool) func(*MLPutJobRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLPutJob) WithExpandWildcards(v ...string) func(*MLPutJobRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLPutJob) WithIgnoreThrottled(v bool) func(*MLPutJobRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLPutJob) WithIgnoreUnavailable(v bool) func(*MLPutJobRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLPutJob) WithPretty() func(*MLPutJobRequest) { _ = "STUB: not implemented"; return nil }

func (f MLPutJob) WithHuman() func(*MLPutJobRequest) { _ = "STUB: not implemented"; return nil }

func (f MLPutJob) WithErrorTrace() func(*MLPutJobRequest) { _ = "STUB: not implemented"; return nil }

func (f MLPutJob) WithFilterPath(v ...string) func(*MLPutJobRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLPutJob) WithHeader(h map[string]string) func(*MLPutJobRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLPutJob) WithOpaqueID(s string) func(*MLPutJobRequest) {
	_ = "STUB: not implemented"
	return nil
}
