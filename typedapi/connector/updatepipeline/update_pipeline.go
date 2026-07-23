package updatepipeline

import (
	gobytes "bytes"
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

const (
	connectoridMask = iota + 1
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type UpdatePipeline struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	req      *Request
	deferred []func(request *Request) error
	buf      *gobytes.Buffer

	paramSet int

	connectorid string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewUpdatePipeline func(connectorid string) *UpdatePipeline

func NewUpdatePipelineFunc(tp elastictransport.Interface) NewUpdatePipeline {
	_ = "STUB: not implemented"
	return *new(NewUpdatePipeline)
}

func New(tp elastictransport.Interface) *UpdatePipeline { _ = "STUB: not implemented"; return nil }

func (r *UpdatePipeline) Raw(raw io.Reader) *UpdatePipeline { _ = "STUB: not implemented"; return nil }

func (r *UpdatePipeline) Request(req *Request) *UpdatePipeline {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdatePipeline) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r UpdatePipeline) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r UpdatePipeline) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *UpdatePipeline) Header(key, value string) *UpdatePipeline {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdatePipeline) _connectorid(connectorid string) *UpdatePipeline {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdatePipeline) ErrorTrace(errortrace bool) *UpdatePipeline {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdatePipeline) FilterPath(filterpaths ...string) *UpdatePipeline {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdatePipeline) Human(human bool) *UpdatePipeline { _ = "STUB: not implemented"; return nil }

func (r *UpdatePipeline) Pretty(pretty bool) *UpdatePipeline { _ = "STUB: not implemented"; return nil }

func (r *UpdatePipeline) Pipeline(pipeline types.IngestPipelineParamsVariant) *UpdatePipeline {
	_ = "STUB: not implemented"
	return nil
}
