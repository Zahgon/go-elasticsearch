package putindextemplate

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

type PutIndexTemplate struct {
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

type NewPutIndexTemplate func(name string) *PutIndexTemplate

func NewPutIndexTemplateFunc(tp elastictransport.Interface) NewPutIndexTemplate {
	_ = "STUB: not implemented"
	return *new(NewPutIndexTemplate)
}

func New(tp elastictransport.Interface) *PutIndexTemplate { _ = "STUB: not implemented"; return nil }

func (r *PutIndexTemplate) Raw(raw io.Reader) *PutIndexTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutIndexTemplate) Request(req *Request) *PutIndexTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutIndexTemplate) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r PutIndexTemplate) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r PutIndexTemplate) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *PutIndexTemplate) Header(key, value string) *PutIndexTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutIndexTemplate) _name(name string) *PutIndexTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutIndexTemplate) Create(create bool) *PutIndexTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutIndexTemplate) MasterTimeout(duration string) *PutIndexTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutIndexTemplate) Cause(cause string) *PutIndexTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutIndexTemplate) ErrorTrace(errortrace bool) *PutIndexTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutIndexTemplate) FilterPath(filterpaths ...string) *PutIndexTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutIndexTemplate) Human(human bool) *PutIndexTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutIndexTemplate) Pretty(pretty bool) *PutIndexTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutIndexTemplate) AllowAutoCreate(allowautocreate bool) *PutIndexTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutIndexTemplate) ComposedOf(composedofs ...string) *PutIndexTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutIndexTemplate) DataStream(datastream types.DataStreamVisibilityVariant) *PutIndexTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutIndexTemplate) Deprecated(deprecated bool) *PutIndexTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutIndexTemplate) IgnoreMissingComponentTemplates(ignoremissingcomponenttemplates ...string) *PutIndexTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutIndexTemplate) IndexPatterns(indices ...string) *PutIndexTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutIndexTemplate) Meta_(metadata types.MetadataVariant) *PutIndexTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutIndexTemplate) Priority(priority int64) *PutIndexTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutIndexTemplate) Template(template types.IndexTemplateMappingVariant) *PutIndexTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutIndexTemplate) Version(versionnumber int64) *PutIndexTemplate {
	_ = "STUB: not implemented"
	return nil
}
