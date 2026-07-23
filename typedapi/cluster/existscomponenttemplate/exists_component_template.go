package existscomponenttemplate

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

type ExistsComponentTemplate struct {
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

type NewExistsComponentTemplate func(name string) *ExistsComponentTemplate

func NewExistsComponentTemplateFunc(tp elastictransport.Interface) NewExistsComponentTemplate {
	_ = "STUB: not implemented"
	return *new(NewExistsComponentTemplate)
}

func New(tp elastictransport.Interface) *ExistsComponentTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (r *ExistsComponentTemplate) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r ExistsComponentTemplate) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r ExistsComponentTemplate) Do(ctx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r ExistsComponentTemplate) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *ExistsComponentTemplate) Header(key, value string) *ExistsComponentTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (r *ExistsComponentTemplate) _name(name string) *ExistsComponentTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (r *ExistsComponentTemplate) MasterTimeout(duration string) *ExistsComponentTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (r *ExistsComponentTemplate) Local(local bool) *ExistsComponentTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (r *ExistsComponentTemplate) ErrorTrace(errortrace bool) *ExistsComponentTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (r *ExistsComponentTemplate) FilterPath(filterpaths ...string) *ExistsComponentTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (r *ExistsComponentTemplate) Human(human bool) *ExistsComponentTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (r *ExistsComponentTemplate) Pretty(pretty bool) *ExistsComponentTemplate {
	_ = "STUB: not implemented"
	return nil
}
