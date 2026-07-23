package esapi

import (
	"context"
	"io"
	"net/http"
	"time"
)

func newInferencePutGroqFunc(t Transport) InferencePutGroq {
	_ = "STUB: not implemented"
	return *new(InferencePutGroq)
}

type InferencePutGroq func(body io.Reader, groq_inference_id string, task_type string, o ...func(*InferencePutGroqRequest)) (*Response, error)

type InferencePutGroqRequest struct {
	Body io.Reader

	GroqInferenceID string
	TaskType        string

	Timeout time.Duration

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r InferencePutGroqRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f InferencePutGroq) WithContext(v context.Context) func(*InferencePutGroqRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferencePutGroq) WithTimeout(v time.Duration) func(*InferencePutGroqRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferencePutGroq) WithPretty() func(*InferencePutGroqRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferencePutGroq) WithHuman() func(*InferencePutGroqRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferencePutGroq) WithErrorTrace() func(*InferencePutGroqRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferencePutGroq) WithFilterPath(v ...string) func(*InferencePutGroqRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferencePutGroq) WithHeader(h map[string]string) func(*InferencePutGroqRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferencePutGroq) WithOpaqueID(s string) func(*InferencePutGroqRequest) {
	_ = "STUB: not implemented"
	return nil
}
