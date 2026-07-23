package esapi

import (
	"context"
	"net/http"
)

func newMLClearTrainedModelDeploymentCacheFunc(t Transport) MLClearTrainedModelDeploymentCache {
	_ = "STUB: not implemented"
	return *new(MLClearTrainedModelDeploymentCache)
}

type MLClearTrainedModelDeploymentCache func(model_id string, o ...func(*MLClearTrainedModelDeploymentCacheRequest)) (*Response, error)

type MLClearTrainedModelDeploymentCacheRequest struct {
	ModelID string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r MLClearTrainedModelDeploymentCacheRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f MLClearTrainedModelDeploymentCache) WithContext(v context.Context) func(*MLClearTrainedModelDeploymentCacheRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLClearTrainedModelDeploymentCache) WithPretty() func(*MLClearTrainedModelDeploymentCacheRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLClearTrainedModelDeploymentCache) WithHuman() func(*MLClearTrainedModelDeploymentCacheRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLClearTrainedModelDeploymentCache) WithErrorTrace() func(*MLClearTrainedModelDeploymentCacheRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLClearTrainedModelDeploymentCache) WithFilterPath(v ...string) func(*MLClearTrainedModelDeploymentCacheRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLClearTrainedModelDeploymentCache) WithHeader(h map[string]string) func(*MLClearTrainedModelDeploymentCacheRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLClearTrainedModelDeploymentCache) WithOpaqueID(s string) func(*MLClearTrainedModelDeploymentCacheRequest) {
	_ = "STUB: not implemented"
	return nil
}
