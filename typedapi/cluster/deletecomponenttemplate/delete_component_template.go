package deletecomponenttemplate

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

type DeleteComponentTemplate struct {
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

type NewDeleteComponentTemplate func(name string) *DeleteComponentTemplate

func NewDeleteComponentTemplateFunc(tp elastictransport.Interface) NewDeleteComponentTemplate {
	_ = "STUB: not implemented"
	return *new(NewDeleteComponentTemplate)
}

func New(tp elastictransport.Interface) *DeleteComponentTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteComponentTemplate) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r DeleteComponentTemplate) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r DeleteComponentTemplate) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r DeleteComponentTemplate) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *DeleteComponentTemplate) Header(key, value string) *DeleteComponentTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteComponentTemplate) _name(name string) *DeleteComponentTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteComponentTemplate) MasterTimeout(duration string) *DeleteComponentTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteComponentTemplate) Timeout(duration string) *DeleteComponentTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteComponentTemplate) ErrorTrace(errortrace bool) *DeleteComponentTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteComponentTemplate) FilterPath(filterpaths ...string) *DeleteComponentTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteComponentTemplate) Human(human bool) *DeleteComponentTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteComponentTemplate) Pretty(pretty bool) *DeleteComponentTemplate {
	_ = "STUB: not implemented"
	return nil
}
