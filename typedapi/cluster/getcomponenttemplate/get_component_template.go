package getcomponenttemplate

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

type GetComponentTemplate struct {
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

type NewGetComponentTemplate func() *GetComponentTemplate

func NewGetComponentTemplateFunc(tp elastictransport.Interface) NewGetComponentTemplate {
	_ = "STUB: not implemented"
	return *new(NewGetComponentTemplate)
}

func New(tp elastictransport.Interface) *GetComponentTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetComponentTemplate) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetComponentTemplate) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetComponentTemplate) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetComponentTemplate) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *GetComponentTemplate) Header(key, value string) *GetComponentTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetComponentTemplate) Name(name string) *GetComponentTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetComponentTemplate) FlatSettings(flatsettings bool) *GetComponentTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetComponentTemplate) SettingsFilter(settingsfilters ...string) *GetComponentTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetComponentTemplate) IncludeDefaults(includedefaults bool) *GetComponentTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetComponentTemplate) Local(local bool) *GetComponentTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetComponentTemplate) MasterTimeout(duration string) *GetComponentTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetComponentTemplate) ErrorTrace(errortrace bool) *GetComponentTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetComponentTemplate) FilterPath(filterpaths ...string) *GetComponentTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetComponentTemplate) Human(human bool) *GetComponentTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetComponentTemplate) Pretty(pretty bool) *GetComponentTemplate {
	_ = "STUB: not implemented"
	return nil
}
