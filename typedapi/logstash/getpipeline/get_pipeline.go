package getpipeline

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
)

const (
	idMask = iota + 1
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type GetPipeline struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	paramSet int

	id string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewGetPipeline func() *GetPipeline

func NewGetPipelineFunc(tp elastictransport.Interface) NewGetPipeline {
	_ = "STUB: not implemented"
	return *new(NewGetPipeline)
}

func New(tp elastictransport.Interface) *GetPipeline { _ = "STUB: not implemented"; return nil }

func (r *GetPipeline) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetPipeline) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetPipeline) Do(providedCtx context.Context) (Response, error) {
	_ = "STUB: not implemented"
	return *new(Response), nil
}

func (r GetPipeline) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *GetPipeline) Header(key, value string) *GetPipeline { _ = "STUB: not implemented"; return nil }

func (r *GetPipeline) Id(id string) *GetPipeline { _ = "STUB: not implemented"; return nil }

func (r *GetPipeline) ErrorTrace(errortrace bool) *GetPipeline {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetPipeline) FilterPath(filterpaths ...string) *GetPipeline {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetPipeline) Human(human bool) *GetPipeline { _ = "STUB: not implemented"; return nil }

func (r *GetPipeline) Pretty(pretty bool) *GetPipeline { _ = "STUB: not implemented"; return nil }
