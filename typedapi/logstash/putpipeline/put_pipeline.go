package putpipeline

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
	idMask = iota + 1
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type PutPipeline struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	req      *Request
	deferred []func(request *Request) error
	buf      *gobytes.Buffer

	paramSet int

	id string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewPutPipeline func(id string) *PutPipeline

func NewPutPipelineFunc(tp elastictransport.Interface) NewPutPipeline {
	_ = "STUB: not implemented"
	return *new(NewPutPipeline)
}

func New(tp elastictransport.Interface) *PutPipeline { _ = "STUB: not implemented"; return nil }

func (r *PutPipeline) Raw(raw io.Reader) *PutPipeline { _ = "STUB: not implemented"; return nil }

func (r *PutPipeline) Request(req *Request) *PutPipeline { _ = "STUB: not implemented"; return nil }

func (r *PutPipeline) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r PutPipeline) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *PutPipeline) Header(key, value string) *PutPipeline { _ = "STUB: not implemented"; return nil }

func (r *PutPipeline) _id(id string) *PutPipeline { _ = "STUB: not implemented"; return nil }

func (r *PutPipeline) ErrorTrace(errortrace bool) *PutPipeline {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutPipeline) FilterPath(filterpaths ...string) *PutPipeline {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutPipeline) Human(human bool) *PutPipeline { _ = "STUB: not implemented"; return nil }

func (r *PutPipeline) Pretty(pretty bool) *PutPipeline { _ = "STUB: not implemented"; return nil }

func (r *PutPipeline) Description(description string) *PutPipeline {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutPipeline) LastModified(datetime types.DateTimeVariant) *PutPipeline {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutPipeline) Pipeline(pipeline string) *PutPipeline { _ = "STUB: not implemented"; return nil }

func (r *PutPipeline) PipelineMetadata(pipelinemetadata types.PipelineMetadataVariant) *PutPipeline {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutPipeline) PipelineSettings(pipelinesettings types.PipelineSettingsVariant) *PutPipeline {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutPipeline) Username(username string) *PutPipeline { _ = "STUB: not implemented"; return nil }
