package knnsearch

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

type KnnSearch struct {
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

type NewKnnSearch func(index string) *KnnSearch

func NewKnnSearchFunc(tp elastictransport.Interface) NewKnnSearch {
	_ = "STUB: not implemented"
	return *new(NewKnnSearch)
}

func New(tp elastictransport.Interface) *KnnSearch { _ = "STUB: not implemented"; return nil }

func (r *KnnSearch) Raw(raw io.Reader) *KnnSearch { _ = "STUB: not implemented"; return nil }

func (r *KnnSearch) Request(req *Request) *KnnSearch { _ = "STUB: not implemented"; return nil }

func (r *KnnSearch) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r KnnSearch) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r KnnSearch) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *KnnSearch) Header(key, value string) *KnnSearch { _ = "STUB: not implemented"; return nil }

func (r *KnnSearch) _index(index string) *KnnSearch { _ = "STUB: not implemented"; return nil }

func (r *KnnSearch) Routing(routings ...string) *KnnSearch { _ = "STUB: not implemented"; return nil }

func (r *KnnSearch) ErrorTrace(errortrace bool) *KnnSearch { _ = "STUB: not implemented"; return nil }

func (r *KnnSearch) FilterPath(filterpaths ...string) *KnnSearch {
	_ = "STUB: not implemented"
	return nil
}

func (r *KnnSearch) Human(human bool) *KnnSearch { _ = "STUB: not implemented"; return nil }

func (r *KnnSearch) Pretty(pretty bool) *KnnSearch { _ = "STUB: not implemented"; return nil }

func (r *KnnSearch) DocvalueFields(docvaluefields ...types.FieldAndFormatVariant) *KnnSearch {
	_ = "STUB: not implemented"
	return nil
}

func (r *KnnSearch) DocvalueFieldsValues(docvaluefieldsvalues []types.FieldAndFormat) *KnnSearch {
	_ = "STUB: not implemented"
	return nil
}

func (r *KnnSearch) Fields(fields ...string) *KnnSearch { _ = "STUB: not implemented"; return nil }

func (r *KnnSearch) Filter(filters ...types.QueryVariant) *KnnSearch {
	_ = "STUB: not implemented"
	return nil
}

func (r *KnnSearch) Knn(knn types.KnnSearchQueryVariant) *KnnSearch {
	_ = "STUB: not implemented"
	return nil
}

func (r *KnnSearch) Source_(sourceconfig types.SourceConfigVariant) *KnnSearch {
	_ = "STUB: not implemented"
	return nil
}

func (r *KnnSearch) StoredFields(fields ...string) *KnnSearch {
	_ = "STUB: not implemented"
	return nil
}
