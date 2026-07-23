package deletescript

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
)

const (
	idMask = iota + 1
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type DeleteScript struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	paramSet int

	id string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewDeleteScript func(id string) *DeleteScript

func NewDeleteScriptFunc(tp elastictransport.Interface) NewDeleteScript {
	_ = "STUB: not implemented"
	return *new(NewDeleteScript)
}

func New(tp elastictransport.Interface) *DeleteScript { _ = "STUB: not implemented"; return nil }

func (r *DeleteScript) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r DeleteScript) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r DeleteScript) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r DeleteScript) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *DeleteScript) Header(key, value string) *DeleteScript {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteScript) _id(id string) *DeleteScript { _ = "STUB: not implemented"; return nil }

func (r *DeleteScript) MasterTimeout(duration string) *DeleteScript {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteScript) Timeout(duration string) *DeleteScript {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteScript) ErrorTrace(errortrace bool) *DeleteScript {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteScript) FilterPath(filterpaths ...string) *DeleteScript {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteScript) Human(human bool) *DeleteScript { _ = "STUB: not implemented"; return nil }

func (r *DeleteScript) Pretty(pretty bool) *DeleteScript { _ = "STUB: not implemented"; return nil }
