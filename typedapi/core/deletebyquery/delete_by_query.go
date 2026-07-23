package deletebyquery

import (
	gobytes "bytes"
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/conflicts"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/expandwildcard"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/operator"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/searchtype"
)

const (
	indexMask = iota + 1
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type DeleteByQuery struct {
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

type NewDeleteByQuery func(index string) *DeleteByQuery

func NewDeleteByQueryFunc(tp elastictransport.Interface) NewDeleteByQuery {
	_ = "STUB: not implemented"
	return *new(NewDeleteByQuery)
}

func New(tp elastictransport.Interface) *DeleteByQuery { _ = "STUB: not implemented"; return nil }

func (r *DeleteByQuery) Raw(raw io.Reader) *DeleteByQuery { _ = "STUB: not implemented"; return nil }

func (r *DeleteByQuery) Request(req *Request) *DeleteByQuery { _ = "STUB: not implemented"; return nil }

func (r *DeleteByQuery) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r DeleteByQuery) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r DeleteByQuery) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *DeleteByQuery) Header(key, value string) *DeleteByQuery {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteByQuery) _index(index string) *DeleteByQuery { _ = "STUB: not implemented"; return nil }

func (r *DeleteByQuery) AllowNoIndices(allownoindices bool) *DeleteByQuery {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteByQuery) Analyzer(analyzer string) *DeleteByQuery {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteByQuery) AnalyzeWildcard(analyzewildcard bool) *DeleteByQuery {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteByQuery) Conflicts(conflicts conflicts.Conflicts) *DeleteByQuery {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteByQuery) DefaultOperator(defaultoperator operator.Operator) *DeleteByQuery {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteByQuery) Df(df string) *DeleteByQuery { _ = "STUB: not implemented"; return nil }

func (r *DeleteByQuery) ExpandWildcards(expandwildcards ...expandwildcard.ExpandWildcard) *DeleteByQuery {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteByQuery) From(from string) *DeleteByQuery { _ = "STUB: not implemented"; return nil }

func (r *DeleteByQuery) IgnoreUnavailable(ignoreunavailable bool) *DeleteByQuery {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteByQuery) Lenient(lenient bool) *DeleteByQuery { _ = "STUB: not implemented"; return nil }

func (r *DeleteByQuery) Preference(preference string) *DeleteByQuery {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteByQuery) Refresh(refresh bool) *DeleteByQuery { _ = "STUB: not implemented"; return nil }

func (r *DeleteByQuery) RequestCache(requestcache bool) *DeleteByQuery {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteByQuery) RequestsPerSecond(requestspersecond string) *DeleteByQuery {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteByQuery) Routing(routings ...string) *DeleteByQuery {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteByQuery) Q(q string) *DeleteByQuery { _ = "STUB: not implemented"; return nil }

func (r *DeleteByQuery) Scroll(duration string) *DeleteByQuery {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteByQuery) ScrollSize(scrollsize string) *DeleteByQuery {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteByQuery) SearchTimeout(duration string) *DeleteByQuery {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteByQuery) SearchType(searchtype searchtype.SearchType) *DeleteByQuery {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteByQuery) Slices(slices string) *DeleteByQuery { _ = "STUB: not implemented"; return nil }

func (r *DeleteByQuery) Stats(stats ...string) *DeleteByQuery {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteByQuery) TerminateAfter(terminateafter string) *DeleteByQuery {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteByQuery) Timeout(duration string) *DeleteByQuery {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteByQuery) Version(version bool) *DeleteByQuery { _ = "STUB: not implemented"; return nil }

func (r *DeleteByQuery) WaitForActiveShards(waitforactiveshards string) *DeleteByQuery {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteByQuery) WaitForCompletion(waitforcompletion bool) *DeleteByQuery {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteByQuery) ErrorTrace(errortrace bool) *DeleteByQuery {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteByQuery) FilterPath(filterpaths ...string) *DeleteByQuery {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteByQuery) Human(human bool) *DeleteByQuery { _ = "STUB: not implemented"; return nil }

func (r *DeleteByQuery) Pretty(pretty bool) *DeleteByQuery { _ = "STUB: not implemented"; return nil }

func (r *DeleteByQuery) MaxDocs(maxdocs int64) *DeleteByQuery {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteByQuery) Query(query types.QueryVariant) *DeleteByQuery {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteByQuery) Slice(slice types.SlicedScrollVariant) *DeleteByQuery {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteByQuery) Sort(sorts ...types.SortCombinationsVariant) *DeleteByQuery {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteByQuery) SortValues(sortvalues []types.SortCombinations) *DeleteByQuery {
	_ = "STUB: not implemented"
	return nil
}
