package esapi

import (
	"context"
	"io"
	"net/http"
	"time"
)

func newInferencePutMistralFunc(t Transport) InferencePutMistral {
	_ = "STUB: not implemented"
	return *new(InferencePutMistral)
}

type InferencePutMistral func(body io.Reader, mistral_inference_id string, task_type string, o ...func(*InferencePutMistralRequest)) (*Response, error)

type InferencePutMistralRequest struct {
	Body io.Reader

	MistralInferenceID string
	TaskType           string

	Timeout time.Duration

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r InferencePutMistralRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f InferencePutMistral) WithContext(v context.Context) func(*InferencePutMistralRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferencePutMistral) WithTimeout(v time.Duration) func(*InferencePutMistralRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferencePutMistral) WithPretty() func(*InferencePutMistralRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferencePutMistral) WithHuman() func(*InferencePutMistralRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferencePutMistral) WithErrorTrace() func(*InferencePutMistralRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferencePutMistral) WithFilterPath(v ...string) func(*InferencePutMistralRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferencePutMistral) WithHeader(h map[string]string) func(*InferencePutMistralRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferencePutMistral) WithOpaqueID(s string) func(*InferencePutMistralRequest) {
	_ = "STUB: not implemented"
	return nil
}
