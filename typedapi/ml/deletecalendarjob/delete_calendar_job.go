package deletecalendarjob

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

	jobidMask
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type DeleteCalendarJob struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	paramSet int

	calendarid string
	jobid      string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewDeleteCalendarJob func(calendarid, jobid string) *DeleteCalendarJob

func NewDeleteCalendarJobFunc(tp elastictransport.Interface) NewDeleteCalendarJob {
	_ = "STUB: not implemented"
	return *new(NewDeleteCalendarJob)
}

func New(tp elastictransport.Interface) *DeleteCalendarJob { _ = "STUB: not implemented"; return nil }

func (r *DeleteCalendarJob) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r DeleteCalendarJob) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r DeleteCalendarJob) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r DeleteCalendarJob) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *DeleteCalendarJob) Header(key, value string) *DeleteCalendarJob {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteCalendarJob) _calendarid(calendarid string) *DeleteCalendarJob {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteCalendarJob) _jobid(jobid string) *DeleteCalendarJob {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteCalendarJob) ErrorTrace(errortrace bool) *DeleteCalendarJob {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteCalendarJob) FilterPath(filterpaths ...string) *DeleteCalendarJob {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteCalendarJob) Human(human bool) *DeleteCalendarJob {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteCalendarJob) Pretty(pretty bool) *DeleteCalendarJob {
	_ = "STUB: not implemented"
	return nil
}
