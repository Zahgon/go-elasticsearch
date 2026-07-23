package search

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
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/resultposition"
)

const (
	indexMask = iota + 1
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type Search struct {
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

type NewSearch func(index string) *Search

func NewSearchFunc(tp elastictransport.Interface) NewSearch {
	_ = "STUB: not implemented"
	return *new(NewSearch)
}

func New(tp elastictransport.Interface) *Search { _ = "STUB: not implemented"; return nil }

func (r *Search) Raw(raw io.Reader) *Search { _ = "STUB: not implemented"; return nil }

func (r *Search) Request(req *Request) *Search { _ = "STUB: not implemented"; return nil }

func (r *Search) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r Search) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r Search) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *Search) Header(key, value string) *Search { _ = "STUB: not implemented"; return nil }

func (r *Search) _index(index string) *Search { _ = "STUB: not implemented"; return nil }

func (r *Search) AllowNoIndices(allownoindices bool) *Search { _ = "STUB: not implemented"; return nil }

func (r *Search) ExpandWildcards(expandwildcards ...expandwildcard.ExpandWildcard) *Search {
	_ = "STUB: not implemented"
	return nil
}

func (r *Search) CcsMinimizeRoundtrips(ccsminimizeroundtrips bool) *Search {
	_ = "STUB: not implemented"
	return nil
}

func (r *Search) IgnoreUnavailable(ignoreunavailable bool) *Search {
	_ = "STUB: not implemented"
	return nil
}

func (r *Search) ErrorTrace(errortrace bool) *Search { _ = "STUB: not implemented"; return nil }

func (r *Search) FilterPath(filterpaths ...string) *Search { _ = "STUB: not implemented"; return nil }

func (r *Search) Human(human bool) *Search { _ = "STUB: not implemented"; return nil }

func (r *Search) Pretty(pretty bool) *Search { _ = "STUB: not implemented"; return nil }

func (r *Search) AllowPartialSearchResults(allowpartialsearchresults bool) *Search {
	_ = "STUB: not implemented"
	return nil
}

func (r *Search) AllowPartialSequenceResults(allowpartialsequenceresults bool) *Search {
	_ = "STUB: not implemented"
	return nil
}

func (r *Search) CaseSensitive(casesensitive bool) *Search { _ = "STUB: not implemented"; return nil }

func (r *Search) EventCategoryField(field string) *Search { _ = "STUB: not implemented"; return nil }

func (r *Search) FetchSize(fetchsize uint) *Search { _ = "STUB: not implemented"; return nil }

func (r *Search) Fields(fields ...types.FieldAndFormatVariant) *Search {
	_ = "STUB: not implemented"
	return nil
}

func (r *Search) Filter(filters ...types.QueryVariant) *Search {
	_ = "STUB: not implemented"
	return nil
}

func (r *Search) KeepAlive(duration types.DurationVariant) *Search {
	_ = "STUB: not implemented"
	return nil
}

func (r *Search) KeepOnCompletion(keeponcompletion bool) *Search {
	_ = "STUB: not implemented"
	return nil
}

func (r *Search) MaxSamplesPerKey(maxsamplesperkey int) *Search {
	_ = "STUB: not implemented"
	return nil
}

func (r *Search) ProjectRouting(projectrouting string) *Search {
	_ = "STUB: not implemented"
	return nil
}

func (r *Search) Query(query string) *Search { _ = "STUB: not implemented"; return nil }

func (r *Search) ResultPosition(resultposition resultposition.ResultPosition) *Search {
	_ = "STUB: not implemented"
	return nil
}

func (r *Search) RuntimeMappings(runtimefields types.RuntimeFieldsVariant) *Search {
	_ = "STUB: not implemented"
	return nil
}

func (r *Search) Size(size uint) *Search { _ = "STUB: not implemented"; return nil }

func (r *Search) TiebreakerField(field string) *Search { _ = "STUB: not implemented"; return nil }

func (r *Search) TimestampField(field string) *Search { _ = "STUB: not implemented"; return nil }

func (r *Search) WaitForCompletionTimeout(duration types.DurationVariant) *Search {
	_ = "STUB: not implemented"
	return nil
}
