package esapi

import (
	"context"
	"io"
	"net/http"
)

func newMLPutTrainedModelDefinitionPartFunc(t Transport) MLPutTrainedModelDefinitionPart {
	_ = "STUB: not implemented"
	return *new(MLPutTrainedModelDefinitionPart)
}

type MLPutTrainedModelDefinitionPart func(body io.Reader, model_id string, part *int, o ...func(*MLPutTrainedModelDefinitionPartRequest)) (*Response, error)

type MLPutTrainedModelDefinitionPartRequest struct {
	Body io.Reader

	ModelID string
	Part    *int

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r MLPutTrainedModelDefinitionPartRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f MLPutTrainedModelDefinitionPart) WithContext(v context.Context) func(*MLPutTrainedModelDefinitionPartRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLPutTrainedModelDefinitionPart) WithPretty() func(*MLPutTrainedModelDefinitionPartRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLPutTrainedModelDefinitionPart) WithHuman() func(*MLPutTrainedModelDefinitionPartRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLPutTrainedModelDefinitionPart) WithErrorTrace() func(*MLPutTrainedModelDefinitionPartRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLPutTrainedModelDefinitionPart) WithFilterPath(v ...string) func(*MLPutTrainedModelDefinitionPartRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLPutTrainedModelDefinitionPart) WithHeader(h map[string]string) func(*MLPutTrainedModelDefinitionPartRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLPutTrainedModelDefinitionPart) WithOpaqueID(s string) func(*MLPutTrainedModelDefinitionPartRequest) {
	_ = "STUB: not implemented"
	return nil
}
