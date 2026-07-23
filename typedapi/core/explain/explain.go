package explain

import (
	gobytes "bytes"
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/operator"
)

const (
	idMask = iota + 1

	indexMask
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type Explain struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	req      *Request
	deferred []func(request *Request) error
	buf      *gobytes.Buffer

	paramSet int

	id    string
	index string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewExplain func(index, id string) *Explain

func NewExplainFunc(tp elastictransport.Interface) NewExplain {
	_ = "STUB: not implemented"
	return *new(NewExplain)
}

func New(tp elastictransport.Interface) *Explain { _ = "STUB: not implemented"; return nil }

func (r *Explain) Raw(raw io.Reader) *Explain { _ = "STUB: not implemented"; return nil }

func (r *Explain) Request(req *Request) *Explain { _ = "STUB: not implemented"; return nil }

func (r *Explain) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r Explain) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r Explain) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *Explain) Header(key, value string) *Explain { _ = "STUB: not implemented"; return nil }

func (r *Explain) _id(id string) *Explain { _ = "STUB: not implemented"; return nil }

func (r *Explain) _index(index string) *Explain { _ = "STUB: not implemented"; return nil }

func (r *Explain) Analyzer(analyzer string) *Explain { _ = "STUB: not implemented"; return nil }

func (r *Explain) AnalyzeWildcard(analyzewildcard bool) *Explain {
	_ = "STUB: not implemented"
	return nil
}

func (r *Explain) DefaultOperator(defaultoperator operator.Operator) *Explain {
	_ = "STUB: not implemented"
	return nil
}

func (r *Explain) Df(df string) *Explain { _ = "STUB: not implemented"; return nil }

func (r *Explain) Lenient(lenient bool) *Explain { _ = "STUB: not implemented"; return nil }

func (r *Explain) Preference(preference string) *Explain { _ = "STUB: not implemented"; return nil }

func (r *Explain) Routing(routings ...string) *Explain { _ = "STUB: not implemented"; return nil }

func (r *Explain) Source_(sourceconfigparam string) *Explain { _ = "STUB: not implemented"; return nil }

func (r *Explain) SourceExcludes_(fields ...string) *Explain { _ = "STUB: not implemented"; return nil }

func (r *Explain) SourceIncludes_(fields ...string) *Explain { _ = "STUB: not implemented"; return nil }

func (r *Explain) StoredFields(fields ...string) *Explain { _ = "STUB: not implemented"; return nil }

func (r *Explain) Q(q string) *Explain { _ = "STUB: not implemented"; return nil }

func (r *Explain) ErrorTrace(errortrace bool) *Explain { _ = "STUB: not implemented"; return nil }

func (r *Explain) FilterPath(filterpaths ...string) *Explain { _ = "STUB: not implemented"; return nil }

func (r *Explain) Human(human bool) *Explain { _ = "STUB: not implemented"; return nil }

func (r *Explain) Pretty(pretty bool) *Explain { _ = "STUB: not implemented"; return nil }

func (r *Explain) Query(query types.QueryVariant) *Explain { _ = "STUB: not implemented"; return nil }
