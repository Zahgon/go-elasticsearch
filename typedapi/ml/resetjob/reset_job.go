package resetjob

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
)

const (
	jobidMask = iota + 1
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type ResetJob struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	paramSet int

	jobid string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewResetJob func(jobid string) *ResetJob

func NewResetJobFunc(tp elastictransport.Interface) NewResetJob {
	_ = "STUB: not implemented"
	return *new(NewResetJob)
}

func New(tp elastictransport.Interface) *ResetJob { _ = "STUB: not implemented"; return nil }

func (r *ResetJob) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r ResetJob) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r ResetJob) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r ResetJob) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *ResetJob) Header(key, value string) *ResetJob { _ = "STUB: not implemented"; return nil }

func (r *ResetJob) _jobid(jobid string) *ResetJob { _ = "STUB: not implemented"; return nil }

func (r *ResetJob) WaitForCompletion(waitforcompletion bool) *ResetJob {
	_ = "STUB: not implemented"
	return nil
}

func (r *ResetJob) DeleteUserAnnotations(deleteuserannotations bool) *ResetJob {
	_ = "STUB: not implemented"
	return nil
}

func (r *ResetJob) ErrorTrace(errortrace bool) *ResetJob { _ = "STUB: not implemented"; return nil }

func (r *ResetJob) FilterPath(filterpaths ...string) *ResetJob {
	_ = "STUB: not implemented"
	return nil
}

func (r *ResetJob) Human(human bool) *ResetJob { _ = "STUB: not implemented"; return nil }

func (r *ResetJob) Pretty(pretty bool) *ResetJob { _ = "STUB: not implemented"; return nil }
