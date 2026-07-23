package getindextemplate

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
)

const (
	nameMask = iota + 1
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type GetIndexTemplate struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	paramSet int

	name string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewGetIndexTemplate func() *GetIndexTemplate

func NewGetIndexTemplateFunc(tp elastictransport.Interface) NewGetIndexTemplate {
	_ = "STUB: not implemented"
	return *new(NewGetIndexTemplate)
}

func New(tp elastictransport.Interface) *GetIndexTemplate { _ = "STUB: not implemented"; return nil }

func (r *GetIndexTemplate) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetIndexTemplate) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetIndexTemplate) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetIndexTemplate) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *GetIndexTemplate) Header(key, value string) *GetIndexTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetIndexTemplate) Name(name string) *GetIndexTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetIndexTemplate) Local(local bool) *GetIndexTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetIndexTemplate) FlatSettings(flatsettings bool) *GetIndexTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetIndexTemplate) MasterTimeout(duration string) *GetIndexTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetIndexTemplate) IncludeDefaults(includedefaults bool) *GetIndexTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetIndexTemplate) ErrorTrace(errortrace bool) *GetIndexTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetIndexTemplate) FilterPath(filterpaths ...string) *GetIndexTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetIndexTemplate) Human(human bool) *GetIndexTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetIndexTemplate) Pretty(pretty bool) *GetIndexTemplate {
	_ = "STUB: not implemented"
	return nil
}
