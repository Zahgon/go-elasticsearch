package esapi

import (
	"context"
	"io"
	"net/http"
)

func newInferenceUpdateFunc(t Transport) InferenceUpdate {
	_ = "STUB: not implemented"
	return *new(InferenceUpdate)
}

type InferenceUpdate func(body io.Reader, inference_id string, o ...func(*InferenceUpdateRequest)) (*Response, error)

type InferenceUpdateRequest struct {
	Body io.Reader

	InferenceID string
	TaskType    string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r InferenceUpdateRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f InferenceUpdate) WithContext(v context.Context) func(*InferenceUpdateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferenceUpdate) WithTaskType(v string) func(*InferenceUpdateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferenceUpdate) WithPretty() func(*InferenceUpdateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferenceUpdate) WithHuman() func(*InferenceUpdateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferenceUpdate) WithErrorTrace() func(*InferenceUpdateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferenceUpdate) WithFilterPath(v ...string) func(*InferenceUpdateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferenceUpdate) WithHeader(h map[string]string) func(*InferenceUpdateRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferenceUpdate) WithOpaqueID(s string) func(*InferenceUpdateRequest) {
	_ = "STUB: not implemented"
	return nil
}
