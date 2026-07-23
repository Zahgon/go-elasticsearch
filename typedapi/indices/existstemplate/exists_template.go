package existstemplate

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

type ExistsTemplate struct {
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

type NewExistsTemplate func(name string) *ExistsTemplate

func NewExistsTemplateFunc(tp elastictransport.Interface) NewExistsTemplate {
	_ = "STUB: not implemented"
	return *new(NewExistsTemplate)
}

func New(tp elastictransport.Interface) *ExistsTemplate { _ = "STUB: not implemented"; return nil }

func (r *ExistsTemplate) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r ExistsTemplate) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r ExistsTemplate) Do(ctx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r ExistsTemplate) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *ExistsTemplate) Header(key, value string) *ExistsTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (r *ExistsTemplate) _name(name string) *ExistsTemplate { _ = "STUB: not implemented"; return nil }

func (r *ExistsTemplate) FlatSettings(flatsettings bool) *ExistsTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (r *ExistsTemplate) Local(local bool) *ExistsTemplate { _ = "STUB: not implemented"; return nil }

func (r *ExistsTemplate) MasterTimeout(duration string) *ExistsTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (r *ExistsTemplate) ErrorTrace(errortrace bool) *ExistsTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (r *ExistsTemplate) FilterPath(filterpaths ...string) *ExistsTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (r *ExistsTemplate) Human(human bool) *ExistsTemplate { _ = "STUB: not implemented"; return nil }

func (r *ExistsTemplate) Pretty(pretty bool) *ExistsTemplate { _ = "STUB: not implemented"; return nil }
