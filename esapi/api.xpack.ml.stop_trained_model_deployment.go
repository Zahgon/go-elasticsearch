package esapi

import (
	"context"
	"io"
	"net/http"
)

func newMLStopTrainedModelDeploymentFunc(t Transport) MLStopTrainedModelDeployment {
	_ = "STUB: not implemented"
	return *new(MLStopTrainedModelDeployment)
}

type MLStopTrainedModelDeployment func(model_id string, o ...func(*MLStopTrainedModelDeploymentRequest)) (*Response, error)

type MLStopTrainedModelDeploymentRequest struct {
	Body io.Reader

	ModelID string

	AllowNoMatch *bool
	Force        *bool

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r MLStopTrainedModelDeploymentRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f MLStopTrainedModelDeployment) WithContext(v context.Context) func(*MLStopTrainedModelDeploymentRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLStopTrainedModelDeployment) WithBody(v io.Reader) func(*MLStopTrainedModelDeploymentRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLStopTrainedModelDeployment) WithAllowNoMatch(v bool) func(*MLStopTrainedModelDeploymentRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLStopTrainedModelDeployment) WithForce(v bool) func(*MLStopTrainedModelDeploymentRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLStopTrainedModelDeployment) WithPretty() func(*MLStopTrainedModelDeploymentRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLStopTrainedModelDeployment) WithHuman() func(*MLStopTrainedModelDeploymentRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLStopTrainedModelDeployment) WithErrorTrace() func(*MLStopTrainedModelDeploymentRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLStopTrainedModelDeployment) WithFilterPath(v ...string) func(*MLStopTrainedModelDeploymentRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLStopTrainedModelDeployment) WithHeader(h map[string]string) func(*MLStopTrainedModelDeploymentRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLStopTrainedModelDeployment) WithOpaqueID(s string) func(*MLStopTrainedModelDeploymentRequest) {
	_ = "STUB: not implemented"
	return nil
}
