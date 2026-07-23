package updatefiltering

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
	connectoridMask = iota + 1
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type UpdateFiltering struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	req      *Request
	deferred []func(request *Request) error
	buf      *gobytes.Buffer

	paramSet int

	connectorid string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewUpdateFiltering func(connectorid string) *UpdateFiltering

func NewUpdateFilteringFunc(tp elastictransport.Interface) NewUpdateFiltering {
	_ = "STUB: not implemented"
	return *new(NewUpdateFiltering)
}

func New(tp elastictransport.Interface) *UpdateFiltering { _ = "STUB: not implemented"; return nil }

func (r *UpdateFiltering) Raw(raw io.Reader) *UpdateFiltering {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateFiltering) Request(req *Request) *UpdateFiltering {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateFiltering) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r UpdateFiltering) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r UpdateFiltering) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *UpdateFiltering) Header(key, value string) *UpdateFiltering {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateFiltering) _connectorid(connectorid string) *UpdateFiltering {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateFiltering) ErrorTrace(errortrace bool) *UpdateFiltering {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateFiltering) FilterPath(filterpaths ...string) *UpdateFiltering {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateFiltering) Human(human bool) *UpdateFiltering { _ = "STUB: not implemented"; return nil }

func (r *UpdateFiltering) Pretty(pretty bool) *UpdateFiltering {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateFiltering) AdvancedSnippet(advancedsnippet types.FilteringAdvancedSnippetVariant) *UpdateFiltering {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateFiltering) Filtering(filterings ...types.FilteringConfigVariant) *UpdateFiltering {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateFiltering) FilteringValues(filteringvalues []types.FilteringConfig) *UpdateFiltering {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateFiltering) Rules(rules ...types.FilteringRuleVariant) *UpdateFiltering {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateFiltering) RulesValues(rulesvalues []types.FilteringRule) *UpdateFiltering {
	_ = "STUB: not implemented"
	return nil
}
