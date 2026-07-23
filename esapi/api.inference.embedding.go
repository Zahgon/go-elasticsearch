package esapi

import (
	"context"
	"io"
	"net/http"
	"time"
)

func newInferenceEmbeddingFunc(t Transport) InferenceEmbedding {
	_ = "STUB: not implemented"
	return *new(InferenceEmbedding)
}

type InferenceEmbedding func(body io.Reader, inference_id string, o ...func(*InferenceEmbeddingRequest)) (*Response, error)

type InferenceEmbeddingRequest struct {
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

func (r InferenceEmbeddingRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f InferenceEmbedding) WithContext(v context.Context) func(*InferenceEmbeddingRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferenceEmbedding) WithTimeout(v time.Duration) func(*InferenceEmbeddingRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferenceEmbedding) WithPretty() func(*InferenceEmbeddingRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferenceEmbedding) WithHuman() func(*InferenceEmbeddingRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferenceEmbedding) WithErrorTrace() func(*InferenceEmbeddingRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferenceEmbedding) WithFilterPath(v ...string) func(*InferenceEmbeddingRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferenceEmbedding) WithHeader(h map[string]string) func(*InferenceEmbeddingRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferenceEmbedding) WithOpaqueID(s string) func(*InferenceEmbeddingRequest) {
	_ = "STUB: not implemented"
	return nil
}
