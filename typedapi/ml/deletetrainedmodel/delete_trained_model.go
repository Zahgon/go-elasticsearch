package deletetrainedmodel

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
)

const (
	modelidMask = iota + 1
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type DeleteTrainedModel struct {
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

type NewDeleteTrainedModel func(modelid string) *DeleteTrainedModel

func NewDeleteTrainedModelFunc(tp elastictransport.Interface) NewDeleteTrainedModel {
	_ = "STUB: not implemented"
	return *new(NewDeleteTrainedModel)
}

func New(tp elastictransport.Interface) *DeleteTrainedModel { _ = "STUB: not implemented"; return nil }

func (r *DeleteTrainedModel) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r DeleteTrainedModel) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r DeleteTrainedModel) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r DeleteTrainedModel) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *DeleteTrainedModel) Header(key, value string) *DeleteTrainedModel {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteTrainedModel) _modelid(modelid string) *DeleteTrainedModel {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteTrainedModel) Force(force bool) *DeleteTrainedModel {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteTrainedModel) Timeout(duration string) *DeleteTrainedModel {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteTrainedModel) ErrorTrace(errortrace bool) *DeleteTrainedModel {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteTrainedModel) FilterPath(filterpaths ...string) *DeleteTrainedModel {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteTrainedModel) Human(human bool) *DeleteTrainedModel {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteTrainedModel) Pretty(pretty bool) *DeleteTrainedModel {
	_ = "STUB: not implemented"
	return nil
}
