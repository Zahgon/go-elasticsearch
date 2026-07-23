package putjob

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
	idMask = iota + 1
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type PutJob struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	req      *Request
	deferred []func(request *Request) error
	buf      *gobytes.Buffer

	paramSet int

	id string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewPutJob func(id string) *PutJob

func NewPutJobFunc(tp elastictransport.Interface) NewPutJob {
	_ = "STUB: not implemented"
	return *new(NewPutJob)
}

func New(tp elastictransport.Interface) *PutJob { _ = "STUB: not implemented"; return nil }

func (r *PutJob) Raw(raw io.Reader) *PutJob { _ = "STUB: not implemented"; return nil }

func (r *PutJob) Request(req *Request) *PutJob { _ = "STUB: not implemented"; return nil }

func (r *PutJob) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r PutJob) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r PutJob) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *PutJob) Header(key, value string) *PutJob { _ = "STUB: not implemented"; return nil }

func (r *PutJob) _id(id string) *PutJob { _ = "STUB: not implemented"; return nil }

func (r *PutJob) ErrorTrace(errortrace bool) *PutJob { _ = "STUB: not implemented"; return nil }

func (r *PutJob) FilterPath(filterpaths ...string) *PutJob { _ = "STUB: not implemented"; return nil }

func (r *PutJob) Human(human bool) *PutJob { _ = "STUB: not implemented"; return nil }

func (r *PutJob) Pretty(pretty bool) *PutJob { _ = "STUB: not implemented"; return nil }

func (r *PutJob) Cron(cron string) *PutJob { _ = "STUB: not implemented"; return nil }

func (r *PutJob) Groups(groups types.GroupingsVariant) *PutJob {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutJob) Headers(httpheaders types.HttpHeadersVariant) *PutJob {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutJob) IndexPattern(indexpattern string) *PutJob { _ = "STUB: not implemented"; return nil }

func (r *PutJob) Metrics(metrics ...types.FieldMetricVariant) *PutJob {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutJob) MetricsValues(metricsvalues []types.FieldMetric) *PutJob {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutJob) PageSize(pagesize int) *PutJob { _ = "STUB: not implemented"; return nil }

func (r *PutJob) RollupIndex(indexname string) *PutJob { _ = "STUB: not implemented"; return nil }

func (r *PutJob) Timeout(duration types.DurationVariant) *PutJob {
	_ = "STUB: not implemented"
	return nil
}
