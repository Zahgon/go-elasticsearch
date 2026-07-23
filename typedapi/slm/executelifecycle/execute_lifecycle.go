package executelifecycle

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

type ExecuteLifecycle struct {
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

type NewExecuteLifecycle func(policyid string) *ExecuteLifecycle

func NewExecuteLifecycleFunc(tp elastictransport.Interface) NewExecuteLifecycle {
	_ = "STUB: not implemented"
	return *new(NewExecuteLifecycle)
}

func New(tp elastictransport.Interface) *ExecuteLifecycle { _ = "STUB: not implemented"; return nil }

func (r *ExecuteLifecycle) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r ExecuteLifecycle) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r ExecuteLifecycle) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r ExecuteLifecycle) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *ExecuteLifecycle) Header(key, value string) *ExecuteLifecycle {
	_ = "STUB: not implemented"
	return nil
}

func (r *ExecuteLifecycle) _policyid(policyid string) *ExecuteLifecycle {
	_ = "STUB: not implemented"
	return nil
}

func (r *ExecuteLifecycle) MasterTimeout(duration string) *ExecuteLifecycle {
	_ = "STUB: not implemented"
	return nil
}

func (r *ExecuteLifecycle) Timeout(duration string) *ExecuteLifecycle {
	_ = "STUB: not implemented"
	return nil
}

func (r *ExecuteLifecycle) ErrorTrace(errortrace bool) *ExecuteLifecycle {
	_ = "STUB: not implemented"
	return nil
}

func (r *ExecuteLifecycle) FilterPath(filterpaths ...string) *ExecuteLifecycle {
	_ = "STUB: not implemented"
	return nil
}

func (r *ExecuteLifecycle) Human(human bool) *ExecuteLifecycle {
	_ = "STUB: not implemented"
	return nil
}

func (r *ExecuteLifecycle) Pretty(pretty bool) *ExecuteLifecycle {
	_ = "STUB: not implemented"
	return nil
}
