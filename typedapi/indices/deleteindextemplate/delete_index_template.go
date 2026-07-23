package deleteindextemplate

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

type DeleteIndexTemplate struct {
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

type NewDeleteIndexTemplate func(name string) *DeleteIndexTemplate

func NewDeleteIndexTemplateFunc(tp elastictransport.Interface) NewDeleteIndexTemplate {
	_ = "STUB: not implemented"
	return *new(NewDeleteIndexTemplate)
}

func New(tp elastictransport.Interface) *DeleteIndexTemplate { _ = "STUB: not implemented"; return nil }

func (r *DeleteIndexTemplate) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r DeleteIndexTemplate) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r DeleteIndexTemplate) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r DeleteIndexTemplate) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *DeleteIndexTemplate) Header(key, value string) *DeleteIndexTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteIndexTemplate) _name(name string) *DeleteIndexTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteIndexTemplate) MasterTimeout(duration string) *DeleteIndexTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteIndexTemplate) Timeout(duration string) *DeleteIndexTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteIndexTemplate) ErrorTrace(errortrace bool) *DeleteIndexTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteIndexTemplate) FilterPath(filterpaths ...string) *DeleteIndexTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteIndexTemplate) Human(human bool) *DeleteIndexTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteIndexTemplate) Pretty(pretty bool) *DeleteIndexTemplate {
	_ = "STUB: not implemented"
	return nil
}
