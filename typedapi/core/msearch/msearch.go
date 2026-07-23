package msearch

import (
	gobytes "bytes"
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/expandwildcard"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/searchtype"
)

const (
	indexMask = iota + 1
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type Msearch struct {
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

type NewMsearch func() *Msearch

func NewMsearchFunc(tp elastictransport.Interface) NewMsearch {
	_ = "STUB: not implemented"
	return *new(NewMsearch)
}

func New(tp elastictransport.Interface) *Msearch { _ = "STUB: not implemented"; return nil }

func (r *Msearch) Raw(raw io.Reader) *Msearch { _ = "STUB: not implemented"; return nil }

func (r *Msearch) Request(req *Request) *Msearch { _ = "STUB: not implemented"; return nil }

func (r *Msearch) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r Msearch) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r Msearch) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *Msearch) Header(key, value string) *Msearch { _ = "STUB: not implemented"; return nil }

func (r *Msearch) Index(index string) *Msearch { _ = "STUB: not implemented"; return nil }

func (r *Msearch) AllowNoIndices(allownoindices bool) *Msearch {
	_ = "STUB: not implemented"
	return nil
}

func (r *Msearch) CcsMinimizeRoundtrips(ccsminimizeroundtrips bool) *Msearch {
	_ = "STUB: not implemented"
	return nil
}

func (r *Msearch) ExpandWildcards(expandwildcards ...expandwildcard.ExpandWildcard) *Msearch {
	_ = "STUB: not implemented"
	return nil
}

func (r *Msearch) IgnoreThrottled(ignorethrottled bool) *Msearch {
	_ = "STUB: not implemented"
	return nil
}

func (r *Msearch) IgnoreUnavailable(ignoreunavailable bool) *Msearch {
	_ = "STUB: not implemented"
	return nil
}

func (r *Msearch) IncludeNamedQueriesScore(includenamedqueriesscore bool) *Msearch {
	_ = "STUB: not implemented"
	return nil
}

func (r *Msearch) MaxConcurrentSearches(maxconcurrentsearches int) *Msearch {
	_ = "STUB: not implemented"
	return nil
}

func (r *Msearch) MaxConcurrentShardRequests(maxconcurrentshardrequests int) *Msearch {
	_ = "STUB: not implemented"
	return nil
}

func (r *Msearch) PreFilterShardSize(prefiltershardsize string) *Msearch {
	_ = "STUB: not implemented"
	return nil
}

func (r *Msearch) ProjectRouting(projectrouting string) *Msearch {
	_ = "STUB: not implemented"
	return nil
}

func (r *Msearch) RestTotalHitsAsInt(resttotalhitsasint bool) *Msearch {
	_ = "STUB: not implemented"
	return nil
}

func (r *Msearch) Routing(routings ...string) *Msearch { _ = "STUB: not implemented"; return nil }

func (r *Msearch) SearchType(searchtype searchtype.SearchType) *Msearch {
	_ = "STUB: not implemented"
	return nil
}

func (r *Msearch) Slice_(slice_ string) *Msearch { _ = "STUB: not implemented"; return nil }

func (r *Msearch) TypedKeys(typedkeys bool) *Msearch { _ = "STUB: not implemented"; return nil }

func (r *Msearch) ErrorTrace(errortrace bool) *Msearch { _ = "STUB: not implemented"; return nil }

func (r *Msearch) FilterPath(filterpaths ...string) *Msearch { _ = "STUB: not implemented"; return nil }

func (r *Msearch) Human(human bool) *Msearch { _ = "STUB: not implemented"; return nil }

func (r *Msearch) Pretty(pretty bool) *Msearch { _ = "STUB: not implemented"; return nil }
