package clearcache

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/expandwildcard"
)

const (
	indexMask = iota + 1
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type ClearCache struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	paramSet int

	index string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewClearCache func() *ClearCache

func NewClearCacheFunc(tp elastictransport.Interface) NewClearCache {
	_ = "STUB: not implemented"
	return *new(NewClearCache)
}

func New(tp elastictransport.Interface) *ClearCache { _ = "STUB: not implemented"; return nil }

func (r *ClearCache) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r ClearCache) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r ClearCache) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r ClearCache) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *ClearCache) Header(key, value string) *ClearCache { _ = "STUB: not implemented"; return nil }

func (r *ClearCache) Index(index string) *ClearCache { _ = "STUB: not implemented"; return nil }

func (r *ClearCache) AllowNoIndices(allownoindices bool) *ClearCache {
	_ = "STUB: not implemented"
	return nil
}

func (r *ClearCache) ExpandWildcards(expandwildcards ...expandwildcard.ExpandWildcard) *ClearCache {
	_ = "STUB: not implemented"
	return nil
}

func (r *ClearCache) Fielddata(fielddata bool) *ClearCache { _ = "STUB: not implemented"; return nil }

func (r *ClearCache) Fields(fields ...string) *ClearCache { _ = "STUB: not implemented"; return nil }

func (r *ClearCache) IgnoreUnavailable(ignoreunavailable bool) *ClearCache {
	_ = "STUB: not implemented"
	return nil
}

func (r *ClearCache) Query(query bool) *ClearCache { _ = "STUB: not implemented"; return nil }

func (r *ClearCache) Request(request bool) *ClearCache { _ = "STUB: not implemented"; return nil }

func (r *ClearCache) ErrorTrace(errortrace bool) *ClearCache { _ = "STUB: not implemented"; return nil }

func (r *ClearCache) FilterPath(filterpaths ...string) *ClearCache {
	_ = "STUB: not implemented"
	return nil
}

func (r *ClearCache) Human(human bool) *ClearCache { _ = "STUB: not implemented"; return nil }

func (r *ClearCache) Pretty(pretty bool) *ClearCache { _ = "STUB: not implemented"; return nil }
