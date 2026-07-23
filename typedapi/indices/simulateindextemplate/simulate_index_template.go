package simulateindextemplate

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

type SimulateIndexTemplate struct {
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

type NewSimulateIndexTemplate func(name string) *SimulateIndexTemplate

func NewSimulateIndexTemplateFunc(tp elastictransport.Interface) NewSimulateIndexTemplate {
	_ = "STUB: not implemented"
	return *new(NewSimulateIndexTemplate)
}

func New(tp elastictransport.Interface) *SimulateIndexTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (r *SimulateIndexTemplate) Raw(raw io.Reader) *SimulateIndexTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (r *SimulateIndexTemplate) Request(req *Request) *SimulateIndexTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (r *SimulateIndexTemplate) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r SimulateIndexTemplate) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r SimulateIndexTemplate) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *SimulateIndexTemplate) Header(key, value string) *SimulateIndexTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (r *SimulateIndexTemplate) _name(name string) *SimulateIndexTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (r *SimulateIndexTemplate) Create(create bool) *SimulateIndexTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (r *SimulateIndexTemplate) Cause(cause string) *SimulateIndexTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (r *SimulateIndexTemplate) MasterTimeout(duration string) *SimulateIndexTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (r *SimulateIndexTemplate) IncludeDefaults(includedefaults bool) *SimulateIndexTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (r *SimulateIndexTemplate) ErrorTrace(errortrace bool) *SimulateIndexTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (r *SimulateIndexTemplate) FilterPath(filterpaths ...string) *SimulateIndexTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (r *SimulateIndexTemplate) Human(human bool) *SimulateIndexTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (r *SimulateIndexTemplate) Pretty(pretty bool) *SimulateIndexTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (r *SimulateIndexTemplate) AllowAutoCreate(allowautocreate bool) *SimulateIndexTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (r *SimulateIndexTemplate) ComposedOf(composedofs ...string) *SimulateIndexTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (r *SimulateIndexTemplate) CreatedDate(datetime types.DateTimeVariant) *SimulateIndexTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (r *SimulateIndexTemplate) CreatedDateMillis(epochtimeunitmillis int64) *SimulateIndexTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (r *SimulateIndexTemplate) DataStream(datastream types.IndexTemplateDataStreamConfigurationVariant) *SimulateIndexTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (r *SimulateIndexTemplate) Deprecated(deprecated bool) *SimulateIndexTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (r *SimulateIndexTemplate) IgnoreMissingComponentTemplates(names ...string) *SimulateIndexTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (r *SimulateIndexTemplate) IndexPatterns(names ...string) *SimulateIndexTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (r *SimulateIndexTemplate) Meta_(metadata types.MetadataVariant) *SimulateIndexTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (r *SimulateIndexTemplate) ModifiedDate(datetime types.DateTimeVariant) *SimulateIndexTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (r *SimulateIndexTemplate) ModifiedDateMillis(epochtimeunitmillis int64) *SimulateIndexTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (r *SimulateIndexTemplate) Priority(priority int64) *SimulateIndexTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (r *SimulateIndexTemplate) Template(template types.IndexTemplateSummaryVariant) *SimulateIndexTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (r *SimulateIndexTemplate) Version(versionnumber int64) *SimulateIndexTemplate {
	_ = "STUB: not implemented"
	return nil
}
