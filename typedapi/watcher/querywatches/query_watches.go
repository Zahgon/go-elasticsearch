package querywatches

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

type QueryWatches struct {
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

type NewQueryWatches func() *QueryWatches

func NewQueryWatchesFunc(tp elastictransport.Interface) NewQueryWatches {
	_ = "STUB: not implemented"
	return *new(NewQueryWatches)
}

func New(tp elastictransport.Interface) *QueryWatches { _ = "STUB: not implemented"; return nil }

func (r *QueryWatches) Raw(raw io.Reader) *QueryWatches { _ = "STUB: not implemented"; return nil }

func (r *QueryWatches) Request(req *Request) *QueryWatches { _ = "STUB: not implemented"; return nil }

func (r *QueryWatches) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r QueryWatches) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r QueryWatches) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *QueryWatches) Header(key, value string) *QueryWatches {
	_ = "STUB: not implemented"
	return nil
}

func (r *QueryWatches) ErrorTrace(errortrace bool) *QueryWatches {
	_ = "STUB: not implemented"
	return nil
}

func (r *QueryWatches) FilterPath(filterpaths ...string) *QueryWatches {
	_ = "STUB: not implemented"
	return nil
}

func (r *QueryWatches) Human(human bool) *QueryWatches { _ = "STUB: not implemented"; return nil }

func (r *QueryWatches) Pretty(pretty bool) *QueryWatches { _ = "STUB: not implemented"; return nil }

func (r *QueryWatches) From(from int) *QueryWatches { _ = "STUB: not implemented"; return nil }

func (r *QueryWatches) Query(query types.QueryVariant) *QueryWatches {
	_ = "STUB: not implemented"
	return nil
}

func (r *QueryWatches) SearchAfter(sortresults ...types.FieldValueVariant) *QueryWatches {
	_ = "STUB: not implemented"
	return nil
}

func (r *QueryWatches) SearchAfterValues(sortresultsvalues []types.FieldValue) *QueryWatches {
	_ = "STUB: not implemented"
	return nil
}

func (r *QueryWatches) Size(size int) *QueryWatches { _ = "STUB: not implemented"; return nil }

func (r *QueryWatches) Sort(sorts ...types.SortCombinationsVariant) *QueryWatches {
	_ = "STUB: not implemented"
	return nil
}

func (r *QueryWatches) SortValues(sortvalues []types.SortCombinations) *QueryWatches {
	_ = "STUB: not implemented"
	return nil
}
