package puttrainedmodelalias

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

type PutTrainedModelAlias struct {
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

type NewPutTrainedModelAlias func(modelid, modelalias string) *PutTrainedModelAlias

func NewPutTrainedModelAliasFunc(tp elastictransport.Interface) NewPutTrainedModelAlias {
	_ = "STUB: not implemented"
	return *new(NewPutTrainedModelAlias)
}

func New(tp elastictransport.Interface) *PutTrainedModelAlias {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutTrainedModelAlias) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r PutTrainedModelAlias) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r PutTrainedModelAlias) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r PutTrainedModelAlias) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *PutTrainedModelAlias) Header(key, value string) *PutTrainedModelAlias {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutTrainedModelAlias) _modelalias(modelalias string) *PutTrainedModelAlias {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutTrainedModelAlias) _modelid(modelid string) *PutTrainedModelAlias {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutTrainedModelAlias) Reassign(reassign bool) *PutTrainedModelAlias {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutTrainedModelAlias) ErrorTrace(errortrace bool) *PutTrainedModelAlias {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutTrainedModelAlias) FilterPath(filterpaths ...string) *PutTrainedModelAlias {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutTrainedModelAlias) Human(human bool) *PutTrainedModelAlias {
	_ = "STUB: not implemented"
	return nil
}

func (r *PutTrainedModelAlias) Pretty(pretty bool) *PutTrainedModelAlias {
	_ = "STUB: not implemented"
	return nil
}
