package esapi

import (
	"context"
	"net/http"
)

func newMLGetCalendarEventsFunc(t Transport) MLGetCalendarEvents {
	_ = "STUB: not implemented"
	return *new(MLGetCalendarEvents)
}

type MLGetCalendarEvents func(calendar_id string, o ...func(*MLGetCalendarEventsRequest)) (*Response, error)

type MLGetCalendarEventsRequest struct {
	CalendarID string

	End   string
	From  *int
	JobID string
	Size  *int
	Start string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r MLGetCalendarEventsRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f MLGetCalendarEvents) WithContext(v context.Context) func(*MLGetCalendarEventsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLGetCalendarEvents) WithEnd(v string) func(*MLGetCalendarEventsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLGetCalendarEvents) WithFrom(v int) func(*MLGetCalendarEventsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLGetCalendarEvents) WithJobID(v string) func(*MLGetCalendarEventsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLGetCalendarEvents) WithSize(v int) func(*MLGetCalendarEventsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLGetCalendarEvents) WithStart(v string) func(*MLGetCalendarEventsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLGetCalendarEvents) WithPretty() func(*MLGetCalendarEventsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLGetCalendarEvents) WithHuman() func(*MLGetCalendarEventsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLGetCalendarEvents) WithErrorTrace() func(*MLGetCalendarEventsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLGetCalendarEvents) WithFilterPath(v ...string) func(*MLGetCalendarEventsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLGetCalendarEvents) WithHeader(h map[string]string) func(*MLGetCalendarEventsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLGetCalendarEvents) WithOpaqueID(s string) func(*MLGetCalendarEventsRequest) {
	_ = "STUB: not implemented"
	return nil
}
