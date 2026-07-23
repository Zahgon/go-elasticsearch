package deletejob

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

type DeleteJob struct {
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

type NewDeleteJob func(id string) *DeleteJob

func NewDeleteJobFunc(tp elastictransport.Interface) NewDeleteJob {
	_ = "STUB: not implemented"
	return *new(NewDeleteJob)
}

func New(tp elastictransport.Interface) *DeleteJob { _ = "STUB: not implemented"; return nil }

func (r *DeleteJob) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r DeleteJob) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r DeleteJob) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r DeleteJob) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *DeleteJob) Header(key, value string) *DeleteJob { _ = "STUB: not implemented"; return nil }

func (r *DeleteJob) _id(id string) *DeleteJob { _ = "STUB: not implemented"; return nil }

func (r *DeleteJob) ErrorTrace(errortrace bool) *DeleteJob { _ = "STUB: not implemented"; return nil }

func (r *DeleteJob) FilterPath(filterpaths ...string) *DeleteJob {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteJob) Human(human bool) *DeleteJob { _ = "STUB: not implemented"; return nil }

func (r *DeleteJob) Pretty(pretty bool) *DeleteJob { _ = "STUB: not implemented"; return nil }
