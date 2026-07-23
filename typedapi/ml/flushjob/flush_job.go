package flushjob

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

type FlushJob struct {
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

type NewFlushJob func(jobid string) *FlushJob

func NewFlushJobFunc(tp elastictransport.Interface) NewFlushJob {
	_ = "STUB: not implemented"
	return *new(NewFlushJob)
}

func New(tp elastictransport.Interface) *FlushJob { _ = "STUB: not implemented"; return nil }

func (r *FlushJob) Raw(raw io.Reader) *FlushJob { _ = "STUB: not implemented"; return nil }

func (r *FlushJob) Request(req *Request) *FlushJob { _ = "STUB: not implemented"; return nil }

func (r *FlushJob) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r FlushJob) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r FlushJob) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *FlushJob) Header(key, value string) *FlushJob { _ = "STUB: not implemented"; return nil }

func (r *FlushJob) _jobid(jobid string) *FlushJob { _ = "STUB: not implemented"; return nil }

func (r *FlushJob) ErrorTrace(errortrace bool) *FlushJob { _ = "STUB: not implemented"; return nil }

func (r *FlushJob) FilterPath(filterpaths ...string) *FlushJob {
	_ = "STUB: not implemented"
	return nil
}

func (r *FlushJob) Human(human bool) *FlushJob { _ = "STUB: not implemented"; return nil }

func (r *FlushJob) Pretty(pretty bool) *FlushJob { _ = "STUB: not implemented"; return nil }

func (r *FlushJob) AdvanceTime(datetime types.DateTimeVariant) *FlushJob {
	_ = "STUB: not implemented"
	return nil
}

func (r *FlushJob) CalcInterim(calcinterim bool) *FlushJob { _ = "STUB: not implemented"; return nil }

func (r *FlushJob) End(datetime types.DateTimeVariant) *FlushJob {
	_ = "STUB: not implemented"
	return nil
}

func (r *FlushJob) SkipTime(datetime types.DateTimeVariant) *FlushJob {
	_ = "STUB: not implemented"
	return nil
}

func (r *FlushJob) Start(datetime types.DateTimeVariant) *FlushJob {
	_ = "STUB: not implemented"
	return nil
}
