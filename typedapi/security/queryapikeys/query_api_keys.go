package queryapikeys

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

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type QueryApiKeys struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	req      *Request
	deferred []func(request *Request) error
	buf      *gobytes.Buffer

	paramSet int

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewQueryApiKeys func() *QueryApiKeys

func NewQueryApiKeysFunc(tp elastictransport.Interface) NewQueryApiKeys {
	_ = "STUB: not implemented"
	return *new(NewQueryApiKeys)
}

func New(tp elastictransport.Interface) *QueryApiKeys { _ = "STUB: not implemented"; return nil }

func (r *QueryApiKeys) Raw(raw io.Reader) *QueryApiKeys { _ = "STUB: not implemented"; return nil }

func (r *QueryApiKeys) Request(req *Request) *QueryApiKeys { _ = "STUB: not implemented"; return nil }

func (r *QueryApiKeys) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r QueryApiKeys) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r QueryApiKeys) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *QueryApiKeys) Header(key, value string) *QueryApiKeys {
	_ = "STUB: not implemented"
	return nil
}

func (r *QueryApiKeys) WithLimitedBy(withlimitedby bool) *QueryApiKeys {
	_ = "STUB: not implemented"
	return nil
}

func (r *QueryApiKeys) WithProfileUid(withprofileuid bool) *QueryApiKeys {
	_ = "STUB: not implemented"
	return nil
}

func (r *QueryApiKeys) TypedKeys(typedkeys bool) *QueryApiKeys {
	_ = "STUB: not implemented"
	return nil
}

func (r *QueryApiKeys) ErrorTrace(errortrace bool) *QueryApiKeys {
	_ = "STUB: not implemented"
	return nil
}

func (r *QueryApiKeys) FilterPath(filterpaths ...string) *QueryApiKeys {
	_ = "STUB: not implemented"
	return nil
}

func (r *QueryApiKeys) Human(human bool) *QueryApiKeys { _ = "STUB: not implemented"; return nil }

func (r *QueryApiKeys) Pretty(pretty bool) *QueryApiKeys { _ = "STUB: not implemented"; return nil }

func (r *QueryApiKeys) Aggregations(aggregations map[string]types.ApiKeyAggregationContainer) *QueryApiKeys {
	_ = "STUB: not implemented"
	return nil
}

func (r *QueryApiKeys) AddAggregation(key string, value types.ApiKeyAggregationContainerVariant) *QueryApiKeys {
	_ = "STUB: not implemented"
	return nil
}

func (r *QueryApiKeys) From(from int) *QueryApiKeys { _ = "STUB: not implemented"; return nil }

func (r *QueryApiKeys) Query(query types.ApiKeyQueryContainerVariant) *QueryApiKeys {
	_ = "STUB: not implemented"
	return nil
}

func (r *QueryApiKeys) SearchAfter(sortresults ...types.FieldValueVariant) *QueryApiKeys {
	_ = "STUB: not implemented"
	return nil
}

func (r *QueryApiKeys) SearchAfterValues(sortresultsvalues []types.FieldValue) *QueryApiKeys {
	_ = "STUB: not implemented"
	return nil
}

func (r *QueryApiKeys) Size(size int) *QueryApiKeys { _ = "STUB: not implemented"; return nil }

func (r *QueryApiKeys) Sort(sorts ...types.SortCombinationsVariant) *QueryApiKeys {
	_ = "STUB: not implemented"
	return nil
}

func (r *QueryApiKeys) SortValues(sortvalues []types.SortCombinations) *QueryApiKeys {
	_ = "STUB: not implemented"
	return nil
}
