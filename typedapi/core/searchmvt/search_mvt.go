package searchmvt

import (
	gobytes "bytes"
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/gridaggregationtype"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/gridtype"
)

const (
	indexMask = iota + 1

	fieldMask

	zoomMask

	xMask

	yMask
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type SearchMvt struct {
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
	field string
	zoom  string
	x     string
	y     string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewSearchMvt func(index, field, zoom, x, y string) *SearchMvt

func NewSearchMvtFunc(tp elastictransport.Interface) NewSearchMvt {
	_ = "STUB: not implemented"
	return *new(NewSearchMvt)
}

func New(tp elastictransport.Interface) *SearchMvt { _ = "STUB: not implemented"; return nil }

func (r *SearchMvt) Raw(raw io.Reader) *SearchMvt { _ = "STUB: not implemented"; return nil }

func (r *SearchMvt) Request(req *Request) *SearchMvt { _ = "STUB: not implemented"; return nil }

func (r *SearchMvt) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r SearchMvt) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r SearchMvt) Do(providedCtx context.Context) (Response, error) {
	_ = "STUB: not implemented"
	return *new(Response), nil
}

func (r *SearchMvt) Header(key, value string) *SearchMvt { _ = "STUB: not implemented"; return nil }

func (r *SearchMvt) _index(index string) *SearchMvt { _ = "STUB: not implemented"; return nil }

func (r *SearchMvt) _field(field string) *SearchMvt { _ = "STUB: not implemented"; return nil }

func (r *SearchMvt) _zoom(zoom string) *SearchMvt { _ = "STUB: not implemented"; return nil }

func (r *SearchMvt) _x(x string) *SearchMvt { _ = "STUB: not implemented"; return nil }

func (r *SearchMvt) _y(y string) *SearchMvt { _ = "STUB: not implemented"; return nil }

func (r *SearchMvt) ErrorTrace(errortrace bool) *SearchMvt { _ = "STUB: not implemented"; return nil }

func (r *SearchMvt) FilterPath(filterpaths ...string) *SearchMvt {
	_ = "STUB: not implemented"
	return nil
}

func (r *SearchMvt) Human(human bool) *SearchMvt { _ = "STUB: not implemented"; return nil }

func (r *SearchMvt) Pretty(pretty bool) *SearchMvt { _ = "STUB: not implemented"; return nil }

func (r *SearchMvt) Aggs(aggs map[string]types.Aggregations) *SearchMvt {
	_ = "STUB: not implemented"
	return nil
}

func (r *SearchMvt) AddAgg(key string, value types.AggregationsVariant) *SearchMvt {
	_ = "STUB: not implemented"
	return nil
}

func (r *SearchMvt) Buffer(buffer int) *SearchMvt { _ = "STUB: not implemented"; return nil }

func (r *SearchMvt) ExactBounds(exactbounds bool) *SearchMvt { _ = "STUB: not implemented"; return nil }

func (r *SearchMvt) Extent(extent int) *SearchMvt { _ = "STUB: not implemented"; return nil }

func (r *SearchMvt) Fields(fields ...string) *SearchMvt { _ = "STUB: not implemented"; return nil }

func (r *SearchMvt) GridAgg(gridagg gridaggregationtype.GridAggregationType) *SearchMvt {
	_ = "STUB: not implemented"
	return nil
}

func (r *SearchMvt) GridPrecision(gridprecision int) *SearchMvt {
	_ = "STUB: not implemented"
	return nil
}

func (r *SearchMvt) GridType(gridtype gridtype.GridType) *SearchMvt {
	_ = "STUB: not implemented"
	return nil
}

func (r *SearchMvt) ProjectRouting(projectrouting string) *SearchMvt {
	_ = "STUB: not implemented"
	return nil
}

func (r *SearchMvt) Query(query types.QueryVariant) *SearchMvt {
	_ = "STUB: not implemented"
	return nil
}

func (r *SearchMvt) RuntimeMappings(runtimefields types.RuntimeFieldsVariant) *SearchMvt {
	_ = "STUB: not implemented"
	return nil
}

func (r *SearchMvt) Size(size int) *SearchMvt { _ = "STUB: not implemented"; return nil }

func (r *SearchMvt) Sort(sorts ...types.SortCombinationsVariant) *SearchMvt {
	_ = "STUB: not implemented"
	return nil
}

func (r *SearchMvt) SortValues(sortvalues []types.SortCombinations) *SearchMvt {
	_ = "STUB: not implemented"
	return nil
}

func (r *SearchMvt) TrackTotalHits(trackhits types.TrackHitsVariant) *SearchMvt {
	_ = "STUB: not implemented"
	return nil
}

func (r *SearchMvt) WithLabels(withlabels bool) *SearchMvt { _ = "STUB: not implemented"; return nil }
