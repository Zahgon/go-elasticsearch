package closejob

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
	jobidMask = iota + 1
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type CloseJob struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	req      *Request
	deferred []func(request *Request) error
	buf      *gobytes.Buffer

	paramSet int

	jobid string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewCloseJob func(jobid string) *CloseJob

func NewCloseJobFunc(tp elastictransport.Interface) NewCloseJob {
	_ = "STUB: not implemented"
	return *new(NewCloseJob)
}

func New(tp elastictransport.Interface) *CloseJob { _ = "STUB: not implemented"; return nil }

func (r *CloseJob) Raw(raw io.Reader) *CloseJob { _ = "STUB: not implemented"; return nil }

func (r *CloseJob) Request(req *Request) *CloseJob { _ = "STUB: not implemented"; return nil }

func (r *CloseJob) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r CloseJob) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r CloseJob) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *CloseJob) Header(key, value string) *CloseJob { _ = "STUB: not implemented"; return nil }

func (r *CloseJob) _jobid(jobid string) *CloseJob { _ = "STUB: not implemented"; return nil }

func (r *CloseJob) ErrorTrace(errortrace bool) *CloseJob { _ = "STUB: not implemented"; return nil }

func (r *CloseJob) FilterPath(filterpaths ...string) *CloseJob {
	_ = "STUB: not implemented"
	return nil
}

func (r *CloseJob) Human(human bool) *CloseJob { _ = "STUB: not implemented"; return nil }

func (r *CloseJob) Pretty(pretty bool) *CloseJob { _ = "STUB: not implemented"; return nil }

func (r *CloseJob) AllowNoMatch(allownomatch bool) *CloseJob { _ = "STUB: not implemented"; return nil }

func (r *CloseJob) Force(force bool) *CloseJob { _ = "STUB: not implemented"; return nil }

func (r *CloseJob) Timeout(duration types.DurationVariant) *CloseJob {
	_ = "STUB: not implemented"
	return nil
}
