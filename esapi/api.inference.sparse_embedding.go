package esapi

import (
	"context"
	"io"
	"net/http"
	"time"
)

func newInferenceSparseEmbeddingFunc(t Transport) InferenceSparseEmbedding {
	_ = "STUB: not implemented"
	return *new(InferenceSparseEmbedding)
}

type InferenceSparseEmbedding func(body io.Reader, inference_id string, o ...func(*InferenceSparseEmbeddingRequest)) (*Response, error)

type InferenceSparseEmbeddingRequest struct {
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

func (r InferenceSparseEmbeddingRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f InferenceSparseEmbedding) WithContext(v context.Context) func(*InferenceSparseEmbeddingRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferenceSparseEmbedding) WithTimeout(v time.Duration) func(*InferenceSparseEmbeddingRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferenceSparseEmbedding) WithPretty() func(*InferenceSparseEmbeddingRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferenceSparseEmbedding) WithHuman() func(*InferenceSparseEmbeddingRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferenceSparseEmbedding) WithErrorTrace() func(*InferenceSparseEmbeddingRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferenceSparseEmbedding) WithFilterPath(v ...string) func(*InferenceSparseEmbeddingRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferenceSparseEmbedding) WithHeader(h map[string]string) func(*InferenceSparseEmbeddingRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferenceSparseEmbedding) WithOpaqueID(s string) func(*InferenceSparseEmbeddingRequest) {
	_ = "STUB: not implemented"
	return nil
}
