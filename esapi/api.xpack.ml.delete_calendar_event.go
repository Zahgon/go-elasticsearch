package esapi

import (
	"context"
	"net/http"
)

func newMLDeleteCalendarEventFunc(t Transport) MLDeleteCalendarEvent {
	_ = "STUB: not implemented"
	return *new(MLDeleteCalendarEvent)
}

type MLDeleteCalendarEvent func(calendar_id string, event_id string, o ...func(*MLDeleteCalendarEventRequest)) (*Response, error)

type MLDeleteCalendarEventRequest struct {
	CalendarID string
	EventID    string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r MLDeleteCalendarEventRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f MLDeleteCalendarEvent) WithContext(v context.Context) func(*MLDeleteCalendarEventRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLDeleteCalendarEvent) WithPretty() func(*MLDeleteCalendarEventRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLDeleteCalendarEvent) WithHuman() func(*MLDeleteCalendarEventRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLDeleteCalendarEvent) WithErrorTrace() func(*MLDeleteCalendarEventRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLDeleteCalendarEvent) WithFilterPath(v ...string) func(*MLDeleteCalendarEventRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLDeleteCalendarEvent) WithHeader(h map[string]string) func(*MLDeleteCalendarEventRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLDeleteCalendarEvent) WithOpaqueID(s string) func(*MLDeleteCalendarEventRequest) {
	_ = "STUB: not implemented"
	return nil
}
