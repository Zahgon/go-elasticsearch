package executeretention

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type ExecuteRetention struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	paramSet int

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewExecuteRetention func() *ExecuteRetention

func NewExecuteRetentionFunc(tp elastictransport.Interface) NewExecuteRetention {
	_ = "STUB: not implemented"
	return *new(NewExecuteRetention)
}

func New(tp elastictransport.Interface) *ExecuteRetention { _ = "STUB: not implemented"; return nil }

func (r *ExecuteRetention) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r ExecuteRetention) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r ExecuteRetention) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r ExecuteRetention) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *ExecuteRetention) Header(key, value string) *ExecuteRetention {
	_ = "STUB: not implemented"
	return nil
}

func (r *ExecuteRetention) MasterTimeout(duration string) *ExecuteRetention {
	_ = "STUB: not implemented"
	return nil
}

func (r *ExecuteRetention) Timeout(duration string) *ExecuteRetention {
	_ = "STUB: not implemented"
	return nil
}

func (r *ExecuteRetention) ErrorTrace(errortrace bool) *ExecuteRetention {
	_ = "STUB: not implemented"
	return nil
}

func (r *ExecuteRetention) FilterPath(filterpaths ...string) *ExecuteRetention {
	_ = "STUB: not implemented"
	return nil
}

func (r *ExecuteRetention) Human(human bool) *ExecuteRetention {
	_ = "STUB: not implemented"
	return nil
}

func (r *ExecuteRetention) Pretty(pretty bool) *ExecuteRetention {
	_ = "STUB: not implemented"
	return nil
}
