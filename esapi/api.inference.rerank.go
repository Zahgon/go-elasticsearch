package esapi

import (
	"context"
	"io"
	"net/http"
	"time"
)

func newInferenceRerankFunc(t Transport) InferenceRerank {
	_ = "STUB: not implemented"
	return *new(InferenceRerank)
}

type InferenceRerank func(body io.Reader, inference_id string, o ...func(*InferenceRerankRequest)) (*Response, error)

type InferenceRerankRequest struct {
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

func (r InferenceRerankRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f InferenceRerank) WithContext(v context.Context) func(*InferenceRerankRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferenceRerank) WithTimeout(v time.Duration) func(*InferenceRerankRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferenceRerank) WithPretty() func(*InferenceRerankRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferenceRerank) WithHuman() func(*InferenceRerankRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferenceRerank) WithErrorTrace() func(*InferenceRerankRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferenceRerank) WithFilterPath(v ...string) func(*InferenceRerankRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferenceRerank) WithHeader(h map[string]string) func(*InferenceRerankRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferenceRerank) WithOpaqueID(s string) func(*InferenceRerankRequest) {
	_ = "STUB: not implemented"
	return nil
}
