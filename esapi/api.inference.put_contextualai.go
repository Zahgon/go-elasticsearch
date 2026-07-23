package esapi

import (
	"context"
	"io"
	"net/http"
	"time"
)

func newInferencePutContextualaiFunc(t Transport) InferencePutContextualai {
	_ = "STUB: not implemented"
	return *new(InferencePutContextualai)
}

type InferencePutContextualai func(body io.Reader, contextualai_inference_id string, task_type string, o ...func(*InferencePutContextualaiRequest)) (*Response, error)

type InferencePutContextualaiRequest struct {
	Body io.Reader

	ContextualaiInferenceID string
	TaskType                string

	Timeout time.Duration

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r InferencePutContextualaiRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f InferencePutContextualai) WithContext(v context.Context) func(*InferencePutContextualaiRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferencePutContextualai) WithTimeout(v time.Duration) func(*InferencePutContextualaiRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferencePutContextualai) WithPretty() func(*InferencePutContextualaiRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferencePutContextualai) WithHuman() func(*InferencePutContextualaiRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferencePutContextualai) WithErrorTrace() func(*InferencePutContextualaiRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferencePutContextualai) WithFilterPath(v ...string) func(*InferencePutContextualaiRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferencePutContextualai) WithHeader(h map[string]string) func(*InferencePutContextualaiRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferencePutContextualai) WithOpaqueID(s string) func(*InferencePutContextualaiRequest) {
	_ = "STUB: not implemented"
	return nil
}
