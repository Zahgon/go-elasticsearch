package putcalendarjob

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

type PutCalendarJob struct {
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

type NewPutCalendarJob func(calendarid, jobid string) *PutCalendarJob

func NewPutCalendarJobFunc(tp elastictransport.Interface) NewPutCalendarJob {
	_ = "STUB: not implemented"
	return *new(NewPutCalendarJob)
}

func New(tp elastictransport.Interface) *PutCalendarJob { _ = "STUB: not implemented"; return nil }

func (r *PutCalendarJob) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r PutCalendarJob) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r PutCalendarJob) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r PutCalendarJob) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *PutCalendarJob) Header(key, value string) *PutCalendarJob {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutCalendarJob) _calendarid(calendarid string) *PutCalendarJob {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutCalendarJob) _jobid(jobid string) *PutCalendarJob {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutCalendarJob) ErrorTrace(errortrace bool) *PutCalendarJob {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutCalendarJob) FilterPath(filterpaths ...string) *PutCalendarJob {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutCalendarJob) Human(human bool) *PutCalendarJob { _ = "STUB: not implemented"; return nil }

func (r *PutCalendarJob) Pretty(pretty bool) *PutCalendarJob { _ = "STUB: not implemented"; return nil }
