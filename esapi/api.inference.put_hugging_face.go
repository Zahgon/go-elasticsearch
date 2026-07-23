package esapi

import (
	"context"
	"io"
	"net/http"
	"time"
)

func newInferencePutHuggingFaceFunc(t Transport) InferencePutHuggingFace {
	_ = "STUB: not implemented"
	return *new(InferencePutHuggingFace)
}

type InferencePutHuggingFace func(body io.Reader, huggingface_inference_id string, task_type string, o ...func(*InferencePutHuggingFaceRequest)) (*Response, error)

type InferencePutHuggingFaceRequest struct {
	Body io.Reader

	HuggingfaceInferenceID string
	TaskType               string

	Timeout time.Duration

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r InferencePutHuggingFaceRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f InferencePutHuggingFace) WithContext(v context.Context) func(*InferencePutHuggingFaceRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferencePutHuggingFace) WithTimeout(v time.Duration) func(*InferencePutHuggingFaceRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferencePutHuggingFace) WithPretty() func(*InferencePutHuggingFaceRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferencePutHuggingFace) WithHuman() func(*InferencePutHuggingFaceRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferencePutHuggingFace) WithErrorTrace() func(*InferencePutHuggingFaceRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferencePutHuggingFace) WithFilterPath(v ...string) func(*InferencePutHuggingFaceRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferencePutHuggingFace) WithHeader(h map[string]string) func(*InferencePutHuggingFaceRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferencePutHuggingFace) WithOpaqueID(s string) func(*InferencePutHuggingFaceRequest) {
	_ = "STUB: not implemented"
	return nil
}
