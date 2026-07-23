package esapi

import (
	"context"
	"io"
	"net/http"
	"time"
)

func newInferencePutAzureaistudioFunc(t Transport) InferencePutAzureaistudio {
	_ = "STUB: not implemented"
	return *new(InferencePutAzureaistudio)
}

type InferencePutAzureaistudio func(body io.Reader, azureaistudio_inference_id string, task_type string, o ...func(*InferencePutAzureaistudioRequest)) (*Response, error)

type InferencePutAzureaistudioRequest struct {
	Body io.Reader

	AzureaistudioInferenceID string
	TaskType                 string

	Timeout time.Duration

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r InferencePutAzureaistudioRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f InferencePutAzureaistudio) WithContext(v context.Context) func(*InferencePutAzureaistudioRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferencePutAzureaistudio) WithTimeout(v time.Duration) func(*InferencePutAzureaistudioRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferencePutAzureaistudio) WithPretty() func(*InferencePutAzureaistudioRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferencePutAzureaistudio) WithHuman() func(*InferencePutAzureaistudioRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferencePutAzureaistudio) WithErrorTrace() func(*InferencePutAzureaistudioRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferencePutAzureaistudio) WithFilterPath(v ...string) func(*InferencePutAzureaistudioRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferencePutAzureaistudio) WithHeader(h map[string]string) func(*InferencePutAzureaistudioRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferencePutAzureaistudio) WithOpaqueID(s string) func(*InferencePutAzureaistudioRequest) {
	_ = "STUB: not implemented"
	return nil
}
