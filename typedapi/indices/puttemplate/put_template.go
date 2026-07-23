package puttemplate

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
	nameMask = iota + 1
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type PutTemplate struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	req      *Request
	deferred []func(request *Request) error
	buf      *gobytes.Buffer

	paramSet int

	name string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewPutTemplate func(name string) *PutTemplate

func NewPutTemplateFunc(tp elastictransport.Interface) NewPutTemplate {
	_ = "STUB: not implemented"
	return *new(NewPutTemplate)
}

func New(tp elastictransport.Interface) *PutTemplate { _ = "STUB: not implemented"; return nil }

func (r *PutTemplate) Raw(raw io.Reader) *PutTemplate { _ = "STUB: not implemented"; return nil }

func (r *PutTemplate) Request(req *Request) *PutTemplate { _ = "STUB: not implemented"; return nil }

func (r *PutTemplate) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r PutTemplate) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r PutTemplate) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *PutTemplate) Header(key, value string) *PutTemplate { _ = "STUB: not implemented"; return nil }

func (r *PutTemplate) _name(name string) *PutTemplate { _ = "STUB: not implemented"; return nil }

func (r *PutTemplate) Create(create bool) *PutTemplate { _ = "STUB: not implemented"; return nil }

func (r *PutTemplate) MasterTimeout(duration string) *PutTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutTemplate) Cause(cause string) *PutTemplate { _ = "STUB: not implemented"; return nil }

func (r *PutTemplate) ErrorTrace(errortrace bool) *PutTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutTemplate) FilterPath(filterpaths ...string) *PutTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutTemplate) Human(human bool) *PutTemplate { _ = "STUB: not implemented"; return nil }

func (r *PutTemplate) Pretty(pretty bool) *PutTemplate { _ = "STUB: not implemented"; return nil }

func (r *PutTemplate) Aliases(aliases map[string]types.Alias) *PutTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutTemplate) AddAlias(key string, value types.AliasVariant) *PutTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutTemplate) IndexPatterns(indexpatterns ...string) *PutTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutTemplate) Mappings(mappings types.TypeMappingVariant) *PutTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutTemplate) Order(order int) *PutTemplate { _ = "STUB: not implemented"; return nil }

func (r *PutTemplate) Settings(settings types.IndexSettingsVariant) *PutTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutTemplate) Version(versionnumber int64) *PutTemplate {
	_ = "STUB: not implemented"
	return nil
}
