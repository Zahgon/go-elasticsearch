package searchshards

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

type SearchShards struct {
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

type NewSearchShards func() *SearchShards

func NewSearchShardsFunc(tp elastictransport.Interface) NewSearchShards {
	_ = "STUB: not implemented"
	return *new(NewSearchShards)
}

func New(tp elastictransport.Interface) *SearchShards { _ = "STUB: not implemented"; return nil }

func (r *SearchShards) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r SearchShards) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r SearchShards) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r SearchShards) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *SearchShards) Header(key, value string) *SearchShards {
	_ = "STUB: not implemented"
	return nil
}

func (r *SearchShards) Index(index string) *SearchShards { _ = "STUB: not implemented"; return nil }

func (r *SearchShards) AllowNoIndices(allownoindices bool) *SearchShards {
	_ = "STUB: not implemented"
	return nil
}

func (r *SearchShards) ExpandWildcards(expandwildcards ...expandwildcard.ExpandWildcard) *SearchShards {
	_ = "STUB: not implemented"
	return nil
}

func (r *SearchShards) IgnoreUnavailable(ignoreunavailable bool) *SearchShards {
	_ = "STUB: not implemented"
	return nil
}

func (r *SearchShards) Local(local bool) *SearchShards { _ = "STUB: not implemented"; return nil }

func (r *SearchShards) MasterTimeout(duration string) *SearchShards {
	_ = "STUB: not implemented"
	return nil
}

func (r *SearchShards) Preference(preference string) *SearchShards {
	_ = "STUB: not implemented"
	return nil
}

func (r *SearchShards) Routing(routings ...string) *SearchShards {
	_ = "STUB: not implemented"
	return nil
}

func (r *SearchShards) Slice_(slice_ string) *SearchShards { _ = "STUB: not implemented"; return nil }

func (r *SearchShards) ErrorTrace(errortrace bool) *SearchShards {
	_ = "STUB: not implemented"
	return nil
}

func (r *SearchShards) FilterPath(filterpaths ...string) *SearchShards {
	_ = "STUB: not implemented"
	return nil
}

func (r *SearchShards) Human(human bool) *SearchShards { _ = "STUB: not implemented"; return nil }

func (r *SearchShards) Pretty(pretty bool) *SearchShards { _ = "STUB: not implemented"; return nil }
