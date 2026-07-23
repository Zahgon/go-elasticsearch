package esapi

import (
	"context"
	"io"
	"net/http"
	"time"
)

func newInferencePutLlamaFunc(t Transport) InferencePutLlama {
	_ = "STUB: not implemented"
	return *new(InferencePutLlama)
}

type InferencePutLlama func(body io.Reader, llama_inference_id string, task_type string, o ...func(*InferencePutLlamaRequest)) (*Response, error)

type InferencePutLlamaRequest struct {
	Body io.Reader

	LlamaInferenceID string
	TaskType         string

	Timeout time.Duration

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r InferencePutLlamaRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f InferencePutLlama) WithContext(v context.Context) func(*InferencePutLlamaRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferencePutLlama) WithTimeout(v time.Duration) func(*InferencePutLlamaRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferencePutLlama) WithPretty() func(*InferencePutLlamaRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferencePutLlama) WithHuman() func(*InferencePutLlamaRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferencePutLlama) WithErrorTrace() func(*InferencePutLlamaRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferencePutLlama) WithFilterPath(v ...string) func(*InferencePutLlamaRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferencePutLlama) WithHeader(h map[string]string) func(*InferencePutLlamaRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferencePutLlama) WithOpaqueID(s string) func(*InferencePutLlamaRequest) {
	_ = "STUB: not implemented"
	return nil
}
