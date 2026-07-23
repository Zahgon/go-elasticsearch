package searchtemplate

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
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/expandwildcard"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/searchtype"
)

const (
	indexMask = iota + 1
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type SearchTemplate struct {
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

type NewSearchTemplate func() *SearchTemplate

func NewSearchTemplateFunc(tp elastictransport.Interface) NewSearchTemplate {
	_ = "STUB: not implemented"
	return *new(NewSearchTemplate)
}

func New(tp elastictransport.Interface) *SearchTemplate { _ = "STUB: not implemented"; return nil }

func (r *SearchTemplate) Raw(raw io.Reader) *SearchTemplate { _ = "STUB: not implemented"; return nil }

func (r *SearchTemplate) Request(req *Request) *SearchTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (r *SearchTemplate) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r SearchTemplate) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r SearchTemplate) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *SearchTemplate) Header(key, value string) *SearchTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (r *SearchTemplate) Index(index string) *SearchTemplate { _ = "STUB: not implemented"; return nil }

func (r *SearchTemplate) AllowNoIndices(allownoindices bool) *SearchTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (r *SearchTemplate) CcsMinimizeRoundtrips(ccsminimizeroundtrips bool) *SearchTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (r *SearchTemplate) ExpandWildcards(expandwildcards ...expandwildcard.ExpandWildcard) *SearchTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (r *SearchTemplate) IgnoreThrottled(ignorethrottled bool) *SearchTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (r *SearchTemplate) IgnoreUnavailable(ignoreunavailable bool) *SearchTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (r *SearchTemplate) Preference(preference string) *SearchTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (r *SearchTemplate) Routing(routings ...string) *SearchTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (r *SearchTemplate) Scroll(duration string) *SearchTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (r *SearchTemplate) SearchType(searchtype searchtype.SearchType) *SearchTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (r *SearchTemplate) RestTotalHitsAsInt(resttotalhitsasint bool) *SearchTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (r *SearchTemplate) TypedKeys(typedkeys bool) *SearchTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (r *SearchTemplate) ErrorTrace(errortrace bool) *SearchTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (r *SearchTemplate) FilterPath(filterpaths ...string) *SearchTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (r *SearchTemplate) Human(human bool) *SearchTemplate { _ = "STUB: not implemented"; return nil }

func (r *SearchTemplate) Pretty(pretty bool) *SearchTemplate { _ = "STUB: not implemented"; return nil }

func (r *SearchTemplate) Explain(explain bool) *SearchTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (r *SearchTemplate) Id(id string) *SearchTemplate { _ = "STUB: not implemented"; return nil }

func (r *SearchTemplate) Params(params map[string]json.RawMessage) *SearchTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (r *SearchTemplate) AddParam(key string, value json.RawMessage) *SearchTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (r *SearchTemplate) Profile(profile bool) *SearchTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (r *SearchTemplate) ProjectRouting(projectrouting string) *SearchTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (r *SearchTemplate) Source(scriptsource types.ScriptSourceVariant) *SearchTemplate {
	_ = "STUB: not implemented"
	return nil
}
