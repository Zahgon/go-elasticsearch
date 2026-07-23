package openpointintime

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
	indexMask = iota + 1
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type OpenPointInTime struct {
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

type NewOpenPointInTime func(index string) *OpenPointInTime

func NewOpenPointInTimeFunc(tp elastictransport.Interface) NewOpenPointInTime {
	_ = "STUB: not implemented"
	return *new(NewOpenPointInTime)
}

func New(tp elastictransport.Interface) *OpenPointInTime { _ = "STUB: not implemented"; return nil }

func (r *OpenPointInTime) Raw(raw io.Reader) *OpenPointInTime {
	_ = "STUB: not implemented"
	return nil
}

func (r *OpenPointInTime) Request(req *Request) *OpenPointInTime {
	_ = "STUB: not implemented"
	return nil
}

func (r *OpenPointInTime) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r OpenPointInTime) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r OpenPointInTime) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *OpenPointInTime) Header(key, value string) *OpenPointInTime {
	_ = "STUB: not implemented"
	return nil
}

func (r *OpenPointInTime) _index(index string) *OpenPointInTime {
	_ = "STUB: not implemented"
	return nil
}

func (r *OpenPointInTime) KeepAlive(duration string) *OpenPointInTime {
	_ = "STUB: not implemented"
	return nil
}

func (r *OpenPointInTime) IgnoreUnavailable(ignoreunavailable bool) *OpenPointInTime {
	_ = "STUB: not implemented"
	return nil
}

func (r *OpenPointInTime) Preference(preference string) *OpenPointInTime {
	_ = "STUB: not implemented"
	return nil
}

func (r *OpenPointInTime) Routing(routings ...string) *OpenPointInTime {
	_ = "STUB: not implemented"
	return nil
}

func (r *OpenPointInTime) ExpandWildcards(expandwildcards ...expandwildcard.ExpandWildcard) *OpenPointInTime {
	_ = "STUB: not implemented"
	return nil
}

func (r *OpenPointInTime) AllowPartialSearchResults(allowpartialsearchresults bool) *OpenPointInTime {
	_ = "STUB: not implemented"
	return nil
}

func (r *OpenPointInTime) MaxConcurrentShardRequests(maxconcurrentshardrequests int) *OpenPointInTime {
	_ = "STUB: not implemented"
	return nil
}

func (r *OpenPointInTime) ErrorTrace(errortrace bool) *OpenPointInTime {
	_ = "STUB: not implemented"
	return nil
}

func (r *OpenPointInTime) FilterPath(filterpaths ...string) *OpenPointInTime {
	_ = "STUB: not implemented"
	return nil
}

func (r *OpenPointInTime) Human(human bool) *OpenPointInTime { _ = "STUB: not implemented"; return nil }

func (r *OpenPointInTime) Pretty(pretty bool) *OpenPointInTime {
	_ = "STUB: not implemented"
	return nil
}

func (r *OpenPointInTime) IndexFilter(indexfilter types.QueryVariant) *OpenPointInTime {
	_ = "STUB: not implemented"
	return nil
}

func (r *OpenPointInTime) ProjectRouting(projectrouting string) *OpenPointInTime {
	_ = "STUB: not implemented"
	return nil
}
