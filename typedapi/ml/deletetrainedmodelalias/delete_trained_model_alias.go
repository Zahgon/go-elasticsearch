package deletetrainedmodelalias

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
)

const (
	modelaliasMask = iota + 1

	modelidMask
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type DeleteTrainedModelAlias struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	paramSet int

	modelalias string
	modelid    string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewDeleteTrainedModelAlias func(modelid, modelalias string) *DeleteTrainedModelAlias

func NewDeleteTrainedModelAliasFunc(tp elastictransport.Interface) NewDeleteTrainedModelAlias {
	_ = "STUB: not implemented"
	return *new(NewDeleteTrainedModelAlias)
}

func New(tp elastictransport.Interface) *DeleteTrainedModelAlias {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteTrainedModelAlias) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r DeleteTrainedModelAlias) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r DeleteTrainedModelAlias) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r DeleteTrainedModelAlias) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *DeleteTrainedModelAlias) Header(key, value string) *DeleteTrainedModelAlias {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteTrainedModelAlias) _modelalias(modelalias string) *DeleteTrainedModelAlias {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteTrainedModelAlias) _modelid(modelid string) *DeleteTrainedModelAlias {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteTrainedModelAlias) ErrorTrace(errortrace bool) *DeleteTrainedModelAlias {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteTrainedModelAlias) FilterPath(filterpaths ...string) *DeleteTrainedModelAlias {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteTrainedModelAlias) Human(human bool) *DeleteTrainedModelAlias {
	_ = "STUB: not implemented"
	return nil
}

func (r *DeleteTrainedModelAlias) Pretty(pretty bool) *DeleteTrainedModelAlias {
	_ = "STUB: not implemented"
	return nil
}
