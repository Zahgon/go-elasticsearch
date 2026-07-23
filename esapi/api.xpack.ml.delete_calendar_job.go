package esapi

import (
	"context"
	"net/http"
)

func newMLDeleteCalendarJobFunc(t Transport) MLDeleteCalendarJob {
	_ = "STUB: not implemented"
	return *new(MLDeleteCalendarJob)
}

type MLDeleteCalendarJob func(calendar_id string, job_id []string, o ...func(*MLDeleteCalendarJobRequest)) (*Response, error)

type MLDeleteCalendarJobRequest struct {
	CalendarID string
	JobID      []string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r MLDeleteCalendarJobRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f MLDeleteCalendarJob) WithContext(v context.Context) func(*MLDeleteCalendarJobRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLDeleteCalendarJob) WithPretty() func(*MLDeleteCalendarJobRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLDeleteCalendarJob) WithHuman() func(*MLDeleteCalendarJobRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLDeleteCalendarJob) WithErrorTrace() func(*MLDeleteCalendarJobRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLDeleteCalendarJob) WithFilterPath(v ...string) func(*MLDeleteCalendarJobRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLDeleteCalendarJob) WithHeader(h map[string]string) func(*MLDeleteCalendarJobRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLDeleteCalendarJob) WithOpaqueID(s string) func(*MLDeleteCalendarJobRequest) {
	_ = "STUB: not implemented"
	return nil
}
