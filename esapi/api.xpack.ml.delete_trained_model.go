package esapi

import (
	"context"
	"net/http"
	"time"
)

func newMLDeleteTrainedModelFunc(t Transport) MLDeleteTrainedModel {
	_ = "STUB: not implemented"
	return *new(MLDeleteTrainedModel)
}

type MLDeleteTrainedModel func(model_id string, o ...func(*MLDeleteTrainedModelRequest)) (*Response, error)

type MLDeleteTrainedModelRequest struct {
	ModelID string

	Force   *bool
	Timeout time.Duration

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r MLDeleteTrainedModelRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f MLDeleteTrainedModel) WithContext(v context.Context) func(*MLDeleteTrainedModelRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLDeleteTrainedModel) WithForce(v bool) func(*MLDeleteTrainedModelRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLDeleteTrainedModel) WithTimeout(v time.Duration) func(*MLDeleteTrainedModelRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLDeleteTrainedModel) WithPretty() func(*MLDeleteTrainedModelRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLDeleteTrainedModel) WithHuman() func(*MLDeleteTrainedModelRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLDeleteTrainedModel) WithErrorTrace() func(*MLDeleteTrainedModelRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLDeleteTrainedModel) WithFilterPath(v ...string) func(*MLDeleteTrainedModelRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLDeleteTrainedModel) WithHeader(h map[string]string) func(*MLDeleteTrainedModelRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLDeleteTrainedModel) WithOpaqueID(s string) func(*MLDeleteTrainedModelRequest) {
	_ = "STUB: not implemented"
	return nil
}
