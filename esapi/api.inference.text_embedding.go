package esapi

import (
	"context"
	"io"
	"net/http"
	"time"
)

func newInferenceTextEmbeddingFunc(t Transport) InferenceTextEmbedding {
	_ = "STUB: not implemented"
	return *new(InferenceTextEmbedding)
}

type InferenceTextEmbedding func(body io.Reader, inference_id string, o ...func(*InferenceTextEmbeddingRequest)) (*Response, error)

type InferenceTextEmbeddingRequest struct {
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

func (r InferenceTextEmbeddingRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f InferenceTextEmbedding) WithContext(v context.Context) func(*InferenceTextEmbeddingRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferenceTextEmbedding) WithTimeout(v time.Duration) func(*InferenceTextEmbeddingRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferenceTextEmbedding) WithPretty() func(*InferenceTextEmbeddingRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferenceTextEmbedding) WithHuman() func(*InferenceTextEmbeddingRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferenceTextEmbedding) WithErrorTrace() func(*InferenceTextEmbeddingRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferenceTextEmbedding) WithFilterPath(v ...string) func(*InferenceTextEmbeddingRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferenceTextEmbedding) WithHeader(h map[string]string) func(*InferenceTextEmbeddingRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferenceTextEmbedding) WithOpaqueID(s string) func(*InferenceTextEmbeddingRequest) {
	_ = "STUB: not implemented"
	return nil
}
