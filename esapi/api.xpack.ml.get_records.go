package esapi

import (
	"context"
	"io"
	"net/http"
)

func newMLGetRecordsFunc(t Transport) MLGetRecords {
	_ = "STUB: not implemented"
	return *new(MLGetRecords)
}

type MLGetRecords func(job_id string, o ...func(*MLGetRecordsRequest)) (*Response, error)

type MLGetRecordsRequest struct {
	Body io.Reader

	JobID string

	Desc           *bool
	End            string
	ExcludeInterim *bool
	From           *int
	RecordScore    interface{}
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

func (r MLGetRecordsRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f MLGetRecords) WithContext(v context.Context) func(*MLGetRecordsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLGetRecords) WithBody(v io.Reader) func(*MLGetRecordsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLGetRecords) WithDesc(v bool) func(*MLGetRecordsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLGetRecords) WithEnd(v string) func(*MLGetRecordsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLGetRecords) WithExcludeInterim(v bool) func(*MLGetRecordsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLGetRecords) WithFrom(v int) func(*MLGetRecordsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLGetRecords) WithRecordScore(v interface{}) func(*MLGetRecordsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLGetRecords) WithSize(v int) func(*MLGetRecordsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLGetRecords) WithSort(v string) func(*MLGetRecordsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLGetRecords) WithStart(v string) func(*MLGetRecordsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLGetRecords) WithPretty() func(*MLGetRecordsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLGetRecords) WithHuman() func(*MLGetRecordsRequest) { _ = "STUB: not implemented"; return nil }

func (f MLGetRecords) WithErrorTrace() func(*MLGetRecordsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLGetRecords) WithFilterPath(v ...string) func(*MLGetRecordsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLGetRecords) WithHeader(h map[string]string) func(*MLGetRecordsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLGetRecords) WithOpaqueID(s string) func(*MLGetRecordsRequest) {
	_ = "STUB: not implemented"
	return nil
}
