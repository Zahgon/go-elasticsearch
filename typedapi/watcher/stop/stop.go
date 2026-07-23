package stop

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type Stop struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	paramSet int

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewStop func() *Stop

func NewStopFunc(tp elastictransport.Interface) NewStop {
	_ = "STUB: not implemented"
	return *new(NewStop)
}

func New(tp elastictransport.Interface) *Stop { _ = "STUB: not implemented"; return nil }

func (r *Stop) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r Stop) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r Stop) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r Stop) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *Stop) Header(key, value string) *Stop { _ = "STUB: not implemented"; return nil }

func (r *Stop) MasterTimeout(duration string) *Stop { _ = "STUB: not implemented"; return nil }

func (r *Stop) ErrorTrace(errortrace bool) *Stop { _ = "STUB: not implemented"; return nil }

func (r *Stop) FilterPath(filterpaths ...string) *Stop { _ = "STUB: not implemented"; return nil }

func (r *Stop) Human(human bool) *Stop { _ = "STUB: not implemented"; return nil }

func (r *Stop) Pretty(pretty bool) *Stop { _ = "STUB: not implemented"; return nil }
