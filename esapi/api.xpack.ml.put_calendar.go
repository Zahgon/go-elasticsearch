package esapi

import (
	"context"
	"io"
	"net/http"
)

func newMLPutCalendarFunc(t Transport) MLPutCalendar {
	_ = "STUB: not implemented"
	return *new(MLPutCalendar)
}

type MLPutCalendar func(calendar_id string, o ...func(*MLPutCalendarRequest)) (*Response, error)

type MLPutCalendarRequest struct {
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

func (r MLPutCalendarRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f MLPutCalendar) WithContext(v context.Context) func(*MLPutCalendarRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLPutCalendar) WithBody(v io.Reader) func(*MLPutCalendarRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLPutCalendar) WithPretty() func(*MLPutCalendarRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLPutCalendar) WithHuman() func(*MLPutCalendarRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLPutCalendar) WithErrorTrace() func(*MLPutCalendarRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLPutCalendar) WithFilterPath(v ...string) func(*MLPutCalendarRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLPutCalendar) WithHeader(h map[string]string) func(*MLPutCalendarRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLPutCalendar) WithOpaqueID(s string) func(*MLPutCalendarRequest) {
	_ = "STUB: not implemented"
	return nil
}
