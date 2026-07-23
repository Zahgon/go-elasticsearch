package esapi

import (
	"context"
	"io"
	"net/http"
)

func newMLFlushJobFunc(t Transport) MLFlushJob { _ = "STUB: not implemented"; return *new(MLFlushJob) }

type MLFlushJob func(job_id string, o ...func(*MLFlushJobRequest)) (*Response, error)

type MLFlushJobRequest struct {
	Body io.Reader

	JobID string

	AdvanceTime string
	CalcInterim *bool
	End         string
	SkipTime    string
	Start       string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r MLFlushJobRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f MLFlushJob) WithContext(v context.Context) func(*MLFlushJobRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLFlushJob) WithBody(v io.Reader) func(*MLFlushJobRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLFlushJob) WithAdvanceTime(v string) func(*MLFlushJobRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLFlushJob) WithCalcInterim(v bool) func(*MLFlushJobRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLFlushJob) WithEnd(v string) func(*MLFlushJobRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLFlushJob) WithSkipTime(v string) func(*MLFlushJobRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLFlushJob) WithStart(v string) func(*MLFlushJobRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLFlushJob) WithPretty() func(*MLFlushJobRequest) { _ = "STUB: not implemented"; return nil }

func (f MLFlushJob) WithHuman() func(*MLFlushJobRequest) { _ = "STUB: not implemented"; return nil }

func (f MLFlushJob) WithErrorTrace() func(*MLFlushJobRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLFlushJob) WithFilterPath(v ...string) func(*MLFlushJobRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLFlushJob) WithHeader(h map[string]string) func(*MLFlushJobRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLFlushJob) WithOpaqueID(s string) func(*MLFlushJobRequest) {
	_ = "STUB: not implemented"
	return nil
}
