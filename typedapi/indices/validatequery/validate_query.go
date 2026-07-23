package validatequery

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
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/operator"
)

const (
	indexMask = iota + 1
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type ValidateQuery struct {
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

type NewValidateQuery func() *ValidateQuery

func NewValidateQueryFunc(tp elastictransport.Interface) NewValidateQuery {
	_ = "STUB: not implemented"
	return *new(NewValidateQuery)
}

func New(tp elastictransport.Interface) *ValidateQuery { _ = "STUB: not implemented"; return nil }

func (r *ValidateQuery) Raw(raw io.Reader) *ValidateQuery { _ = "STUB: not implemented"; return nil }

func (r *ValidateQuery) Request(req *Request) *ValidateQuery { _ = "STUB: not implemented"; return nil }

func (r *ValidateQuery) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r ValidateQuery) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r ValidateQuery) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *ValidateQuery) Header(key, value string) *ValidateQuery {
	_ = "STUB: not implemented"
	return nil
}

func (r *ValidateQuery) Index(index string) *ValidateQuery { _ = "STUB: not implemented"; return nil }

func (r *ValidateQuery) AllowNoIndices(allownoindices bool) *ValidateQuery {
	_ = "STUB: not implemented"
	return nil
}

func (r *ValidateQuery) AllShards(allshards bool) *ValidateQuery {
	_ = "STUB: not implemented"
	return nil
}

func (r *ValidateQuery) Analyzer(analyzer string) *ValidateQuery {
	_ = "STUB: not implemented"
	return nil
}

func (r *ValidateQuery) AnalyzeWildcard(analyzewildcard bool) *ValidateQuery {
	_ = "STUB: not implemented"
	return nil
}

func (r *ValidateQuery) DefaultOperator(defaultoperator operator.Operator) *ValidateQuery {
	_ = "STUB: not implemented"
	return nil
}

func (r *ValidateQuery) Df(df string) *ValidateQuery { _ = "STUB: not implemented"; return nil }

func (r *ValidateQuery) ExpandWildcards(expandwildcards ...expandwildcard.ExpandWildcard) *ValidateQuery {
	_ = "STUB: not implemented"
	return nil
}

func (r *ValidateQuery) Explain(explain bool) *ValidateQuery { _ = "STUB: not implemented"; return nil }

func (r *ValidateQuery) IgnoreUnavailable(ignoreunavailable bool) *ValidateQuery {
	_ = "STUB: not implemented"
	return nil
}

func (r *ValidateQuery) Lenient(lenient bool) *ValidateQuery { _ = "STUB: not implemented"; return nil }

func (r *ValidateQuery) Rewrite(rewrite bool) *ValidateQuery { _ = "STUB: not implemented"; return nil }

func (r *ValidateQuery) Q(q string) *ValidateQuery { _ = "STUB: not implemented"; return nil }

func (r *ValidateQuery) ErrorTrace(errortrace bool) *ValidateQuery {
	_ = "STUB: not implemented"
	return nil
}

func (r *ValidateQuery) FilterPath(filterpaths ...string) *ValidateQuery {
	_ = "STUB: not implemented"
	return nil
}

func (r *ValidateQuery) Human(human bool) *ValidateQuery { _ = "STUB: not implemented"; return nil }

func (r *ValidateQuery) Pretty(pretty bool) *ValidateQuery { _ = "STUB: not implemented"; return nil }

func (r *ValidateQuery) Query(query types.QueryVariant) *ValidateQuery {
	_ = "STUB: not implemented"
	return nil
}
