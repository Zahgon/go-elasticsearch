package rollupsearch

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
	indexMask = iota + 1
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type RollupSearch struct {
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

type NewRollupSearch func(index string) *RollupSearch

func NewRollupSearchFunc(tp elastictransport.Interface) NewRollupSearch {
	_ = "STUB: not implemented"
	return *new(NewRollupSearch)
}

func New(tp elastictransport.Interface) *RollupSearch { _ = "STUB: not implemented"; return nil }

func (r *RollupSearch) Raw(raw io.Reader) *RollupSearch { _ = "STUB: not implemented"; return nil }

func (r *RollupSearch) Request(req *Request) *RollupSearch { _ = "STUB: not implemented"; return nil }

func (r *RollupSearch) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r RollupSearch) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r RollupSearch) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *RollupSearch) Header(key, value string) *RollupSearch {
	_ = "STUB: not implemented"
	return nil
}

func (r *RollupSearch) _index(index string) *RollupSearch { _ = "STUB: not implemented"; return nil }

func (r *RollupSearch) RestTotalHitsAsInt(resttotalhitsasint bool) *RollupSearch {
	_ = "STUB: not implemented"
	return nil
}

func (r *RollupSearch) TypedKeys(typedkeys bool) *RollupSearch {
	_ = "STUB: not implemented"
	return nil
}

func (r *RollupSearch) ErrorTrace(errortrace bool) *RollupSearch {
	_ = "STUB: not implemented"
	return nil
}

func (r *RollupSearch) FilterPath(filterpaths ...string) *RollupSearch {
	_ = "STUB: not implemented"
	return nil
}

func (r *RollupSearch) Human(human bool) *RollupSearch { _ = "STUB: not implemented"; return nil }

func (r *RollupSearch) Pretty(pretty bool) *RollupSearch { _ = "STUB: not implemented"; return nil }

func (r *RollupSearch) Aggregations(aggregations map[string]types.Aggregations) *RollupSearch {
	_ = "STUB: not implemented"
	return nil
}

func (r *RollupSearch) AddAggregation(key string, value types.AggregationsVariant) *RollupSearch {
	_ = "STUB: not implemented"
	return nil
}

func (r *RollupSearch) Query(query types.QueryVariant) *RollupSearch {
	_ = "STUB: not implemented"
	return nil
}

func (r *RollupSearch) Size(size int) *RollupSearch { _ = "STUB: not implemented"; return nil }
