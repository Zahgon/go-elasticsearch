package deletepipeline

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

type DeletePipeline struct {
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

type NewDeletePipeline func(id string) *DeletePipeline

func NewDeletePipelineFunc(tp elastictransport.Interface) NewDeletePipeline {
	_ = "STUB: not implemented"
	return *new(NewDeletePipeline)
}

func New(tp elastictransport.Interface) *DeletePipeline { _ = "STUB: not implemented"; return nil }

func (r *DeletePipeline) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r DeletePipeline) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r DeletePipeline) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r DeletePipeline) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *DeletePipeline) Header(key, value string) *DeletePipeline {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeletePipeline) _id(id string) *DeletePipeline { _ = "STUB: not implemented"; return nil }

func (r *DeletePipeline) MasterTimeout(duration string) *DeletePipeline {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeletePipeline) Timeout(duration string) *DeletePipeline {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeletePipeline) ErrorTrace(errortrace bool) *DeletePipeline {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeletePipeline) FilterPath(filterpaths ...string) *DeletePipeline {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeletePipeline) Human(human bool) *DeletePipeline { _ = "STUB: not implemented"; return nil }

func (r *DeletePipeline) Pretty(pretty bool) *DeletePipeline { _ = "STUB: not implemented"; return nil }
