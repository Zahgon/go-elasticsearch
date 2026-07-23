package getbuckets

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

	timestampMask
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type GetBuckets struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	req      *Request
	deferred []func(request *Request) error
	buf      *gobytes.Buffer

	paramSet int

	jobid     string
	timestamp string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewGetBuckets func(jobid string) *GetBuckets

func NewGetBucketsFunc(tp elastictransport.Interface) NewGetBuckets {
	_ = "STUB: not implemented"
	return *new(NewGetBuckets)
}

func New(tp elastictransport.Interface) *GetBuckets { _ = "STUB: not implemented"; return nil }

func (r *GetBuckets) Raw(raw io.Reader) *GetBuckets { _ = "STUB: not implemented"; return nil }

func (r *GetBuckets) Request(req *Request) *GetBuckets { _ = "STUB: not implemented"; return nil }

func (r *GetBuckets) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetBuckets) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetBuckets) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *GetBuckets) Header(key, value string) *GetBuckets { _ = "STUB: not implemented"; return nil }

func (r *GetBuckets) _jobid(jobid string) *GetBuckets { _ = "STUB: not implemented"; return nil }

func (r *GetBuckets) Timestamp(timestamp string) *GetBuckets { _ = "STUB: not implemented"; return nil }

func (r *GetBuckets) From(from int) *GetBuckets { _ = "STUB: not implemented"; return nil }

func (r *GetBuckets) Size(size int) *GetBuckets { _ = "STUB: not implemented"; return nil }

func (r *GetBuckets) ErrorTrace(errortrace bool) *GetBuckets { _ = "STUB: not implemented"; return nil }

func (r *GetBuckets) FilterPath(filterpaths ...string) *GetBuckets {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetBuckets) Human(human bool) *GetBuckets { _ = "STUB: not implemented"; return nil }

func (r *GetBuckets) Pretty(pretty bool) *GetBuckets { _ = "STUB: not implemented"; return nil }

func (r *GetBuckets) AnomalyScore(anomalyscore types.Float64) *GetBuckets {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetBuckets) Desc(desc bool) *GetBuckets { _ = "STUB: not implemented"; return nil }

func (r *GetBuckets) End(datetime types.DateTimeVariant) *GetBuckets {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetBuckets) ExcludeInterim(excludeinterim bool) *GetBuckets {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetBuckets) Expand(expand bool) *GetBuckets { _ = "STUB: not implemented"; return nil }

func (r *GetBuckets) Page(page types.PageVariant) *GetBuckets {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetBuckets) Sort(field string) *GetBuckets { _ = "STUB: not implemented"; return nil }

func (r *GetBuckets) Start(datetime types.DateTimeVariant) *GetBuckets {
	_ = "STUB: not implemented"
	return nil
}
