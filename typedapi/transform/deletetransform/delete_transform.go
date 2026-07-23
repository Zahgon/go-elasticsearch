package deletetransform

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
)

const (
	transformidMask = iota + 1
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type DeleteTransform struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	paramSet int

	transformid string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewDeleteTransform func(transformid string) *DeleteTransform

func NewDeleteTransformFunc(tp elastictransport.Interface) NewDeleteTransform {
	_ = "STUB: not implemented"
	return *new(NewDeleteTransform)
}

func New(tp elastictransport.Interface) *DeleteTransform { _ = "STUB: not implemented"; return nil }

func (r *DeleteTransform) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r DeleteTransform) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r DeleteTransform) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r DeleteTransform) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *DeleteTransform) Header(key, value string) *DeleteTransform {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteTransform) _transformid(transformid string) *DeleteTransform {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteTransform) Force(force bool) *DeleteTransform { _ = "STUB: not implemented"; return nil }

func (r *DeleteTransform) DeleteDestIndex(deletedestindex bool) *DeleteTransform {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteTransform) Timeout(duration string) *DeleteTransform {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteTransform) ErrorTrace(errortrace bool) *DeleteTransform {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteTransform) FilterPath(filterpaths ...string) *DeleteTransform {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteTransform) Human(human bool) *DeleteTransform { _ = "STUB: not implemented"; return nil }

func (r *DeleteTransform) Pretty(pretty bool) *DeleteTransform {
	_ = "STUB: not implemented"
	return nil
}
