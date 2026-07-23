package deletecalendar

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

type DeleteCalendar struct {
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

type NewDeleteCalendar func(calendarid string) *DeleteCalendar

func NewDeleteCalendarFunc(tp elastictransport.Interface) NewDeleteCalendar {
	_ = "STUB: not implemented"
	return *new(NewDeleteCalendar)
}

func New(tp elastictransport.Interface) *DeleteCalendar { _ = "STUB: not implemented"; return nil }

func (r *DeleteCalendar) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r DeleteCalendar) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r DeleteCalendar) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r DeleteCalendar) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *DeleteCalendar) Header(key, value string) *DeleteCalendar {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteCalendar) _calendarid(calendarid string) *DeleteCalendar {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteCalendar) ErrorTrace(errortrace bool) *DeleteCalendar {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteCalendar) FilterPath(filterpaths ...string) *DeleteCalendar {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteCalendar) Human(human bool) *DeleteCalendar { _ = "STUB: not implemented"; return nil }

func (r *DeleteCalendar) Pretty(pretty bool) *DeleteCalendar { _ = "STUB: not implemented"; return nil }
