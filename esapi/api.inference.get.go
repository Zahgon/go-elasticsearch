package esapi

import (
	"context"
	"net/http"
)

func newInferenceGetFunc(t Transport) InferenceGet {
	_ = "STUB: not implemented"
	return *new(InferenceGet)
}

type InferenceGet func(o ...func(*InferenceGetRequest)) (*Response, error)

type InferenceGetRequest struct {
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

func (r InferenceGetRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f InferenceGet) WithContext(v context.Context) func(*InferenceGetRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferenceGet) WithInferenceID(v string) func(*InferenceGetRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferenceGet) WithTaskType(v string) func(*InferenceGetRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferenceGet) WithPretty() func(*InferenceGetRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferenceGet) WithHuman() func(*InferenceGetRequest) { _ = "STUB: not implemented"; return nil }

func (f InferenceGet) WithErrorTrace() func(*InferenceGetRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferenceGet) WithFilterPath(v ...string) func(*InferenceGetRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferenceGet) WithHeader(h map[string]string) func(*InferenceGetRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferenceGet) WithOpaqueID(s string) func(*InferenceGetRequest) {
	_ = "STUB: not implemented"
	return nil
}
