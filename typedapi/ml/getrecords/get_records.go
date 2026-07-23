package getrecords

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

type GetRecords struct {
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

type NewGetRecords func(jobid string) *GetRecords

func NewGetRecordsFunc(tp elastictransport.Interface) NewGetRecords {
	_ = "STUB: not implemented"
	return *new(NewGetRecords)
}

func New(tp elastictransport.Interface) *GetRecords { _ = "STUB: not implemented"; return nil }

func (r *GetRecords) Raw(raw io.Reader) *GetRecords { _ = "STUB: not implemented"; return nil }

func (r *GetRecords) Request(req *Request) *GetRecords { _ = "STUB: not implemented"; return nil }

func (r *GetRecords) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetRecords) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetRecords) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *GetRecords) Header(key, value string) *GetRecords { _ = "STUB: not implemented"; return nil }

func (r *GetRecords) _jobid(jobid string) *GetRecords { _ = "STUB: not implemented"; return nil }

func (r *GetRecords) From(from int) *GetRecords { _ = "STUB: not implemented"; return nil }

func (r *GetRecords) Size(size int) *GetRecords { _ = "STUB: not implemented"; return nil }

func (r *GetRecords) ErrorTrace(errortrace bool) *GetRecords { _ = "STUB: not implemented"; return nil }

func (r *GetRecords) FilterPath(filterpaths ...string) *GetRecords {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetRecords) Human(human bool) *GetRecords { _ = "STUB: not implemented"; return nil }

func (r *GetRecords) Pretty(pretty bool) *GetRecords { _ = "STUB: not implemented"; return nil }

func (r *GetRecords) Desc(desc bool) *GetRecords { _ = "STUB: not implemented"; return nil }

func (r *GetRecords) End(datetime types.DateTimeVariant) *GetRecords {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetRecords) ExcludeInterim(excludeinterim bool) *GetRecords {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetRecords) Page(page types.PageVariant) *GetRecords {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetRecords) RecordScore(recordscore types.Float64) *GetRecords {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetRecords) Sort(field string) *GetRecords { _ = "STUB: not implemented"; return nil }

func (r *GetRecords) Start(datetime types.DateTimeVariant) *GetRecords {
	_ = "STUB: not implemented"
	return nil
}
