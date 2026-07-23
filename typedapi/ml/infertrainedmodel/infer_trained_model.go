package infertrainedmodel

import (
	gobytes "bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

const (
	modelidMask = iota + 1
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type InferTrainedModel struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	req      *Request
	deferred []func(request *Request) error
	buf      *gobytes.Buffer

	paramSet int

	modelid string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewInferTrainedModel func(modelid string) *InferTrainedModel

func NewInferTrainedModelFunc(tp elastictransport.Interface) NewInferTrainedModel {
	_ = "STUB: not implemented"
	return *new(NewInferTrainedModel)
}

func New(tp elastictransport.Interface) *InferTrainedModel { _ = "STUB: not implemented"; return nil }

func (r *InferTrainedModel) Raw(raw io.Reader) *InferTrainedModel {
	_ = "STUB: not implemented"
	return nil
}

func (r *InferTrainedModel) Request(req *Request) *InferTrainedModel {
	_ = "STUB: not implemented"
	return nil
}

func (r *InferTrainedModel) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r InferTrainedModel) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r InferTrainedModel) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *InferTrainedModel) Header(key, value string) *InferTrainedModel {
	_ = "STUB: not implemented"
	return nil
}

func (r *InferTrainedModel) _modelid(modelid string) *InferTrainedModel {
	_ = "STUB: not implemented"
	return nil
}

func (r *InferTrainedModel) Timeout(duration string) *InferTrainedModel {
	_ = "STUB: not implemented"
	return nil
}

func (r *InferTrainedModel) ErrorTrace(errortrace bool) *InferTrainedModel {
	_ = "STUB: not implemented"
	return nil
}

func (r *InferTrainedModel) FilterPath(filterpaths ...string) *InferTrainedModel {
	_ = "STUB: not implemented"
	return nil
}

func (r *InferTrainedModel) Human(human bool) *InferTrainedModel {
	_ = "STUB: not implemented"
	return nil
}

func (r *InferTrainedModel) Pretty(pretty bool) *InferTrainedModel {
	_ = "STUB: not implemented"
	return nil
}

func (r *InferTrainedModel) Docs(docs []map[string]json.RawMessage) *InferTrainedModel {
	_ = "STUB: not implemented"
	return nil
}

func (r *InferTrainedModel) InferenceConfig(inferenceconfig types.InferenceConfigUpdateContainerVariant) *InferTrainedModel {
	_ = "STUB: not implemented"
	return nil
}
