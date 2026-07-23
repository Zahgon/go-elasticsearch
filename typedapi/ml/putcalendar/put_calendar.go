package putcalendar

import (
	gobytes "bytes"
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

type PutCalendar struct {
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

type NewPutCalendar func(calendarid string) *PutCalendar

func NewPutCalendarFunc(tp elastictransport.Interface) NewPutCalendar {
	_ = "STUB: not implemented"
	return *new(NewPutCalendar)
}

func New(tp elastictransport.Interface) *PutCalendar { _ = "STUB: not implemented"; return nil }

func (r *PutCalendar) Raw(raw io.Reader) *PutCalendar { _ = "STUB: not implemented"; return nil }

func (r *PutCalendar) Request(req *Request) *PutCalendar { _ = "STUB: not implemented"; return nil }

func (r *PutCalendar) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r PutCalendar) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r PutCalendar) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *PutCalendar) Header(key, value string) *PutCalendar { _ = "STUB: not implemented"; return nil }

func (r *PutCalendar) _calendarid(calendarid string) *PutCalendar {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutCalendar) ErrorTrace(errortrace bool) *PutCalendar {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutCalendar) FilterPath(filterpaths ...string) *PutCalendar {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutCalendar) Human(human bool) *PutCalendar { _ = "STUB: not implemented"; return nil }

func (r *PutCalendar) Pretty(pretty bool) *PutCalendar { _ = "STUB: not implemented"; return nil }

func (r *PutCalendar) Description(description string) *PutCalendar {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutCalendar) JobIds(jobids ...string) *PutCalendar { _ = "STUB: not implemented"; return nil }
