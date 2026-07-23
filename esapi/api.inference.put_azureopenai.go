package esapi

import (
	"context"
	"io"
	"net/http"
	"time"
)

func newInferencePutAzureopenaiFunc(t Transport) InferencePutAzureopenai {
	_ = "STUB: not implemented"
	return *new(InferencePutAzureopenai)
}

type InferencePutAzureopenai func(body io.Reader, azureopenai_inference_id string, task_type string, o ...func(*InferencePutAzureopenaiRequest)) (*Response, error)

type InferencePutAzureopenaiRequest struct {
	Body io.Reader

	AzureopenaiInferenceID string
	TaskType               string

	Timeout time.Duration

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r InferencePutAzureopenaiRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f InferencePutAzureopenai) WithContext(v context.Context) func(*InferencePutAzureopenaiRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferencePutAzureopenai) WithTimeout(v time.Duration) func(*InferencePutAzureopenaiRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferencePutAzureopenai) WithPretty() func(*InferencePutAzureopenaiRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferencePutAzureopenai) WithHuman() func(*InferencePutAzureopenaiRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferencePutAzureopenai) WithErrorTrace() func(*InferencePutAzureopenaiRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferencePutAzureopenai) WithFilterPath(v ...string) func(*InferencePutAzureopenaiRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferencePutAzureopenai) WithHeader(h map[string]string) func(*InferencePutAzureopenaiRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferencePutAzureopenai) WithOpaqueID(s string) func(*InferencePutAzureopenaiRequest) {
	_ = "STUB: not implemented"
	return nil
}
