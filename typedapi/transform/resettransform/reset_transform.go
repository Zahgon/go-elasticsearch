package resettransform

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

type ResetTransform struct {
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

type NewResetTransform func(transformid string) *ResetTransform

func NewResetTransformFunc(tp elastictransport.Interface) NewResetTransform {
	_ = "STUB: not implemented"
	return *new(NewResetTransform)
}

func New(tp elastictransport.Interface) *ResetTransform { _ = "STUB: not implemented"; return nil }

func (r *ResetTransform) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r ResetTransform) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r ResetTransform) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r ResetTransform) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *ResetTransform) Header(key, value string) *ResetTransform {
	_ = "STUB: not implemented"
	return nil
}

func (r *ResetTransform) _transformid(transformid string) *ResetTransform {
	_ = "STUB: not implemented"
	return nil
}

func (r *ResetTransform) Force(force bool) *ResetTransform { _ = "STUB: not implemented"; return nil }

func (r *ResetTransform) Timeout(duration string) *ResetTransform {
	_ = "STUB: not implemented"
	return nil
}

func (r *ResetTransform) ErrorTrace(errortrace bool) *ResetTransform {
	_ = "STUB: not implemented"
	return nil
}

func (r *ResetTransform) FilterPath(filterpaths ...string) *ResetTransform {
	_ = "STUB: not implemented"
	return nil
}

func (r *ResetTransform) Human(human bool) *ResetTransform { _ = "STUB: not implemented"; return nil }

func (r *ResetTransform) Pretty(pretty bool) *ResetTransform { _ = "STUB: not implemented"; return nil }
