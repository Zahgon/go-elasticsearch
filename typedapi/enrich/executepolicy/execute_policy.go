package executepolicy

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
)

const (
	nameMask = iota + 1
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type ExecutePolicy struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	paramSet int

	name string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewExecutePolicy func(name string) *ExecutePolicy

func NewExecutePolicyFunc(tp elastictransport.Interface) NewExecutePolicy {
	_ = "STUB: not implemented"
	return *new(NewExecutePolicy)
}

func New(tp elastictransport.Interface) *ExecutePolicy { _ = "STUB: not implemented"; return nil }

func (r *ExecutePolicy) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r ExecutePolicy) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r ExecutePolicy) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r ExecutePolicy) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *ExecutePolicy) Header(key, value string) *ExecutePolicy {
	_ = "STUB: not implemented"
	return nil
}

func (r *ExecutePolicy) _name(name string) *ExecutePolicy { _ = "STUB: not implemented"; return nil }

func (r *ExecutePolicy) MasterTimeout(duration string) *ExecutePolicy {
	_ = "STUB: not implemented"
	return nil
}

func (r *ExecutePolicy) WaitForCompletion(waitforcompletion bool) *ExecutePolicy {
	_ = "STUB: not implemented"
	return nil
}

func (r *ExecutePolicy) ErrorTrace(errortrace bool) *ExecutePolicy {
	_ = "STUB: not implemented"
	return nil
}

func (r *ExecutePolicy) FilterPath(filterpaths ...string) *ExecutePolicy {
	_ = "STUB: not implemented"
	return nil
}

func (r *ExecutePolicy) Human(human bool) *ExecutePolicy { _ = "STUB: not implemented"; return nil }

func (r *ExecutePolicy) Pretty(pretty bool) *ExecutePolicy { _ = "STUB: not implemented"; return nil }
