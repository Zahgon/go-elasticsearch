package esapi

import (
	"context"
	"io"
	"net/http"
	"time"
)

func newInferenceInferenceFunc(t Transport) InferenceInference {
	_ = "STUB: not implemented"
	return *new(InferenceInference)
}

type InferenceInference func(body io.Reader, inference_id string, o ...func(*InferenceInferenceRequest)) (*Response, error)

type InferenceInferenceRequest struct {
	Body io.Reader

	InferenceID string
	TaskType    string

	Timeout time.Duration

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r InferenceInferenceRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f InferenceInference) WithContext(v context.Context) func(*InferenceInferenceRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferenceInference) WithTaskType(v string) func(*InferenceInferenceRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferenceInference) WithTimeout(v time.Duration) func(*InferenceInferenceRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferenceInference) WithPretty() func(*InferenceInferenceRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferenceInference) WithHuman() func(*InferenceInferenceRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferenceInference) WithErrorTrace() func(*InferenceInferenceRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferenceInference) WithFilterPath(v ...string) func(*InferenceInferenceRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferenceInference) WithHeader(h map[string]string) func(*InferenceInferenceRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferenceInference) WithOpaqueID(s string) func(*InferenceInferenceRequest) {
	_ = "STUB: not implemented"
	return nil
}
