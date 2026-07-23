package esapi

import (
	"context"
	"net/http"
)

func newMLPutCalendarJobFunc(t Transport) MLPutCalendarJob {
	_ = "STUB: not implemented"
	return *new(MLPutCalendarJob)
}

type MLPutCalendarJob func(calendar_id string, job_id []string, o ...func(*MLPutCalendarJobRequest)) (*Response, error)

type MLPutCalendarJobRequest struct {
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

func (r MLPutCalendarJobRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f MLPutCalendarJob) WithContext(v context.Context) func(*MLPutCalendarJobRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLPutCalendarJob) WithPretty() func(*MLPutCalendarJobRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLPutCalendarJob) WithHuman() func(*MLPutCalendarJobRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLPutCalendarJob) WithErrorTrace() func(*MLPutCalendarJobRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLPutCalendarJob) WithFilterPath(v ...string) func(*MLPutCalendarJobRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLPutCalendarJob) WithHeader(h map[string]string) func(*MLPutCalendarJobRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLPutCalendarJob) WithOpaqueID(s string) func(*MLPutCalendarJobRequest) {
	_ = "STUB: not implemented"
	return nil
}
