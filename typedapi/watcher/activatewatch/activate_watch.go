package activatewatch

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

type ActivateWatch struct {
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

type NewActivateWatch func(watchid string) *ActivateWatch

func NewActivateWatchFunc(tp elastictransport.Interface) NewActivateWatch {
	_ = "STUB: not implemented"
	return *new(NewActivateWatch)
}

func New(tp elastictransport.Interface) *ActivateWatch { _ = "STUB: not implemented"; return nil }

func (r *ActivateWatch) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r ActivateWatch) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r ActivateWatch) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r ActivateWatch) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *ActivateWatch) Header(key, value string) *ActivateWatch {
	_ = "STUB: not implemented"
	return nil
}

func (r *ActivateWatch) _watchid(watchid string) *ActivateWatch {
	_ = "STUB: not implemented"
	return nil
}

func (r *ActivateWatch) ErrorTrace(errortrace bool) *ActivateWatch {
	_ = "STUB: not implemented"
	return nil
}

func (r *ActivateWatch) FilterPath(filterpaths ...string) *ActivateWatch {
	_ = "STUB: not implemented"
	return nil
}

func (r *ActivateWatch) Human(human bool) *ActivateWatch { _ = "STUB: not implemented"; return nil }

func (r *ActivateWatch) Pretty(pretty bool) *ActivateWatch { _ = "STUB: not implemented"; return nil }
