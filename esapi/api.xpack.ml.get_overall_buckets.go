package esapi

import (
	"context"
	"io"
	"net/http"
	"time"
)

func newMLGetOverallBucketsFunc(t Transport) MLGetOverallBuckets {
	_ = "STUB: not implemented"
	return *new(MLGetOverallBuckets)
}

type MLGetOverallBuckets func(job_id string, o ...func(*MLGetOverallBucketsRequest)) (*Response, error)

type MLGetOverallBucketsRequest struct {
	Body io.Reader

	JobID string

	AllowNoMatch   *bool
	BucketSpan     time.Duration
	End            string
	ExcludeInterim *bool
	OverallScore   interface{}
	Start          string
	TopN           *int

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r MLGetOverallBucketsRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f MLGetOverallBuckets) WithContext(v context.Context) func(*MLGetOverallBucketsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLGetOverallBuckets) WithBody(v io.Reader) func(*MLGetOverallBucketsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLGetOverallBuckets) WithAllowNoMatch(v bool) func(*MLGetOverallBucketsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLGetOverallBuckets) WithBucketSpan(v time.Duration) func(*MLGetOverallBucketsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLGetOverallBuckets) WithEnd(v string) func(*MLGetOverallBucketsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLGetOverallBuckets) WithExcludeInterim(v bool) func(*MLGetOverallBucketsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLGetOverallBuckets) WithOverallScore(v interface{}) func(*MLGetOverallBucketsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLGetOverallBuckets) WithStart(v string) func(*MLGetOverallBucketsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLGetOverallBuckets) WithTopN(v int) func(*MLGetOverallBucketsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLGetOverallBuckets) WithPretty() func(*MLGetOverallBucketsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLGetOverallBuckets) WithHuman() func(*MLGetOverallBucketsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLGetOverallBuckets) WithErrorTrace() func(*MLGetOverallBucketsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLGetOverallBuckets) WithFilterPath(v ...string) func(*MLGetOverallBucketsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLGetOverallBuckets) WithHeader(h map[string]string) func(*MLGetOverallBucketsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLGetOverallBuckets) WithOpaqueID(s string) func(*MLGetOverallBucketsRequest) {
	_ = "STUB: not implemented"
	return nil
}
