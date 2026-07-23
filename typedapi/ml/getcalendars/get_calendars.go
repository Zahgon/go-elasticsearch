package getcalendars

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

type GetCalendars struct {
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

type NewGetCalendars func() *GetCalendars

func NewGetCalendarsFunc(tp elastictransport.Interface) NewGetCalendars {
	_ = "STUB: not implemented"
	return *new(NewGetCalendars)
}

func New(tp elastictransport.Interface) *GetCalendars { _ = "STUB: not implemented"; return nil }

func (r *GetCalendars) Raw(raw io.Reader) *GetCalendars { _ = "STUB: not implemented"; return nil }

func (r *GetCalendars) Request(req *Request) *GetCalendars { _ = "STUB: not implemented"; return nil }

func (r *GetCalendars) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetCalendars) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetCalendars) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *GetCalendars) Header(key, value string) *GetCalendars {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetCalendars) CalendarId(calendarid string) *GetCalendars {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetCalendars) From(from int) *GetCalendars { _ = "STUB: not implemented"; return nil }

func (r *GetCalendars) Size(size int) *GetCalendars { _ = "STUB: not implemented"; return nil }

func (r *GetCalendars) ErrorTrace(errortrace bool) *GetCalendars {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetCalendars) FilterPath(filterpaths ...string) *GetCalendars {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetCalendars) Human(human bool) *GetCalendars { _ = "STUB: not implemented"; return nil }

func (r *GetCalendars) Pretty(pretty bool) *GetCalendars { _ = "STUB: not implemented"; return nil }

func (r *GetCalendars) Page(page types.PageVariant) *GetCalendars {
	_ = "STUB: not implemented"
	return nil
}
