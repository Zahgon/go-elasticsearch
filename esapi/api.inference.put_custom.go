package esapi

import (
	"context"
	"io"
	"net/http"
)

func newInferencePutCustomFunc(t Transport) InferencePutCustom {
	_ = "STUB: not implemented"
	return *new(InferencePutCustom)
}

type InferencePutCustom func(body io.Reader, custom_inference_id string, task_type string, o ...func(*InferencePutCustomRequest)) (*Response, error)

type InferencePutCustomRequest struct {
	Body io.Reader

	CustomInferenceID string
	TaskType          string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r InferencePutCustomRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f InferencePutCustom) WithContext(v context.Context) func(*InferencePutCustomRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferencePutCustom) WithPretty() func(*InferencePutCustomRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferencePutCustom) WithHuman() func(*InferencePutCustomRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferencePutCustom) WithErrorTrace() func(*InferencePutCustomRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferencePutCustom) WithFilterPath(v ...string) func(*InferencePutCustomRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferencePutCustom) WithHeader(h map[string]string) func(*InferencePutCustomRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferencePutCustom) WithOpaqueID(s string) func(*InferencePutCustomRequest) {
	_ = "STUB: not implemented"
	return nil
}
