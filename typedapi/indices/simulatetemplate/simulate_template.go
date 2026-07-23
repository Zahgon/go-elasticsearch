package simulatetemplate

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

type SimulateTemplate struct {
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

type NewSimulateTemplate func() *SimulateTemplate

func NewSimulateTemplateFunc(tp elastictransport.Interface) NewSimulateTemplate {
	_ = "STUB: not implemented"
	return *new(NewSimulateTemplate)
}

func New(tp elastictransport.Interface) *SimulateTemplate { _ = "STUB: not implemented"; return nil }

func (r *SimulateTemplate) Raw(raw io.Reader) *SimulateTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (r *SimulateTemplate) Request(req *Request) *SimulateTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (r *SimulateTemplate) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r SimulateTemplate) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r SimulateTemplate) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *SimulateTemplate) Header(key, value string) *SimulateTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (r *SimulateTemplate) Name(name string) *SimulateTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (r *SimulateTemplate) Create(create bool) *SimulateTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (r *SimulateTemplate) Cause(cause string) *SimulateTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (r *SimulateTemplate) MasterTimeout(duration string) *SimulateTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (r *SimulateTemplate) IncludeDefaults(includedefaults bool) *SimulateTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (r *SimulateTemplate) ErrorTrace(errortrace bool) *SimulateTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (r *SimulateTemplate) FilterPath(filterpaths ...string) *SimulateTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (r *SimulateTemplate) Human(human bool) *SimulateTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (r *SimulateTemplate) Pretty(pretty bool) *SimulateTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (r *SimulateTemplate) AllowAutoCreate(allowautocreate bool) *SimulateTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (r *SimulateTemplate) ComposedOf(composedofs ...string) *SimulateTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (r *SimulateTemplate) DataStream(datastream types.DataStreamVisibilityVariant) *SimulateTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (r *SimulateTemplate) Deprecated(deprecated bool) *SimulateTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (r *SimulateTemplate) IgnoreMissingComponentTemplates(ignoremissingcomponenttemplates ...string) *SimulateTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (r *SimulateTemplate) IndexPatterns(indices ...string) *SimulateTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (r *SimulateTemplate) Meta_(metadata types.MetadataVariant) *SimulateTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (r *SimulateTemplate) Priority(priority int64) *SimulateTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (r *SimulateTemplate) Template(template types.IndexTemplateMappingVariant) *SimulateTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (r *SimulateTemplate) Version(versionnumber int64) *SimulateTemplate {
	_ = "STUB: not implemented"
	return nil
}
