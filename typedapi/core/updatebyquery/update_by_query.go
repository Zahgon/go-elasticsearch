package updatebyquery

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

type UpdateByQuery struct {
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

type NewUpdateByQuery func(index string) *UpdateByQuery

func NewUpdateByQueryFunc(tp elastictransport.Interface) NewUpdateByQuery {
	_ = "STUB: not implemented"
	return *new(NewUpdateByQuery)
}

func New(tp elastictransport.Interface) *UpdateByQuery { _ = "STUB: not implemented"; return nil }

func (r *UpdateByQuery) Raw(raw io.Reader) *UpdateByQuery { _ = "STUB: not implemented"; return nil }

func (r *UpdateByQuery) Request(req *Request) *UpdateByQuery { _ = "STUB: not implemented"; return nil }

func (r *UpdateByQuery) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r UpdateByQuery) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r UpdateByQuery) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *UpdateByQuery) Header(key, value string) *UpdateByQuery {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateByQuery) _index(index string) *UpdateByQuery { _ = "STUB: not implemented"; return nil }

func (r *UpdateByQuery) AllowNoIndices(allownoindices bool) *UpdateByQuery {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateByQuery) Analyzer(analyzer string) *UpdateByQuery {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateByQuery) AnalyzeWildcard(analyzewildcard bool) *UpdateByQuery {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateByQuery) DefaultOperator(defaultoperator operator.Operator) *UpdateByQuery {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateByQuery) Df(df string) *UpdateByQuery { _ = "STUB: not implemented"; return nil }

func (r *UpdateByQuery) ExpandWildcards(expandwildcards ...expandwildcard.ExpandWildcard) *UpdateByQuery {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateByQuery) From(from string) *UpdateByQuery { _ = "STUB: not implemented"; return nil }

func (r *UpdateByQuery) IgnoreUnavailable(ignoreunavailable bool) *UpdateByQuery {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateByQuery) Lenient(lenient bool) *UpdateByQuery { _ = "STUB: not implemented"; return nil }

func (r *UpdateByQuery) Pipeline(pipeline string) *UpdateByQuery {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateByQuery) Preference(preference string) *UpdateByQuery {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateByQuery) Q(q string) *UpdateByQuery { _ = "STUB: not implemented"; return nil }

func (r *UpdateByQuery) Refresh(refresh bool) *UpdateByQuery { _ = "STUB: not implemented"; return nil }

func (r *UpdateByQuery) RequestCache(requestcache bool) *UpdateByQuery {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateByQuery) RequestsPerSecond(requestspersecond string) *UpdateByQuery {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateByQuery) Routing(routings ...string) *UpdateByQuery {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateByQuery) Scroll(duration string) *UpdateByQuery {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateByQuery) ScrollSize(scrollsize string) *UpdateByQuery {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateByQuery) SearchTimeout(duration string) *UpdateByQuery {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateByQuery) SearchType(searchtype searchtype.SearchType) *UpdateByQuery {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateByQuery) Slices(slices string) *UpdateByQuery { _ = "STUB: not implemented"; return nil }

func (r *UpdateByQuery) Sort(sorts ...string) *UpdateByQuery { _ = "STUB: not implemented"; return nil }

func (r *UpdateByQuery) Stats(stats ...string) *UpdateByQuery {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateByQuery) TerminateAfter(terminateafter string) *UpdateByQuery {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateByQuery) Timeout(duration string) *UpdateByQuery {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateByQuery) Version(version bool) *UpdateByQuery { _ = "STUB: not implemented"; return nil }

func (r *UpdateByQuery) VersionType(versiontype bool) *UpdateByQuery {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateByQuery) WaitForActiveShards(waitforactiveshards string) *UpdateByQuery {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateByQuery) WaitForCompletion(waitforcompletion bool) *UpdateByQuery {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateByQuery) ErrorTrace(errortrace bool) *UpdateByQuery {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateByQuery) FilterPath(filterpaths ...string) *UpdateByQuery {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateByQuery) Human(human bool) *UpdateByQuery { _ = "STUB: not implemented"; return nil }

func (r *UpdateByQuery) Pretty(pretty bool) *UpdateByQuery { _ = "STUB: not implemented"; return nil }

func (r *UpdateByQuery) Conflicts(conflicts conflicts.Conflicts) *UpdateByQuery {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateByQuery) MaxDocs(maxdocs int64) *UpdateByQuery {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateByQuery) Query(query types.QueryVariant) *UpdateByQuery {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateByQuery) Script(script types.ScriptVariant) *UpdateByQuery {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateByQuery) Slice(slice types.SlicedScrollVariant) *UpdateByQuery {
	_ = "STUB: not implemented"
	return nil
}
