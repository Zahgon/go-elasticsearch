package stoptrainedmodeldeployment

import (
	gobytes "bytes"
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

type StopTrainedModelDeployment struct {
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

type NewStopTrainedModelDeployment func(modelid string) *StopTrainedModelDeployment

func NewStopTrainedModelDeploymentFunc(tp elastictransport.Interface) NewStopTrainedModelDeployment {
	_ = "STUB: not implemented"
	return *new(NewStopTrainedModelDeployment)
}

func New(tp elastictransport.Interface) *StopTrainedModelDeployment {
	_ = "STUB: not implemented"
	return nil
}

func (r *StopTrainedModelDeployment) Raw(raw io.Reader) *StopTrainedModelDeployment {
	_ = "STUB: not implemented"
	return nil
}

func (r *StopTrainedModelDeployment) Request(req *Request) *StopTrainedModelDeployment {
	_ = "STUB: not implemented"
	return nil
}

func (r *StopTrainedModelDeployment) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r StopTrainedModelDeployment) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r StopTrainedModelDeployment) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *StopTrainedModelDeployment) Header(key, value string) *StopTrainedModelDeployment {
	_ = "STUB: not implemented"
	return nil
}

func (r *StopTrainedModelDeployment) _modelid(modelid string) *StopTrainedModelDeployment {
	_ = "STUB: not implemented"
	return nil
}

func (r *StopTrainedModelDeployment) ErrorTrace(errortrace bool) *StopTrainedModelDeployment {
	_ = "STUB: not implemented"
	return nil
}

func (r *StopTrainedModelDeployment) FilterPath(filterpaths ...string) *StopTrainedModelDeployment {
	_ = "STUB: not implemented"
	return nil
}

func (r *StopTrainedModelDeployment) Human(human bool) *StopTrainedModelDeployment {
	_ = "STUB: not implemented"
	return nil
}

func (r *StopTrainedModelDeployment) Pretty(pretty bool) *StopTrainedModelDeployment {
	_ = "STUB: not implemented"
	return nil
}

func (r *StopTrainedModelDeployment) AllowNoMatch(allownomatch bool) *StopTrainedModelDeployment {
	_ = "STUB: not implemented"
	return nil
}

func (r *StopTrainedModelDeployment) Force(force bool) *StopTrainedModelDeployment {
	_ = "STUB: not implemented"
	return nil
}

func (r *StopTrainedModelDeployment) Id(id string) *StopTrainedModelDeployment {
	_ = "STUB: not implemented"
	return nil
}
