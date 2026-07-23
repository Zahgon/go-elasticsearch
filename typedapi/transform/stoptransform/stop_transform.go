package stoptransform

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

type StopTransform struct {
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

type NewStopTransform func(transformid string) *StopTransform

func NewStopTransformFunc(tp elastictransport.Interface) NewStopTransform {
	_ = "STUB: not implemented"
	return *new(NewStopTransform)
}

func New(tp elastictransport.Interface) *StopTransform { _ = "STUB: not implemented"; return nil }

func (r *StopTransform) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r StopTransform) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r StopTransform) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r StopTransform) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *StopTransform) Header(key, value string) *StopTransform {
	_ = "STUB: not implemented"
	return nil
}

func (r *StopTransform) _transformid(transformid string) *StopTransform {
	_ = "STUB: not implemented"
	return nil
}

func (r *StopTransform) AllowNoMatch(allownomatch bool) *StopTransform {
	_ = "STUB: not implemented"
	return nil
}

func (r *StopTransform) Force(force bool) *StopTransform { _ = "STUB: not implemented"; return nil }

func (r *StopTransform) Timeout(duration string) *StopTransform {
	_ = "STUB: not implemented"
	return nil
}

func (r *StopTransform) WaitForCheckpoint(waitforcheckpoint bool) *StopTransform {
	_ = "STUB: not implemented"
	return nil
}

func (r *StopTransform) WaitForCompletion(waitforcompletion bool) *StopTransform {
	_ = "STUB: not implemented"
	return nil
}

func (r *StopTransform) ErrorTrace(errortrace bool) *StopTransform {
	_ = "STUB: not implemented"
	return nil
}

func (r *StopTransform) FilterPath(filterpaths ...string) *StopTransform {
	_ = "STUB: not implemented"
	return nil
}

func (r *StopTransform) Human(human bool) *StopTransform { _ = "STUB: not implemented"; return nil }

func (r *StopTransform) Pretty(pretty bool) *StopTransform { _ = "STUB: not implemented"; return nil }
