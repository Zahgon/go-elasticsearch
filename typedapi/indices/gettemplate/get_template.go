package gettemplate

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

type GetTemplate struct {
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

type NewGetTemplate func() *GetTemplate

func NewGetTemplateFunc(tp elastictransport.Interface) NewGetTemplate {
	_ = "STUB: not implemented"
	return *new(NewGetTemplate)
}

func New(tp elastictransport.Interface) *GetTemplate { _ = "STUB: not implemented"; return nil }

func (r *GetTemplate) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetTemplate) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetTemplate) Do(providedCtx context.Context) (Response, error) {
	_ = "STUB: not implemented"
	return *new(Response), nil
}

func (r GetTemplate) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *GetTemplate) Header(key, value string) *GetTemplate { _ = "STUB: not implemented"; return nil }

func (r *GetTemplate) Name(name string) *GetTemplate { _ = "STUB: not implemented"; return nil }

func (r *GetTemplate) FlatSettings(flatsettings bool) *GetTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetTemplate) Local(local bool) *GetTemplate { _ = "STUB: not implemented"; return nil }

func (r *GetTemplate) MasterTimeout(duration string) *GetTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetTemplate) ErrorTrace(errortrace bool) *GetTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetTemplate) FilterPath(filterpaths ...string) *GetTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetTemplate) Human(human bool) *GetTemplate { _ = "STUB: not implemented"; return nil }

func (r *GetTemplate) Pretty(pretty bool) *GetTemplate { _ = "STUB: not implemented"; return nil }
