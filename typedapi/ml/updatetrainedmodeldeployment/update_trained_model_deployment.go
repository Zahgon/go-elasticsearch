package updatetrainedmodeldeployment

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
	modelidMask = iota + 1
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type UpdateTrainedModelDeployment struct {
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

type NewUpdateTrainedModelDeployment func(modelid string) *UpdateTrainedModelDeployment

func NewUpdateTrainedModelDeploymentFunc(tp elastictransport.Interface) NewUpdateTrainedModelDeployment {
	_ = "STUB: not implemented"
	return *new(NewUpdateTrainedModelDeployment)
}

func New(tp elastictransport.Interface) *UpdateTrainedModelDeployment {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateTrainedModelDeployment) Raw(raw io.Reader) *UpdateTrainedModelDeployment {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateTrainedModelDeployment) Request(req *Request) *UpdateTrainedModelDeployment {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateTrainedModelDeployment) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r UpdateTrainedModelDeployment) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r UpdateTrainedModelDeployment) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *UpdateTrainedModelDeployment) Header(key, value string) *UpdateTrainedModelDeployment {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateTrainedModelDeployment) _modelid(modelid string) *UpdateTrainedModelDeployment {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateTrainedModelDeployment) ErrorTrace(errortrace bool) *UpdateTrainedModelDeployment {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateTrainedModelDeployment) FilterPath(filterpaths ...string) *UpdateTrainedModelDeployment {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateTrainedModelDeployment) Human(human bool) *UpdateTrainedModelDeployment {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateTrainedModelDeployment) Pretty(pretty bool) *UpdateTrainedModelDeployment {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateTrainedModelDeployment) AdaptiveAllocations(adaptiveallocations types.AdaptiveAllocationsSettingsVariant) *UpdateTrainedModelDeployment {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateTrainedModelDeployment) NumberOfAllocations(numberofallocations int) *UpdateTrainedModelDeployment {
	_ = "STUB: not implemented"
	return nil
}
