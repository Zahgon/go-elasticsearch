package updatedatafeed

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

type UpdateDatafeed struct {
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

type NewUpdateDatafeed func(datafeedid string) *UpdateDatafeed

func NewUpdateDatafeedFunc(tp elastictransport.Interface) NewUpdateDatafeed {
	_ = "STUB: not implemented"
	return *new(NewUpdateDatafeed)
}

func New(tp elastictransport.Interface) *UpdateDatafeed { _ = "STUB: not implemented"; return nil }

func (r *UpdateDatafeed) Raw(raw io.Reader) *UpdateDatafeed { _ = "STUB: not implemented"; return nil }

func (r *UpdateDatafeed) Request(req *Request) *UpdateDatafeed {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateDatafeed) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r UpdateDatafeed) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r UpdateDatafeed) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *UpdateDatafeed) Header(key, value string) *UpdateDatafeed {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateDatafeed) _datafeedid(datafeedid string) *UpdateDatafeed {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateDatafeed) AllowNoIndices(allownoindices bool) *UpdateDatafeed {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateDatafeed) ExpandWildcards(expandwildcards ...expandwildcard.ExpandWildcard) *UpdateDatafeed {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateDatafeed) IgnoreThrottled(ignorethrottled bool) *UpdateDatafeed {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateDatafeed) IgnoreUnavailable(ignoreunavailable bool) *UpdateDatafeed {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateDatafeed) ErrorTrace(errortrace bool) *UpdateDatafeed {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateDatafeed) FilterPath(filterpaths ...string) *UpdateDatafeed {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateDatafeed) Human(human bool) *UpdateDatafeed { _ = "STUB: not implemented"; return nil }

func (r *UpdateDatafeed) Pretty(pretty bool) *UpdateDatafeed { _ = "STUB: not implemented"; return nil }

func (r *UpdateDatafeed) Aggregations(aggregations map[string]types.Aggregations) *UpdateDatafeed {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateDatafeed) AddAggregation(key string, value types.AggregationsVariant) *UpdateDatafeed {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateDatafeed) ChunkingConfig(chunkingconfig types.ChunkingConfigVariant) *UpdateDatafeed {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateDatafeed) DelayedDataCheckConfig(delayeddatacheckconfig types.DelayedDataCheckConfigVariant) *UpdateDatafeed {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateDatafeed) Frequency(duration types.DurationVariant) *UpdateDatafeed {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateDatafeed) Indices(indices ...string) *UpdateDatafeed {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateDatafeed) IndicesOptions(indicesoptions types.IndicesOptionsVariant) *UpdateDatafeed {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateDatafeed) JobId(id string) *UpdateDatafeed { _ = "STUB: not implemented"; return nil }

func (r *UpdateDatafeed) MaxEmptySearches(maxemptysearches int) *UpdateDatafeed {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateDatafeed) Query(query types.QueryVariant) *UpdateDatafeed {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateDatafeed) QueryDelay(duration types.DurationVariant) *UpdateDatafeed {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateDatafeed) RuntimeMappings(runtimefields types.RuntimeFieldsVariant) *UpdateDatafeed {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateDatafeed) ScriptFields(scriptfields map[string]types.ScriptField) *UpdateDatafeed {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateDatafeed) AddScriptField(key string, value types.ScriptFieldVariant) *UpdateDatafeed {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateDatafeed) ScrollSize(scrollsize int) *UpdateDatafeed {
	_ = "STUB: not implemented"
	return nil
}
