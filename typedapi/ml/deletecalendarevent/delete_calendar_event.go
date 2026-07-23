package deletecalendarevent

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
)

const (
	calendaridMask = iota + 1

	eventidMask
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type DeleteCalendarEvent struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	paramSet int

	calendarid string
	eventid    string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewDeleteCalendarEvent func(calendarid, eventid string) *DeleteCalendarEvent

func NewDeleteCalendarEventFunc(tp elastictransport.Interface) NewDeleteCalendarEvent {
	_ = "STUB: not implemented"
	return *new(NewDeleteCalendarEvent)
}

func New(tp elastictransport.Interface) *DeleteCalendarEvent { _ = "STUB: not implemented"; return nil }

func (r *DeleteCalendarEvent) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r DeleteCalendarEvent) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r DeleteCalendarEvent) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r DeleteCalendarEvent) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *DeleteCalendarEvent) Header(key, value string) *DeleteCalendarEvent {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteCalendarEvent) _calendarid(calendarid string) *DeleteCalendarEvent {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteCalendarEvent) _eventid(eventid string) *DeleteCalendarEvent {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteCalendarEvent) ErrorTrace(errortrace bool) *DeleteCalendarEvent {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteCalendarEvent) FilterPath(filterpaths ...string) *DeleteCalendarEvent {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteCalendarEvent) Human(human bool) *DeleteCalendarEvent {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteCalendarEvent) Pretty(pretty bool) *DeleteCalendarEvent {
	_ = "STUB: not implemented"
	return nil
}
