package getcalendarevents

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
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type GetCalendarEvents struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	paramSet int

	calendarid string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewGetCalendarEvents func(calendarid string) *GetCalendarEvents

func NewGetCalendarEventsFunc(tp elastictransport.Interface) NewGetCalendarEvents {
	_ = "STUB: not implemented"
	return *new(NewGetCalendarEvents)
}

func New(tp elastictransport.Interface) *GetCalendarEvents { _ = "STUB: not implemented"; return nil }

func (r *GetCalendarEvents) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetCalendarEvents) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetCalendarEvents) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetCalendarEvents) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *GetCalendarEvents) Header(key, value string) *GetCalendarEvents {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetCalendarEvents) _calendarid(calendarid string) *GetCalendarEvents {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetCalendarEvents) End(datetime string) *GetCalendarEvents {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetCalendarEvents) From(from int) *GetCalendarEvents {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetCalendarEvents) JobId(id string) *GetCalendarEvents {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetCalendarEvents) Size(size int) *GetCalendarEvents {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetCalendarEvents) Start(datetime string) *GetCalendarEvents {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetCalendarEvents) ErrorTrace(errortrace bool) *GetCalendarEvents {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetCalendarEvents) FilterPath(filterpaths ...string) *GetCalendarEvents {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetCalendarEvents) Human(human bool) *GetCalendarEvents {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetCalendarEvents) Pretty(pretty bool) *GetCalendarEvents {
	_ = "STUB: not implemented"
	return nil
}
