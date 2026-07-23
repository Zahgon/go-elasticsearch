package openjob

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

type OpenJob struct {
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

type NewOpenJob func(jobid string) *OpenJob

func NewOpenJobFunc(tp elastictransport.Interface) NewOpenJob {
	_ = "STUB: not implemented"
	return *new(NewOpenJob)
}

func New(tp elastictransport.Interface) *OpenJob { _ = "STUB: not implemented"; return nil }

func (r *OpenJob) Raw(raw io.Reader) *OpenJob { _ = "STUB: not implemented"; return nil }

func (r *OpenJob) Request(req *Request) *OpenJob { _ = "STUB: not implemented"; return nil }

func (r *OpenJob) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r OpenJob) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r OpenJob) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *OpenJob) Header(key, value string) *OpenJob { _ = "STUB: not implemented"; return nil }

func (r *OpenJob) _jobid(jobid string) *OpenJob { _ = "STUB: not implemented"; return nil }

func (r *OpenJob) ErrorTrace(errortrace bool) *OpenJob { _ = "STUB: not implemented"; return nil }

func (r *OpenJob) FilterPath(filterpaths ...string) *OpenJob { _ = "STUB: not implemented"; return nil }

func (r *OpenJob) Human(human bool) *OpenJob { _ = "STUB: not implemented"; return nil }

func (r *OpenJob) Pretty(pretty bool) *OpenJob { _ = "STUB: not implemented"; return nil }

func (r *OpenJob) Timeout(duration types.DurationVariant) *OpenJob {
	_ = "STUB: not implemented"
	return nil
}
