package esapi

import (
	"context"
	"net/http"
)

func newMLDeleteCalendarFunc(t Transport) MLDeleteCalendar {
	_ = "STUB: not implemented"
	return *new(MLDeleteCalendar)
}

type MLDeleteCalendar func(calendar_id string, o ...func(*MLDeleteCalendarRequest)) (*Response, error)

type MLDeleteCalendarRequest struct {
	CalendarID string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r MLDeleteCalendarRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f MLDeleteCalendar) WithContext(v context.Context) func(*MLDeleteCalendarRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLDeleteCalendar) WithPretty() func(*MLDeleteCalendarRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLDeleteCalendar) WithHuman() func(*MLDeleteCalendarRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLDeleteCalendar) WithErrorTrace() func(*MLDeleteCalendarRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLDeleteCalendar) WithFilterPath(v ...string) func(*MLDeleteCalendarRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLDeleteCalendar) WithHeader(h map[string]string) func(*MLDeleteCalendarRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLDeleteCalendar) WithOpaqueID(s string) func(*MLDeleteCalendarRequest) {
	_ = "STUB: not implemented"
	return nil
}
