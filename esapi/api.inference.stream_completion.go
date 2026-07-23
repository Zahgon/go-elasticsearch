package esapi

import (
	"context"
	"io"
	"net/http"
	"time"
)

func newInferenceStreamCompletionFunc(t Transport) InferenceStreamCompletion {
	_ = "STUB: not implemented"
	return *new(InferenceStreamCompletion)
}

type InferenceStreamCompletion func(body io.Reader, inference_id string, o ...func(*InferenceStreamCompletionRequest)) (*Response, error)

type InferenceStreamCompletionRequest struct {
	Body io.Reader

	InferenceID string

	Timeout time.Duration

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r InferenceStreamCompletionRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f InferenceStreamCompletion) WithContext(v context.Context) func(*InferenceStreamCompletionRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferenceStreamCompletion) WithTimeout(v time.Duration) func(*InferenceStreamCompletionRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferenceStreamCompletion) WithPretty() func(*InferenceStreamCompletionRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferenceStreamCompletion) WithHuman() func(*InferenceStreamCompletionRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferenceStreamCompletion) WithErrorTrace() func(*InferenceStreamCompletionRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferenceStreamCompletion) WithFilterPath(v ...string) func(*InferenceStreamCompletionRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferenceStreamCompletion) WithHeader(h map[string]string) func(*InferenceStreamCompletionRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferenceStreamCompletion) WithOpaqueID(s string) func(*InferenceStreamCompletionRequest) {
	_ = "STUB: not implemented"
	return nil
}
