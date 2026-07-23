package esapi

import (
	"context"
	"io"
	"net/http"
)

func newMLGetBucketsFunc(t Transport) MLGetBuckets {
	_ = "STUB: not implemented"
	return *new(MLGetBuckets)
}

type MLGetBuckets func(job_id string, o ...func(*MLGetBucketsRequest)) (*Response, error)

type MLGetBucketsRequest struct {
	Body io.Reader

	JobID     string
	Timestamp string

	AnomalyScore   interface{}
	Desc           *bool
	End            string
	ExcludeInterim *bool
	Expand         *bool
	From           *int
	Size           *int
	Sort           string
	Start          string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r MLGetBucketsRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f MLGetBuckets) WithContext(v context.Context) func(*MLGetBucketsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLGetBuckets) WithBody(v io.Reader) func(*MLGetBucketsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLGetBuckets) WithTimestamp(v string) func(*MLGetBucketsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLGetBuckets) WithAnomalyScore(v interface{}) func(*MLGetBucketsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLGetBuckets) WithDesc(v bool) func(*MLGetBucketsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLGetBuckets) WithEnd(v string) func(*MLGetBucketsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLGetBuckets) WithExcludeInterim(v bool) func(*MLGetBucketsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLGetBuckets) WithExpand(v bool) func(*MLGetBucketsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLGetBuckets) WithFrom(v int) func(*MLGetBucketsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLGetBuckets) WithSize(v int) func(*MLGetBucketsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLGetBuckets) WithSort(v string) func(*MLGetBucketsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLGetBuckets) WithStart(v string) func(*MLGetBucketsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLGetBuckets) WithPretty() func(*MLGetBucketsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLGetBuckets) WithHuman() func(*MLGetBucketsRequest) { _ = "STUB: not implemented"; return nil }

func (f MLGetBuckets) WithErrorTrace() func(*MLGetBucketsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLGetBuckets) WithFilterPath(v ...string) func(*MLGetBucketsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLGetBuckets) WithHeader(h map[string]string) func(*MLGetBucketsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLGetBuckets) WithOpaqueID(s string) func(*MLGetBucketsRequest) {
	_ = "STUB: not implemented"
	return nil
}
