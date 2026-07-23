package rankeval

import (
	gobytes "bytes"
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/expandwildcard"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/searchtype"
)

const (
	indexMask = iota + 1
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type RankEval struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	req      *Request
	deferred []func(request *Request) error
	buf      *gobytes.Buffer

	paramSet int

	index string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewRankEval func() *RankEval

func NewRankEvalFunc(tp elastictransport.Interface) NewRankEval {
	_ = "STUB: not implemented"
	return *new(NewRankEval)
}

func New(tp elastictransport.Interface) *RankEval { _ = "STUB: not implemented"; return nil }

func (r *RankEval) Raw(raw io.Reader) *RankEval { _ = "STUB: not implemented"; return nil }

func (r *RankEval) Request(req *Request) *RankEval { _ = "STUB: not implemented"; return nil }

func (r *RankEval) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r RankEval) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r RankEval) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *RankEval) Header(key, value string) *RankEval { _ = "STUB: not implemented"; return nil }

func (r *RankEval) Index(index string) *RankEval { _ = "STUB: not implemented"; return nil }

func (r *RankEval) AllowNoIndices(allownoindices bool) *RankEval {
	_ = "STUB: not implemented"
	return nil
}

func (r *RankEval) ExpandWildcards(expandwildcards ...expandwildcard.ExpandWildcard) *RankEval {
	_ = "STUB: not implemented"
	return nil
}

func (r *RankEval) IgnoreUnavailable(ignoreunavailable bool) *RankEval {
	_ = "STUB: not implemented"
	return nil
}

func (r *RankEval) SearchType(searchtype searchtype.SearchType) *RankEval {
	_ = "STUB: not implemented"
	return nil
}

func (r *RankEval) ErrorTrace(errortrace bool) *RankEval { _ = "STUB: not implemented"; return nil }

func (r *RankEval) FilterPath(filterpaths ...string) *RankEval {
	_ = "STUB: not implemented"
	return nil
}

func (r *RankEval) Human(human bool) *RankEval { _ = "STUB: not implemented"; return nil }

func (r *RankEval) Pretty(pretty bool) *RankEval { _ = "STUB: not implemented"; return nil }

func (r *RankEval) Metric(metric types.RankEvalMetricVariant) *RankEval {
	_ = "STUB: not implemented"
	return nil
}

func (r *RankEval) Requests(requests ...types.RankEvalRequestItemVariant) *RankEval {
	_ = "STUB: not implemented"
	return nil
}

func (r *RankEval) RequestsValues(requestsvalues []types.RankEvalRequestItem) *RankEval {
	_ = "STUB: not implemented"
	return nil
}
