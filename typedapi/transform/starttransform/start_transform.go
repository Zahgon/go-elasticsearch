package starttransform

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

type StartTransform struct {
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

type NewStartTransform func(transformid string) *StartTransform

func NewStartTransformFunc(tp elastictransport.Interface) NewStartTransform {
	_ = "STUB: not implemented"
	return *new(NewStartTransform)
}

func New(tp elastictransport.Interface) *StartTransform { _ = "STUB: not implemented"; return nil }

func (r *StartTransform) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r StartTransform) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r StartTransform) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r StartTransform) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *StartTransform) Header(key, value string) *StartTransform {
	_ = "STUB: not implemented"
	return nil
}

func (r *StartTransform) _transformid(transformid string) *StartTransform {
	_ = "STUB: not implemented"
	return nil
}

func (r *StartTransform) Timeout(duration string) *StartTransform {
	_ = "STUB: not implemented"
	return nil
}

func (r *StartTransform) From(from string) *StartTransform { _ = "STUB: not implemented"; return nil }

func (r *StartTransform) ErrorTrace(errortrace bool) *StartTransform {
	_ = "STUB: not implemented"
	return nil
}

func (r *StartTransform) FilterPath(filterpaths ...string) *StartTransform {
	_ = "STUB: not implemented"
	return nil
}

func (r *StartTransform) Human(human bool) *StartTransform { _ = "STUB: not implemented"; return nil }

func (r *StartTransform) Pretty(pretty bool) *StartTransform { _ = "STUB: not implemented"; return nil }
