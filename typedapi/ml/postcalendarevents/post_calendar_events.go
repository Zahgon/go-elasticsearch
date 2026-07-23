package postcalendarevents

import (
	gobytes "bytes"
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

const (
	calendaridMask = iota + 1
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type PostCalendarEvents struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	req      *Request
	deferred []func(request *Request) error
	buf      *gobytes.Buffer

	paramSet int

	calendarid string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewPostCalendarEvents func(calendarid string) *PostCalendarEvents

func NewPostCalendarEventsFunc(tp elastictransport.Interface) NewPostCalendarEvents {
	_ = "STUB: not implemented"
	return *new(NewPostCalendarEvents)
}

func New(tp elastictransport.Interface) *PostCalendarEvents { _ = "STUB: not implemented"; return nil }

func (r *PostCalendarEvents) Raw(raw io.Reader) *PostCalendarEvents {
	_ = "STUB: not implemented"
	return nil
}

func (r *PostCalendarEvents) Request(req *Request) *PostCalendarEvents {
	_ = "STUB: not implemented"
	return nil
}

func (r *PostCalendarEvents) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r PostCalendarEvents) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r PostCalendarEvents) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *PostCalendarEvents) Header(key, value string) *PostCalendarEvents {
	_ = "STUB: not implemented"
	return nil
}

func (r *PostCalendarEvents) _calendarid(calendarid string) *PostCalendarEvents {
	_ = "STUB: not implemented"
	return nil
}

func (r *PostCalendarEvents) ErrorTrace(errortrace bool) *PostCalendarEvents {
	_ = "STUB: not implemented"
	return nil
}

func (r *PostCalendarEvents) FilterPath(filterpaths ...string) *PostCalendarEvents {
	_ = "STUB: not implemented"
	return nil
}

func (r *PostCalendarEvents) Human(human bool) *PostCalendarEvents {
	_ = "STUB: not implemented"
	return nil
}

func (r *PostCalendarEvents) Pretty(pretty bool) *PostCalendarEvents {
	_ = "STUB: not implemented"
	return nil
}

func (r *PostCalendarEvents) Events(events ...types.CalendarEventVariant) *PostCalendarEvents {
	_ = "STUB: not implemented"
	return nil
}

func (r *PostCalendarEvents) EventsValues(eventsvalues []types.CalendarEvent) *PostCalendarEvents {
	_ = "STUB: not implemented"
	return nil
}
