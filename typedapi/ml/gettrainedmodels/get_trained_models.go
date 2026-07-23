package gettrainedmodels

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/include"
)

const (
	modelidMask = iota + 1
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type GetTrainedModels struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	paramSet int

	modelid string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewGetTrainedModels func() *GetTrainedModels

func NewGetTrainedModelsFunc(tp elastictransport.Interface) NewGetTrainedModels {
	_ = "STUB: not implemented"
	return *new(NewGetTrainedModels)
}

func New(tp elastictransport.Interface) *GetTrainedModels { _ = "STUB: not implemented"; return nil }

func (r *GetTrainedModels) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetTrainedModels) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetTrainedModels) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r GetTrainedModels) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *GetTrainedModels) Header(key, value string) *GetTrainedModels {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetTrainedModels) ModelId(modelid string) *GetTrainedModels {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetTrainedModels) AllowNoMatch(allownomatch bool) *GetTrainedModels {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetTrainedModels) DecompressDefinition(decompressdefinition bool) *GetTrainedModels {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetTrainedModels) ExcludeGenerated(excludegenerated bool) *GetTrainedModels {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetTrainedModels) From(from int) *GetTrainedModels { _ = "STUB: not implemented"; return nil }

func (r *GetTrainedModels) Include(include include.Include) *GetTrainedModels {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetTrainedModels) Size(size int) *GetTrainedModels { _ = "STUB: not implemented"; return nil }

func (r *GetTrainedModels) Tags(tags ...string) *GetTrainedModels {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetTrainedModels) ErrorTrace(errortrace bool) *GetTrainedModels {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetTrainedModels) FilterPath(filterpaths ...string) *GetTrainedModels {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetTrainedModels) Human(human bool) *GetTrainedModels {
	_ = "STUB: not implemented"
	return nil
}

func (r *GetTrainedModels) Pretty(pretty bool) *GetTrainedModels {
	_ = "STUB: not implemented"
	return nil
}
