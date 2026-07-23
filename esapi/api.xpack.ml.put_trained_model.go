package esapi

import (
	"context"
	"io"
	"net/http"
)

func newMLPutTrainedModelFunc(t Transport) MLPutTrainedModel {
	_ = "STUB: not implemented"
	return *new(MLPutTrainedModel)
}

type MLPutTrainedModel func(body io.Reader, model_id string, o ...func(*MLPutTrainedModelRequest)) (*Response, error)

type MLPutTrainedModelRequest struct {
	Body io.Reader

	ModelID string

	DeferDefinitionDecompression *bool
	WaitForCompletion            *bool

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r MLPutTrainedModelRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f MLPutTrainedModel) WithContext(v context.Context) func(*MLPutTrainedModelRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLPutTrainedModel) WithDeferDefinitionDecompression(v bool) func(*MLPutTrainedModelRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLPutTrainedModel) WithWaitForCompletion(v bool) func(*MLPutTrainedModelRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLPutTrainedModel) WithPretty() func(*MLPutTrainedModelRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLPutTrainedModel) WithHuman() func(*MLPutTrainedModelRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLPutTrainedModel) WithErrorTrace() func(*MLPutTrainedModelRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLPutTrainedModel) WithFilterPath(v ...string) func(*MLPutTrainedModelRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLPutTrainedModel) WithHeader(h map[string]string) func(*MLPutTrainedModelRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLPutTrainedModel) WithOpaqueID(s string) func(*MLPutTrainedModelRequest) {
	_ = "STUB: not implemented"
	return nil
}
