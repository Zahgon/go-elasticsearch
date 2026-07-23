package getoverallbuckets

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

type GetOverallBuckets struct {
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

type NewGetOverallBuckets func(jobid string) *GetOverallBuckets

func NewGetOverallBucketsFunc(tp elastictransport.Interface) NewGetOverallBuckets {
	_ = "STUB: not implemented"
	return *new(NewGetOverallBuckets)
}

func New(tp elastictransport.Interface) *GetOverallBuckets { _ = "STUB: not implemented"; return nil }

func (r *GetOverallBuckets) Raw(raw io.Reader) *GetOverallBuckets {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetOverallBuckets) Request(req *Request) *GetOverallBuckets {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetOverallBuckets) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetOverallBuckets) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetOverallBuckets) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *GetOverallBuckets) Header(key, value string) *GetOverallBuckets {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetOverallBuckets) _jobid(jobid string) *GetOverallBuckets {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetOverallBuckets) ErrorTrace(errortrace bool) *GetOverallBuckets {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetOverallBuckets) FilterPath(filterpaths ...string) *GetOverallBuckets {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetOverallBuckets) Human(human bool) *GetOverallBuckets {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetOverallBuckets) Pretty(pretty bool) *GetOverallBuckets {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetOverallBuckets) AllowNoMatch(allownomatch bool) *GetOverallBuckets {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetOverallBuckets) BucketSpan(duration types.DurationVariant) *GetOverallBuckets {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetOverallBuckets) End(datetime types.DateTimeVariant) *GetOverallBuckets {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetOverallBuckets) ExcludeInterim(excludeinterim bool) *GetOverallBuckets {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetOverallBuckets) OverallScore(overallscore types.Float64) *GetOverallBuckets {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetOverallBuckets) Start(datetime types.DateTimeVariant) *GetOverallBuckets {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetOverallBuckets) TopN(topn int) *GetOverallBuckets {
	_ = "STUB: not implemented"
	return nil
}
