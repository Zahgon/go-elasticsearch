package startjob

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

type StartJob struct {
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

type NewStartJob func(id string) *StartJob

func NewStartJobFunc(tp elastictransport.Interface) NewStartJob {
	_ = "STUB: not implemented"
	return *new(NewStartJob)
}

func New(tp elastictransport.Interface) *StartJob { _ = "STUB: not implemented"; return nil }

func (r *StartJob) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r StartJob) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r StartJob) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r StartJob) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *StartJob) Header(key, value string) *StartJob { _ = "STUB: not implemented"; return nil }

func (r *StartJob) _id(id string) *StartJob { _ = "STUB: not implemented"; return nil }

func (r *StartJob) ErrorTrace(errortrace bool) *StartJob { _ = "STUB: not implemented"; return nil }

func (r *StartJob) FilterPath(filterpaths ...string) *StartJob {
	_ = "STUB: not implemented"
	return nil
}

func (r *StartJob) Human(human bool) *StartJob { _ = "STUB: not implemented"; return nil }

func (r *StartJob) Pretty(pretty bool) *StartJob { _ = "STUB: not implemented"; return nil }
