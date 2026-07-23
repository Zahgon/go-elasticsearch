package esapi

import (
	"context"
	"io"
	"net/http"
)

func newMLUpdateTrainedModelDeploymentFunc(t Transport) MLUpdateTrainedModelDeployment {
	_ = "STUB: not implemented"
	return *new(MLUpdateTrainedModelDeployment)
}

type MLUpdateTrainedModelDeployment func(model_id string, o ...func(*MLUpdateTrainedModelDeploymentRequest)) (*Response, error)

type MLUpdateTrainedModelDeploymentRequest struct {
	Body io.Reader

	ModelID string

	NumberOfAllocations *int

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r MLUpdateTrainedModelDeploymentRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f MLUpdateTrainedModelDeployment) WithContext(v context.Context) func(*MLUpdateTrainedModelDeploymentRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLUpdateTrainedModelDeployment) WithBody(v io.Reader) func(*MLUpdateTrainedModelDeploymentRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLUpdateTrainedModelDeployment) WithNumberOfAllocations(v int) func(*MLUpdateTrainedModelDeploymentRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLUpdateTrainedModelDeployment) WithPretty() func(*MLUpdateTrainedModelDeploymentRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLUpdateTrainedModelDeployment) WithHuman() func(*MLUpdateTrainedModelDeploymentRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLUpdateTrainedModelDeployment) WithErrorTrace() func(*MLUpdateTrainedModelDeploymentRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLUpdateTrainedModelDeployment) WithFilterPath(v ...string) func(*MLUpdateTrainedModelDeploymentRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLUpdateTrainedModelDeployment) WithHeader(h map[string]string) func(*MLUpdateTrainedModelDeploymentRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLUpdateTrainedModelDeployment) WithOpaqueID(s string) func(*MLUpdateTrainedModelDeploymentRequest) {
	_ = "STUB: not implemented"
	return nil
}
