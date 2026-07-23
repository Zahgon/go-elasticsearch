package esapi

import (
	"context"
	"io"
	"net/http"
	"time"
)

func newInferenceCompletionFunc(t Transport) InferenceCompletion {
	_ = "STUB: not implemented"
	return *new(InferenceCompletion)
}

type InferenceCompletion func(body io.Reader, inference_id string, o ...func(*InferenceCompletionRequest)) (*Response, error)

type InferenceCompletionRequest struct {
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

func (r InferenceCompletionRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f InferenceCompletion) WithContext(v context.Context) func(*InferenceCompletionRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferenceCompletion) WithTimeout(v time.Duration) func(*InferenceCompletionRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferenceCompletion) WithPretty() func(*InferenceCompletionRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferenceCompletion) WithHuman() func(*InferenceCompletionRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferenceCompletion) WithErrorTrace() func(*InferenceCompletionRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferenceCompletion) WithFilterPath(v ...string) func(*InferenceCompletionRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferenceCompletion) WithHeader(h map[string]string) func(*InferenceCompletionRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferenceCompletion) WithOpaqueID(s string) func(*InferenceCompletionRequest) {
	_ = "STUB: not implemented"
	return nil
}
