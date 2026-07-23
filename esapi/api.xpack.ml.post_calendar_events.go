package esapi

import (
	"context"
	"io"
	"net/http"
)

func newMLPostCalendarEventsFunc(t Transport) MLPostCalendarEvents {
	_ = "STUB: not implemented"
	return *new(MLPostCalendarEvents)
}

type MLPostCalendarEvents func(calendar_id string, body io.Reader, o ...func(*MLPostCalendarEventsRequest)) (*Response, error)

type MLPostCalendarEventsRequest struct {
	Body io.Reader

	CalendarID string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r MLPostCalendarEventsRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f MLPostCalendarEvents) WithContext(v context.Context) func(*MLPostCalendarEventsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLPostCalendarEvents) WithPretty() func(*MLPostCalendarEventsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLPostCalendarEvents) WithHuman() func(*MLPostCalendarEventsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLPostCalendarEvents) WithErrorTrace() func(*MLPostCalendarEventsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLPostCalendarEvents) WithFilterPath(v ...string) func(*MLPostCalendarEventsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLPostCalendarEvents) WithHeader(h map[string]string) func(*MLPostCalendarEventsRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLPostCalendarEvents) WithOpaqueID(s string) func(*MLPostCalendarEventsRequest) {
	_ = "STUB: not implemented"
	return nil
}
