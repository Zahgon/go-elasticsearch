package esapi

import (
	"context"
	"io"
	"net/http"
	"time"
)

func newInferencePutDeepseekFunc(t Transport) InferencePutDeepseek {
	_ = "STUB: not implemented"
	return *new(InferencePutDeepseek)
}

type InferencePutDeepseek func(body io.Reader, deepseek_inference_id string, task_type string, o ...func(*InferencePutDeepseekRequest)) (*Response, error)

type InferencePutDeepseekRequest struct {
	Body io.Reader

	DeepseekInferenceID string
	TaskType            string

	Timeout time.Duration

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r InferencePutDeepseekRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f InferencePutDeepseek) WithContext(v context.Context) func(*InferencePutDeepseekRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferencePutDeepseek) WithTimeout(v time.Duration) func(*InferencePutDeepseekRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferencePutDeepseek) WithPretty() func(*InferencePutDeepseekRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferencePutDeepseek) WithHuman() func(*InferencePutDeepseekRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferencePutDeepseek) WithErrorTrace() func(*InferencePutDeepseekRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferencePutDeepseek) WithFilterPath(v ...string) func(*InferencePutDeepseekRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferencePutDeepseek) WithHeader(h map[string]string) func(*InferencePutDeepseekRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferencePutDeepseek) WithOpaqueID(s string) func(*InferencePutDeepseekRequest) {
	_ = "STUB: not implemented"
	return nil
}
