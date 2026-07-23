package query

import (
	gobytes "bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/sqlformat"
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type Query struct {
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

type NewQuery func() *Query

func NewQueryFunc(tp elastictransport.Interface) NewQuery {
	_ = "STUB: not implemented"
	return *new(NewQuery)
}

func New(tp elastictransport.Interface) *Query { _ = "STUB: not implemented"; return nil }

func (r *Query) Raw(raw io.Reader) *Query { _ = "STUB: not implemented"; return nil }

func (r *Query) Request(req *Request) *Query { _ = "STUB: not implemented"; return nil }

func (r *Query) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r Query) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r Query) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *Query) Header(key, value string) *Query { _ = "STUB: not implemented"; return nil }

func (r *Query) Format(format sqlformat.SqlFormat) *Query { _ = "STUB: not implemented"; return nil }

func (r *Query) ErrorTrace(errortrace bool) *Query { _ = "STUB: not implemented"; return nil }

func (r *Query) FilterPath(filterpaths ...string) *Query { _ = "STUB: not implemented"; return nil }

func (r *Query) Human(human bool) *Query { _ = "STUB: not implemented"; return nil }

func (r *Query) Pretty(pretty bool) *Query { _ = "STUB: not implemented"; return nil }

func (r *Query) AllowPartialSearchResults(allowpartialsearchresults bool) *Query {
	_ = "STUB: not implemented"
	return nil
}

func (r *Query) Catalog(catalog string) *Query { _ = "STUB: not implemented"; return nil }

func (r *Query) Columnar(columnar bool) *Query { _ = "STUB: not implemented"; return nil }

func (r *Query) Cursor(cursor string) *Query { _ = "STUB: not implemented"; return nil }

func (r *Query) FetchSize(fetchsize int) *Query { _ = "STUB: not implemented"; return nil }

func (r *Query) FieldMultiValueLeniency(fieldmultivalueleniency bool) *Query {
	_ = "STUB: not implemented"
	return nil
}

func (r *Query) Filter(filter types.QueryVariant) *Query { _ = "STUB: not implemented"; return nil }

func (r *Query) IndexUsingFrozen(indexusingfrozen bool) *Query {
	_ = "STUB: not implemented"
	return nil
}

func (r *Query) KeepAlive(duration types.DurationVariant) *Query {
	_ = "STUB: not implemented"
	return nil
}

func (r *Query) KeepOnCompletion(keeponcompletion bool) *Query {
	_ = "STUB: not implemented"
	return nil
}

func (r *Query) PageTimeout(duration types.DurationVariant) *Query {
	_ = "STUB: not implemented"
	return nil
}

func (r *Query) Params(params ...json.RawMessage) *Query { _ = "STUB: not implemented"; return nil }

func (r *Query) ProjectRouting(projectrouting string) *Query { _ = "STUB: not implemented"; return nil }

func (r *Query) Query(query string) *Query { _ = "STUB: not implemented"; return nil }

func (r *Query) RequestTimeout(duration types.DurationVariant) *Query {
	_ = "STUB: not implemented"
	return nil
}

func (r *Query) RuntimeMappings(runtimefields types.RuntimeFieldsVariant) *Query {
	_ = "STUB: not implemented"
	return nil
}

func (r *Query) TimeZone(timezone string) *Query { _ = "STUB: not implemented"; return nil }

func (r *Query) WaitForCompletionTimeout(duration types.DurationVariant) *Query {
	_ = "STUB: not implemented"
	return nil
}
