package mltrainedmodels

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/bytes"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/cattrainedmodelscolumn"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/timeunit"
)

const (
	modelidMask = iota + 1
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type MlTrainedModels struct {
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

type NewMlTrainedModels func() *MlTrainedModels

func NewMlTrainedModelsFunc(tp elastictransport.Interface) NewMlTrainedModels {
	_ = "STUB: not implemented"
	return *new(NewMlTrainedModels)
}

func New(tp elastictransport.Interface) *MlTrainedModels { _ = "STUB: not implemented"; return nil }

func (r *MlTrainedModels) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r MlTrainedModels) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r MlTrainedModels) Do(providedCtx context.Context) (Response, error) {
	_ = "STUB: not implemented"
	return *new(Response), nil
}

func (r MlTrainedModels) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *MlTrainedModels) Header(key, value string) *MlTrainedModels {
	_ = "STUB: not implemented"
	return nil
}

func (r *MlTrainedModels) ModelId(modelid string) *MlTrainedModels {
	_ = "STUB: not implemented"
	return nil
}

func (r *MlTrainedModels) AllowNoMatch(allownomatch bool) *MlTrainedModels {
	_ = "STUB: not implemented"
	return nil
}

func (r *MlTrainedModels) H(cattrainedmodelscolumns ...cattrainedmodelscolumn.CatTrainedModelsColumn) *MlTrainedModels {
	_ = "STUB: not implemented"
	return nil
}

func (r *MlTrainedModels) S(cattrainedmodelscolumns ...cattrainedmodelscolumn.CatTrainedModelsColumn) *MlTrainedModels {
	_ = "STUB: not implemented"
	return nil
}

func (r *MlTrainedModels) From(from int) *MlTrainedModels { _ = "STUB: not implemented"; return nil }

func (r *MlTrainedModels) Size(size int) *MlTrainedModels { _ = "STUB: not implemented"; return nil }

func (r *MlTrainedModels) Bytes(bytes bytes.Bytes) *MlTrainedModels {
	_ = "STUB: not implemented"
	return nil
}

func (r *MlTrainedModels) Format(format string) *MlTrainedModels {
	_ = "STUB: not implemented"
	return nil
}

func (r *MlTrainedModels) Help(help bool) *MlTrainedModels { _ = "STUB: not implemented"; return nil }

func (r *MlTrainedModels) Time(time timeunit.TimeUnit) *MlTrainedModels {
	_ = "STUB: not implemented"
	return nil
}

func (r *MlTrainedModels) V(v bool) *MlTrainedModels { _ = "STUB: not implemented"; return nil }

func (r *MlTrainedModels) ErrorTrace(errortrace bool) *MlTrainedModels {
	_ = "STUB: not implemented"
	return nil
}

func (r *MlTrainedModels) FilterPath(filterpaths ...string) *MlTrainedModels {
	_ = "STUB: not implemented"
	return nil
}

func (r *MlTrainedModels) Human(human bool) *MlTrainedModels { _ = "STUB: not implemented"; return nil }

func (r *MlTrainedModels) Pretty(pretty bool) *MlTrainedModels {
	_ = "STUB: not implemented"
	return nil
}
