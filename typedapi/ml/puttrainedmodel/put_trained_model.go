package puttrainedmodel

import (
	gobytes "bytes"
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/trainedmodeltype"
)

const (
	modelidMask = iota + 1
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type PutTrainedModel struct {
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

type NewPutTrainedModel func(modelid string) *PutTrainedModel

func NewPutTrainedModelFunc(tp elastictransport.Interface) NewPutTrainedModel {
	_ = "STUB: not implemented"
	return *new(NewPutTrainedModel)
}

func New(tp elastictransport.Interface) *PutTrainedModel { _ = "STUB: not implemented"; return nil }

func (r *PutTrainedModel) Raw(raw io.Reader) *PutTrainedModel {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutTrainedModel) Request(req *Request) *PutTrainedModel {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutTrainedModel) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r PutTrainedModel) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r PutTrainedModel) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *PutTrainedModel) Header(key, value string) *PutTrainedModel {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutTrainedModel) _modelid(modelid string) *PutTrainedModel {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutTrainedModel) DeferDefinitionDecompression(deferdefinitiondecompression bool) *PutTrainedModel {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutTrainedModel) WaitForCompletion(waitforcompletion bool) *PutTrainedModel {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutTrainedModel) ErrorTrace(errortrace bool) *PutTrainedModel {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutTrainedModel) FilterPath(filterpaths ...string) *PutTrainedModel {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutTrainedModel) Human(human bool) *PutTrainedModel { _ = "STUB: not implemented"; return nil }

func (r *PutTrainedModel) Pretty(pretty bool) *PutTrainedModel {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutTrainedModel) CompressedDefinition(compresseddefinition string) *PutTrainedModel {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutTrainedModel) Definition(definition types.DefinitionVariant) *PutTrainedModel {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutTrainedModel) Description(description string) *PutTrainedModel {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutTrainedModel) InferenceConfig(inferenceconfig types.InferenceConfigCreateContainerVariant) *PutTrainedModel {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutTrainedModel) Input(input types.InputVariant) *PutTrainedModel {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutTrainedModel) Metadata(metadata any) *PutTrainedModel {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutTrainedModel) ModelSizeBytes(modelsizebytes int64) *PutTrainedModel {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutTrainedModel) ModelType(modeltype trainedmodeltype.TrainedModelType) *PutTrainedModel {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutTrainedModel) PlatformArchitecture(platformarchitecture string) *PutTrainedModel {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutTrainedModel) PrefixStrings(prefixstrings types.TrainedModelPrefixStringsVariant) *PutTrainedModel {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutTrainedModel) Tags(tags ...string) *PutTrainedModel {
	_ = "STUB: not implemented"
	return nil
}
