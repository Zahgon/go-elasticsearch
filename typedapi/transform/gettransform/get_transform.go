package gettransform

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

type GetTransform struct {
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

type NewGetTransform func() *GetTransform

func NewGetTransformFunc(tp elastictransport.Interface) NewGetTransform {
	_ = "STUB: not implemented"
	return *new(NewGetTransform)
}

func New(tp elastictransport.Interface) *GetTransform { _ = "STUB: not implemented"; return nil }

func (r *GetTransform) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetTransform) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetTransform) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetTransform) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *GetTransform) Header(key, value string) *GetTransform {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetTransform) TransformId(transformid string) *GetTransform {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetTransform) AllowNoMatch(allownomatch bool) *GetTransform {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetTransform) From(from int) *GetTransform { _ = "STUB: not implemented"; return nil }

func (r *GetTransform) Size(size int) *GetTransform { _ = "STUB: not implemented"; return nil }

func (r *GetTransform) ExcludeGenerated(excludegenerated bool) *GetTransform {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetTransform) ErrorTrace(errortrace bool) *GetTransform {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetTransform) FilterPath(filterpaths ...string) *GetTransform {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetTransform) Human(human bool) *GetTransform { _ = "STUB: not implemented"; return nil }

func (r *GetTransform) Pretty(pretty bool) *GetTransform { _ = "STUB: not implemented"; return nil }
