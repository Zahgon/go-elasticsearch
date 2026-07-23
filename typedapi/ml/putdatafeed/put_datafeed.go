package putdatafeed

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
)

const (
	datafeedidMask = iota + 1
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type PutDatafeed struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	req      *Request
	deferred []func(request *Request) error
	buf      *gobytes.Buffer

	paramSet int

	datafeedid string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewPutDatafeed func(datafeedid string) *PutDatafeed

func NewPutDatafeedFunc(tp elastictransport.Interface) NewPutDatafeed {
	_ = "STUB: not implemented"
	return *new(NewPutDatafeed)
}

func New(tp elastictransport.Interface) *PutDatafeed { _ = "STUB: not implemented"; return nil }

func (r *PutDatafeed) Raw(raw io.Reader) *PutDatafeed { _ = "STUB: not implemented"; return nil }

func (r *PutDatafeed) Request(req *Request) *PutDatafeed { _ = "STUB: not implemented"; return nil }

func (r *PutDatafeed) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r PutDatafeed) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r PutDatafeed) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *PutDatafeed) Header(key, value string) *PutDatafeed { _ = "STUB: not implemented"; return nil }

func (r *PutDatafeed) _datafeedid(datafeedid string) *PutDatafeed {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutDatafeed) AllowNoIndices(allownoindices bool) *PutDatafeed {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutDatafeed) ExpandWildcards(expandwildcards ...expandwildcard.ExpandWildcard) *PutDatafeed {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutDatafeed) IgnoreThrottled(ignorethrottled bool) *PutDatafeed {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutDatafeed) IgnoreUnavailable(ignoreunavailable bool) *PutDatafeed {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutDatafeed) ErrorTrace(errortrace bool) *PutDatafeed {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutDatafeed) FilterPath(filterpaths ...string) *PutDatafeed {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutDatafeed) Human(human bool) *PutDatafeed { _ = "STUB: not implemented"; return nil }

func (r *PutDatafeed) Pretty(pretty bool) *PutDatafeed { _ = "STUB: not implemented"; return nil }

func (r *PutDatafeed) Aggregations(aggregations map[string]types.Aggregations) *PutDatafeed {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutDatafeed) AddAggregation(key string, value types.AggregationsVariant) *PutDatafeed {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutDatafeed) ChunkingConfig(chunkingconfig types.ChunkingConfigVariant) *PutDatafeed {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutDatafeed) DelayedDataCheckConfig(delayeddatacheckconfig types.DelayedDataCheckConfigVariant) *PutDatafeed {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutDatafeed) Frequency(duration types.DurationVariant) *PutDatafeed {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutDatafeed) Headers(httpheaders types.HttpHeadersVariant) *PutDatafeed {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutDatafeed) Indices(indices ...string) *PutDatafeed {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutDatafeed) IndicesOptions(indicesoptions types.IndicesOptionsVariant) *PutDatafeed {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutDatafeed) JobId(id string) *PutDatafeed { _ = "STUB: not implemented"; return nil }

func (r *PutDatafeed) MaxEmptySearches(maxemptysearches int) *PutDatafeed {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutDatafeed) Query(query types.QueryVariant) *PutDatafeed {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutDatafeed) QueryDelay(duration types.DurationVariant) *PutDatafeed {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutDatafeed) RuntimeMappings(runtimefields types.RuntimeFieldsVariant) *PutDatafeed {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutDatafeed) ScriptFields(scriptfields map[string]types.ScriptField) *PutDatafeed {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutDatafeed) AddScriptField(key string, value types.ScriptFieldVariant) *PutDatafeed {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutDatafeed) ScrollSize(scrollsize int) *PutDatafeed {
	_ = "STUB: not implemented"
	return nil
}
