package putcomponenttemplate

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

type PutComponentTemplate struct {
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

type NewPutComponentTemplate func(name string) *PutComponentTemplate

func NewPutComponentTemplateFunc(tp elastictransport.Interface) NewPutComponentTemplate {
	_ = "STUB: not implemented"
	return *new(NewPutComponentTemplate)
}

func New(tp elastictransport.Interface) *PutComponentTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutComponentTemplate) Raw(raw io.Reader) *PutComponentTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutComponentTemplate) Request(req *Request) *PutComponentTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutComponentTemplate) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r PutComponentTemplate) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r PutComponentTemplate) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *PutComponentTemplate) Header(key, value string) *PutComponentTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutComponentTemplate) _name(name string) *PutComponentTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutComponentTemplate) Create(create bool) *PutComponentTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutComponentTemplate) Cause(cause string) *PutComponentTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutComponentTemplate) MasterTimeout(duration string) *PutComponentTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutComponentTemplate) ErrorTrace(errortrace bool) *PutComponentTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutComponentTemplate) FilterPath(filterpaths ...string) *PutComponentTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutComponentTemplate) Human(human bool) *PutComponentTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutComponentTemplate) Pretty(pretty bool) *PutComponentTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutComponentTemplate) Deprecated(deprecated bool) *PutComponentTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutComponentTemplate) Meta_(metadata types.MetadataVariant) *PutComponentTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutComponentTemplate) Template(template types.IndexTemplateMappingVariant) *PutComponentTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutComponentTemplate) Version(versionnumber int64) *PutComponentTemplate {
	_ = "STUB: not implemented"
	return nil
}
