package esapi

import (
	"context"
	"io"
	"net/http"
	"time"
)

func newMLInferTrainedModelFunc(t Transport) MLInferTrainedModel {
	_ = "STUB: not implemented"
	return *new(MLInferTrainedModel)
}

type MLInferTrainedModel func(body io.Reader, model_id string, o ...func(*MLInferTrainedModelRequest)) (*Response, error)

type MLInferTrainedModelRequest struct {
	Body io.Reader

	ModelID string

	Timeout time.Duration

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r MLInferTrainedModelRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f MLInferTrainedModel) WithContext(v context.Context) func(*MLInferTrainedModelRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLInferTrainedModel) WithTimeout(v time.Duration) func(*MLInferTrainedModelRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLInferTrainedModel) WithPretty() func(*MLInferTrainedModelRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLInferTrainedModel) WithHuman() func(*MLInferTrainedModelRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLInferTrainedModel) WithErrorTrace() func(*MLInferTrainedModelRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLInferTrainedModel) WithFilterPath(v ...string) func(*MLInferTrainedModelRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLInferTrainedModel) WithHeader(h map[string]string) func(*MLInferTrainedModelRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLInferTrainedModel) WithOpaqueID(s string) func(*MLInferTrainedModelRequest) {
	_ = "STUB: not implemented"
	return nil
}
