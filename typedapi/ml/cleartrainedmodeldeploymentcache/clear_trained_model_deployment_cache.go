package cleartrainedmodeldeploymentcache

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

type ClearTrainedModelDeploymentCache struct {
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

type NewClearTrainedModelDeploymentCache func(modelid string) *ClearTrainedModelDeploymentCache

func NewClearTrainedModelDeploymentCacheFunc(tp elastictransport.Interface) NewClearTrainedModelDeploymentCache {
	_ = "STUB: not implemented"
	return *new(NewClearTrainedModelDeploymentCache)
}

func New(tp elastictransport.Interface) *ClearTrainedModelDeploymentCache {
	_ = "STUB: not implemented"
	return nil
}

func (r *ClearTrainedModelDeploymentCache) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r ClearTrainedModelDeploymentCache) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r ClearTrainedModelDeploymentCache) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r ClearTrainedModelDeploymentCache) IsSuccess(providedCtx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *ClearTrainedModelDeploymentCache) Header(key, value string) *ClearTrainedModelDeploymentCache {
	_ = "STUB: not implemented"
	return nil
}

func (r *ClearTrainedModelDeploymentCache) _modelid(modelid string) *ClearTrainedModelDeploymentCache {
	_ = "STUB: not implemented"
	return nil
}

func (r *ClearTrainedModelDeploymentCache) ErrorTrace(errortrace bool) *ClearTrainedModelDeploymentCache {
	_ = "STUB: not implemented"
	return nil
}

func (r *ClearTrainedModelDeploymentCache) FilterPath(filterpaths ...string) *ClearTrainedModelDeploymentCache {
	_ = "STUB: not implemented"
	return nil
}

func (r *ClearTrainedModelDeploymentCache) Human(human bool) *ClearTrainedModelDeploymentCache {
	_ = "STUB: not implemented"
	return nil
}

func (r *ClearTrainedModelDeploymentCache) Pretty(pretty bool) *ClearTrainedModelDeploymentCache {
	_ = "STUB: not implemented"
	return nil
}
