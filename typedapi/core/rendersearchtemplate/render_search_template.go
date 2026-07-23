package rendersearchtemplate

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
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type RenderSearchTemplate struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	req      *Request
	deferred []func(request *Request) error
	buf      *gobytes.Buffer

	paramSet int

	id string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewRenderSearchTemplate func() *RenderSearchTemplate

func NewRenderSearchTemplateFunc(tp elastictransport.Interface) NewRenderSearchTemplate {
	_ = "STUB: not implemented"
	return *new(NewRenderSearchTemplate)
}

func New(tp elastictransport.Interface) *RenderSearchTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (r *RenderSearchTemplate) Raw(raw io.Reader) *RenderSearchTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (r *RenderSearchTemplate) Request(req *Request) *RenderSearchTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (r *RenderSearchTemplate) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r RenderSearchTemplate) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r RenderSearchTemplate) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *RenderSearchTemplate) Header(key, value string) *RenderSearchTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (r *RenderSearchTemplate) ErrorTrace(errortrace bool) *RenderSearchTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (r *RenderSearchTemplate) FilterPath(filterpaths ...string) *RenderSearchTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (r *RenderSearchTemplate) Human(human bool) *RenderSearchTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (r *RenderSearchTemplate) Pretty(pretty bool) *RenderSearchTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (r *RenderSearchTemplate) File(file string) *RenderSearchTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (r *RenderSearchTemplate) Id(id string) *RenderSearchTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (r *RenderSearchTemplate) Params(params map[string]json.RawMessage) *RenderSearchTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (r *RenderSearchTemplate) AddParam(key string, value json.RawMessage) *RenderSearchTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (r *RenderSearchTemplate) Source(scriptsource types.ScriptSourceVariant) *RenderSearchTemplate {
	_ = "STUB: not implemented"
	return nil
}
