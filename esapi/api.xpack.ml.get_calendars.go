package esapi

import (
	"context"
	"io"
	"net/http"
)

func newMLGetCalendarsFunc(t Transport) MLGetCalendars {
	_ = "STUB: not implemented"
	return *new(MLGetCalendars)
}

type MLGetCalendars func(o ...func(*MLGetCalendarsRequest)) (*Response, error)

type MLGetCalendarsRequest struct {
	Body io.Reader

	CalendarID string

	From *int
	Size *int

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r MLGetCalendarsRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f MLGetCalendars) WithContext(v context.Context) func(*MLGetCalendarsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLGetCalendars) WithBody(v io.Reader) func(*MLGetCalendarsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLGetCalendars) WithCalendarID(v string) func(*MLGetCalendarsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLGetCalendars) WithFrom(v int) func(*MLGetCalendarsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLGetCalendars) WithSize(v int) func(*MLGetCalendarsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLGetCalendars) WithPretty() func(*MLGetCalendarsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLGetCalendars) WithHuman() func(*MLGetCalendarsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLGetCalendars) WithErrorTrace() func(*MLGetCalendarsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLGetCalendars) WithFilterPath(v ...string) func(*MLGetCalendarsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLGetCalendars) WithHeader(h map[string]string) func(*MLGetCalendarsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLGetCalendars) WithOpaqueID(s string) func(*MLGetCalendarsRequest) {
	_ = "STUB: not implemented"
	return nil
}
