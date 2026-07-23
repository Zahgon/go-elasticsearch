package stopjob

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

type StopJob struct {
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

type NewStopJob func(id string) *StopJob

func NewStopJobFunc(tp elastictransport.Interface) NewStopJob {
	_ = "STUB: not implemented"
	return *new(NewStopJob)
}

func New(tp elastictransport.Interface) *StopJob { _ = "STUB: not implemented"; return nil }

func (r *StopJob) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r StopJob) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r StopJob) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r StopJob) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *StopJob) Header(key, value string) *StopJob { _ = "STUB: not implemented"; return nil }

func (r *StopJob) _id(id string) *StopJob { _ = "STUB: not implemented"; return nil }

func (r *StopJob) Timeout(duration string) *StopJob { _ = "STUB: not implemented"; return nil }

func (r *StopJob) WaitForCompletion(waitforcompletion bool) *StopJob {
	_ = "STUB: not implemented"
	return nil
}

func (r *StopJob) ErrorTrace(errortrace bool) *StopJob { _ = "STUB: not implemented"; return nil }

func (r *StopJob) FilterPath(filterpaths ...string) *StopJob { _ = "STUB: not implemented"; return nil }

func (r *StopJob) Human(human bool) *StopJob { _ = "STUB: not implemented"; return nil }

func (r *StopJob) Pretty(pretty bool) *StopJob { _ = "STUB: not implemented"; return nil }
