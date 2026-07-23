package existsindextemplate

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

type ExistsIndexTemplate struct {
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

type NewExistsIndexTemplate func(name string) *ExistsIndexTemplate

func NewExistsIndexTemplateFunc(tp elastictransport.Interface) NewExistsIndexTemplate {
	_ = "STUB: not implemented"
	return *new(NewExistsIndexTemplate)
}

func New(tp elastictransport.Interface) *ExistsIndexTemplate { _ = "STUB: not implemented"; return nil }

func (r *ExistsIndexTemplate) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r ExistsIndexTemplate) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r ExistsIndexTemplate) Do(ctx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r ExistsIndexTemplate) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *ExistsIndexTemplate) Header(key, value string) *ExistsIndexTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (r *ExistsIndexTemplate) _name(name string) *ExistsIndexTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (r *ExistsIndexTemplate) Local(local bool) *ExistsIndexTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (r *ExistsIndexTemplate) FlatSettings(flatsettings bool) *ExistsIndexTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (r *ExistsIndexTemplate) MasterTimeout(duration string) *ExistsIndexTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (r *ExistsIndexTemplate) ErrorTrace(errortrace bool) *ExistsIndexTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (r *ExistsIndexTemplate) FilterPath(filterpaths ...string) *ExistsIndexTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (r *ExistsIndexTemplate) Human(human bool) *ExistsIndexTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (r *ExistsIndexTemplate) Pretty(pretty bool) *ExistsIndexTemplate {
	_ = "STUB: not implemented"
	return nil
}
