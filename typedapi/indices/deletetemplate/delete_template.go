package deletetemplate

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

type DeleteTemplate struct {
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

type NewDeleteTemplate func(name string) *DeleteTemplate

func NewDeleteTemplateFunc(tp elastictransport.Interface) NewDeleteTemplate {
	_ = "STUB: not implemented"
	return *new(NewDeleteTemplate)
}

func New(tp elastictransport.Interface) *DeleteTemplate { _ = "STUB: not implemented"; return nil }

func (r *DeleteTemplate) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r DeleteTemplate) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r DeleteTemplate) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r DeleteTemplate) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *DeleteTemplate) Header(key, value string) *DeleteTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteTemplate) _name(name string) *DeleteTemplate { _ = "STUB: not implemented"; return nil }

func (r *DeleteTemplate) MasterTimeout(duration string) *DeleteTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteTemplate) Timeout(duration string) *DeleteTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteTemplate) ErrorTrace(errortrace bool) *DeleteTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteTemplate) FilterPath(filterpaths ...string) *DeleteTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteTemplate) Human(human bool) *DeleteTemplate { _ = "STUB: not implemented"; return nil }

func (r *DeleteTemplate) Pretty(pretty bool) *DeleteTemplate { _ = "STUB: not implemented"; return nil }
