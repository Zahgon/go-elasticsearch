package esapi

import (
	"context"
	"io"
	"net/http"
	"time"
)

func newInferencePutAnthropicFunc(t Transport) InferencePutAnthropic {
	_ = "STUB: not implemented"
	return *new(InferencePutAnthropic)
}

type InferencePutAnthropic func(body io.Reader, anthropic_inference_id string, task_type string, o ...func(*InferencePutAnthropicRequest)) (*Response, error)

type InferencePutAnthropicRequest struct {
	Body io.Reader

	AnthropicInferenceID string
	TaskType             string

	Timeout time.Duration

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r InferencePutAnthropicRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f InferencePutAnthropic) WithContext(v context.Context) func(*InferencePutAnthropicRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferencePutAnthropic) WithTimeout(v time.Duration) func(*InferencePutAnthropicRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferencePutAnthropic) WithPretty() func(*InferencePutAnthropicRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferencePutAnthropic) WithHuman() func(*InferencePutAnthropicRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferencePutAnthropic) WithErrorTrace() func(*InferencePutAnthropicRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferencePutAnthropic) WithFilterPath(v ...string) func(*InferencePutAnthropicRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferencePutAnthropic) WithHeader(h map[string]string) func(*InferencePutAnthropicRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferencePutAnthropic) WithOpaqueID(s string) func(*InferencePutAnthropicRequest) {
	_ = "STUB: not implemented"
	return nil
}
