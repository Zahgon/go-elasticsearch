package getautoscalingpolicy

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

type GetAutoscalingPolicy struct {
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

type NewGetAutoscalingPolicy func(name string) *GetAutoscalingPolicy

func NewGetAutoscalingPolicyFunc(tp elastictransport.Interface) NewGetAutoscalingPolicy {
	_ = "STUB: not implemented"
	return *new(NewGetAutoscalingPolicy)
}

func New(tp elastictransport.Interface) *GetAutoscalingPolicy {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetAutoscalingPolicy) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetAutoscalingPolicy) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetAutoscalingPolicy) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetAutoscalingPolicy) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *GetAutoscalingPolicy) Header(key, value string) *GetAutoscalingPolicy {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetAutoscalingPolicy) _name(name string) *GetAutoscalingPolicy {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetAutoscalingPolicy) MasterTimeout(duration string) *GetAutoscalingPolicy {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetAutoscalingPolicy) ErrorTrace(errortrace bool) *GetAutoscalingPolicy {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetAutoscalingPolicy) FilterPath(filterpaths ...string) *GetAutoscalingPolicy {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetAutoscalingPolicy) Human(human bool) *GetAutoscalingPolicy {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetAutoscalingPolicy) Pretty(pretty bool) *GetAutoscalingPolicy {
	_ = "STUB: not implemented"
	return nil
}
