package esapi

import (
	"context"
	"net/http"
)

func newInferenceDeleteFunc(t Transport) InferenceDelete {
	_ = "STUB: not implemented"
	return *new(InferenceDelete)
}

type InferenceDelete func(inference_id string, o ...func(*InferenceDeleteRequest)) (*Response, error)

type InferenceDeleteRequest struct {
	InferenceID string
	TaskType    string

	DryRun *bool
	Force  *bool

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r InferenceDeleteRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f InferenceDelete) WithContext(v context.Context) func(*InferenceDeleteRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferenceDelete) WithTaskType(v string) func(*InferenceDeleteRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferenceDelete) WithDryRun(v bool) func(*InferenceDeleteRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferenceDelete) WithForce(v bool) func(*InferenceDeleteRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferenceDelete) WithPretty() func(*InferenceDeleteRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferenceDelete) WithHuman() func(*InferenceDeleteRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferenceDelete) WithErrorTrace() func(*InferenceDeleteRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferenceDelete) WithFilterPath(v ...string) func(*InferenceDeleteRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferenceDelete) WithHeader(h map[string]string) func(*InferenceDeleteRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferenceDelete) WithOpaqueID(s string) func(*InferenceDeleteRequest) {
	_ = "STUB: not implemented"
	return nil
}
