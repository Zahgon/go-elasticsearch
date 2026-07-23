package getautoscalingcapacity

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type GetAutoscalingCapacity struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	paramSet int

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewGetAutoscalingCapacity func() *GetAutoscalingCapacity

func NewGetAutoscalingCapacityFunc(tp elastictransport.Interface) NewGetAutoscalingCapacity {
	_ = "STUB: not implemented"
	return *new(NewGetAutoscalingCapacity)
}

func New(tp elastictransport.Interface) *GetAutoscalingCapacity {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetAutoscalingCapacity) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetAutoscalingCapacity) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetAutoscalingCapacity) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetAutoscalingCapacity) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *GetAutoscalingCapacity) Header(key, value string) *GetAutoscalingCapacity {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetAutoscalingCapacity) MasterTimeout(duration string) *GetAutoscalingCapacity {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetAutoscalingCapacity) ErrorTrace(errortrace bool) *GetAutoscalingCapacity {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetAutoscalingCapacity) FilterPath(filterpaths ...string) *GetAutoscalingCapacity {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetAutoscalingCapacity) Human(human bool) *GetAutoscalingCapacity {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetAutoscalingCapacity) Pretty(pretty bool) *GetAutoscalingCapacity {
	_ = "STUB: not implemented"
	return nil
}
