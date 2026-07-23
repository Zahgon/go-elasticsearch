package esapi

import (
	"context"
	"io"
	"net/http"
	"time"
)

func newInferenceChatCompletionUnifiedFunc(t Transport) InferenceChatCompletionUnified {
	_ = "STUB: not implemented"
	return *new(InferenceChatCompletionUnified)
}

type InferenceChatCompletionUnified func(body io.Reader, inference_id string, o ...func(*InferenceChatCompletionUnifiedRequest)) (*Response, error)

type InferenceChatCompletionUnifiedRequest struct {
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

func (r InferenceChatCompletionUnifiedRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f InferenceChatCompletionUnified) WithContext(v context.Context) func(*InferenceChatCompletionUnifiedRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferenceChatCompletionUnified) WithTimeout(v time.Duration) func(*InferenceChatCompletionUnifiedRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferenceChatCompletionUnified) WithPretty() func(*InferenceChatCompletionUnifiedRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferenceChatCompletionUnified) WithHuman() func(*InferenceChatCompletionUnifiedRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferenceChatCompletionUnified) WithErrorTrace() func(*InferenceChatCompletionUnifiedRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferenceChatCompletionUnified) WithFilterPath(v ...string) func(*InferenceChatCompletionUnifiedRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferenceChatCompletionUnified) WithHeader(h map[string]string) func(*InferenceChatCompletionUnifiedRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferenceChatCompletionUnified) WithOpaqueID(s string) func(*InferenceChatCompletionUnifiedRequest) {
	_ = "STUB: not implemented"
	return nil
}
