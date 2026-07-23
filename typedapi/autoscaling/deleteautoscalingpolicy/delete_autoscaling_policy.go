package deleteautoscalingpolicy

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

type DeleteAutoscalingPolicy struct {
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

type NewDeleteAutoscalingPolicy func(name string) *DeleteAutoscalingPolicy

func NewDeleteAutoscalingPolicyFunc(tp elastictransport.Interface) NewDeleteAutoscalingPolicy {
	_ = "STUB: not implemented"
	return *new(NewDeleteAutoscalingPolicy)
}

func New(tp elastictransport.Interface) *DeleteAutoscalingPolicy {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteAutoscalingPolicy) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r DeleteAutoscalingPolicy) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r DeleteAutoscalingPolicy) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r DeleteAutoscalingPolicy) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *DeleteAutoscalingPolicy) Header(key, value string) *DeleteAutoscalingPolicy {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteAutoscalingPolicy) _name(name string) *DeleteAutoscalingPolicy {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteAutoscalingPolicy) MasterTimeout(duration string) *DeleteAutoscalingPolicy {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteAutoscalingPolicy) Timeout(duration string) *DeleteAutoscalingPolicy {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteAutoscalingPolicy) ErrorTrace(errortrace bool) *DeleteAutoscalingPolicy {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteAutoscalingPolicy) FilterPath(filterpaths ...string) *DeleteAutoscalingPolicy {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteAutoscalingPolicy) Human(human bool) *DeleteAutoscalingPolicy {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteAutoscalingPolicy) Pretty(pretty bool) *DeleteAutoscalingPolicy {
	_ = "STUB: not implemented"
	return nil
}
