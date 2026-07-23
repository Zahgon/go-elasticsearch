package deactivatewatch

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
)

const (
	watchidMask = iota + 1
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type DeactivateWatch struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	paramSet int

	watchid string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewDeactivateWatch func(watchid string) *DeactivateWatch

func NewDeactivateWatchFunc(tp elastictransport.Interface) NewDeactivateWatch {
	_ = "STUB: not implemented"
	return *new(NewDeactivateWatch)
}

func New(tp elastictransport.Interface) *DeactivateWatch { _ = "STUB: not implemented"; return nil }

func (r *DeactivateWatch) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r DeactivateWatch) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r DeactivateWatch) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r DeactivateWatch) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *DeactivateWatch) Header(key, value string) *DeactivateWatch {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeactivateWatch) _watchid(watchid string) *DeactivateWatch {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeactivateWatch) ErrorTrace(errortrace bool) *DeactivateWatch {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeactivateWatch) FilterPath(filterpaths ...string) *DeactivateWatch {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeactivateWatch) Human(human bool) *DeactivateWatch { _ = "STUB: not implemented"; return nil }

func (r *DeactivateWatch) Pretty(pretty bool) *DeactivateWatch {
	_ = "STUB: not implemented"
	return nil
}
