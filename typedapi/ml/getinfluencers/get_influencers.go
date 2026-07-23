package getinfluencers

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

type GetInfluencers struct {
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

type NewGetInfluencers func(jobid string) *GetInfluencers

func NewGetInfluencersFunc(tp elastictransport.Interface) NewGetInfluencers {
	_ = "STUB: not implemented"
	return *new(NewGetInfluencers)
}

func New(tp elastictransport.Interface) *GetInfluencers { _ = "STUB: not implemented"; return nil }

func (r *GetInfluencers) Raw(raw io.Reader) *GetInfluencers { _ = "STUB: not implemented"; return nil }

func (r *GetInfluencers) Request(req *Request) *GetInfluencers {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetInfluencers) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetInfluencers) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetInfluencers) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *GetInfluencers) Header(key, value string) *GetInfluencers {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetInfluencers) _jobid(jobid string) *GetInfluencers {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetInfluencers) Desc(desc bool) *GetInfluencers { _ = "STUB: not implemented"; return nil }

func (r *GetInfluencers) End(datetime string) *GetInfluencers {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetInfluencers) ExcludeInterim(excludeinterim bool) *GetInfluencers {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetInfluencers) InfluencerScore(influencerscore string) *GetInfluencers {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetInfluencers) From(from int) *GetInfluencers { _ = "STUB: not implemented"; return nil }

func (r *GetInfluencers) Size(size int) *GetInfluencers { _ = "STUB: not implemented"; return nil }

func (r *GetInfluencers) Sort(field string) *GetInfluencers { _ = "STUB: not implemented"; return nil }

func (r *GetInfluencers) Start(datetime string) *GetInfluencers {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetInfluencers) ErrorTrace(errortrace bool) *GetInfluencers {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetInfluencers) FilterPath(filterpaths ...string) *GetInfluencers {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetInfluencers) Human(human bool) *GetInfluencers { _ = "STUB: not implemented"; return nil }

func (r *GetInfluencers) Pretty(pretty bool) *GetInfluencers { _ = "STUB: not implemented"; return nil }

func (r *GetInfluencers) Page(page types.PageVariant) *GetInfluencers {
	_ = "STUB: not implemented"
	return nil
}
