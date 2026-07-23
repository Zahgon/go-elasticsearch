package getjobs

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
)

const (
	idMask = iota + 1
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type GetJobs struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	paramSet int

	id string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewGetJobs func() *GetJobs

func NewGetJobsFunc(tp elastictransport.Interface) NewGetJobs {
	_ = "STUB: not implemented"
	return *new(NewGetJobs)
}

func New(tp elastictransport.Interface) *GetJobs { _ = "STUB: not implemented"; return nil }

func (r *GetJobs) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetJobs) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetJobs) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetJobs) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *GetJobs) Header(key, value string) *GetJobs { _ = "STUB: not implemented"; return nil }

func (r *GetJobs) Id(id string) *GetJobs { _ = "STUB: not implemented"; return nil }

func (r *GetJobs) ErrorTrace(errortrace bool) *GetJobs { _ = "STUB: not implemented"; return nil }

func (r *GetJobs) FilterPath(filterpaths ...string) *GetJobs { _ = "STUB: not implemented"; return nil }

func (r *GetJobs) Human(human bool) *GetJobs { _ = "STUB: not implemented"; return nil }

func (r *GetJobs) Pretty(pretty bool) *GetJobs { _ = "STUB: not implemented"; return nil }
