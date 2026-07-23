package analyze

import (
	gobytes "bytes"
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

const (
	indexMask = iota + 1
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type Analyze struct {
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

type NewAnalyze func() *Analyze

func NewAnalyzeFunc(tp elastictransport.Interface) NewAnalyze {
	_ = "STUB: not implemented"
	return *new(NewAnalyze)
}

func New(tp elastictransport.Interface) *Analyze { _ = "STUB: not implemented"; return nil }

func (r *Analyze) Raw(raw io.Reader) *Analyze { _ = "STUB: not implemented"; return nil }

func (r *Analyze) Request(req *Request) *Analyze { _ = "STUB: not implemented"; return nil }

func (r *Analyze) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r Analyze) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r Analyze) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *Analyze) Header(key, value string) *Analyze { _ = "STUB: not implemented"; return nil }

func (r *Analyze) Index(index string) *Analyze { _ = "STUB: not implemented"; return nil }

func (r *Analyze) ErrorTrace(errortrace bool) *Analyze { _ = "STUB: not implemented"; return nil }

func (r *Analyze) FilterPath(filterpaths ...string) *Analyze { _ = "STUB: not implemented"; return nil }

func (r *Analyze) Human(human bool) *Analyze { _ = "STUB: not implemented"; return nil }

func (r *Analyze) Pretty(pretty bool) *Analyze { _ = "STUB: not implemented"; return nil }

func (r *Analyze) Analyzer(analyzer string) *Analyze { _ = "STUB: not implemented"; return nil }

func (r *Analyze) Attributes(attributes ...string) *Analyze { _ = "STUB: not implemented"; return nil }

func (r *Analyze) CharFilter(charfilters ...types.CharFilterVariant) *Analyze {
	_ = "STUB: not implemented"
	return nil
}

func (r *Analyze) CharFilterValues(charfiltervalues []types.CharFilter) *Analyze {
	_ = "STUB: not implemented"
	return nil
}

func (r *Analyze) Explain(explain bool) *Analyze { _ = "STUB: not implemented"; return nil }

func (r *Analyze) Field(field string) *Analyze { _ = "STUB: not implemented"; return nil }

func (r *Analyze) Filter(filters ...types.TokenFilterVariant) *Analyze {
	_ = "STUB: not implemented"
	return nil
}

func (r *Analyze) FilterValues(filtervalues []types.TokenFilter) *Analyze {
	_ = "STUB: not implemented"
	return nil
}

func (r *Analyze) Normalizer(normalizer string) *Analyze { _ = "STUB: not implemented"; return nil }

func (r *Analyze) Text(texttoanalyzes ...string) *Analyze { _ = "STUB: not implemented"; return nil }

func (r *Analyze) Tokenizer(tokenizer types.TokenizerVariant) *Analyze {
	_ = "STUB: not implemented"
	return nil
}
