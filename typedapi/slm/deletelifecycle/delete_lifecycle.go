package deletelifecycle

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
)

const (
	policyidMask = iota + 1
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type DeleteLifecycle struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	paramSet int

	policyid string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewDeleteLifecycle func(policyid string) *DeleteLifecycle

func NewDeleteLifecycleFunc(tp elastictransport.Interface) NewDeleteLifecycle {
	_ = "STUB: not implemented"
	return *new(NewDeleteLifecycle)
}

func New(tp elastictransport.Interface) *DeleteLifecycle { _ = "STUB: not implemented"; return nil }

func (r *DeleteLifecycle) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r DeleteLifecycle) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r DeleteLifecycle) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r DeleteLifecycle) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *DeleteLifecycle) Header(key, value string) *DeleteLifecycle {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteLifecycle) _policyid(policyid string) *DeleteLifecycle {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteLifecycle) MasterTimeout(duration string) *DeleteLifecycle {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteLifecycle) Timeout(duration string) *DeleteLifecycle {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteLifecycle) ErrorTrace(errortrace bool) *DeleteLifecycle {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteLifecycle) FilterPath(filterpaths ...string) *DeleteLifecycle {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteLifecycle) Human(human bool) *DeleteLifecycle { _ = "STUB: not implemented"; return nil }

func (r *DeleteLifecycle) Pretty(pretty bool) *DeleteLifecycle {
	_ = "STUB: not implemented"
	return nil
}
