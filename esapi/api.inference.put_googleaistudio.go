package esapi

import (
	"context"
	"io"
	"net/http"
	"time"
)

func newInferencePutGoogleaistudioFunc(t Transport) InferencePutGoogleaistudio {
	_ = "STUB: not implemented"
	return *new(InferencePutGoogleaistudio)
}

type InferencePutGoogleaistudio func(body io.Reader, googleaistudio_inference_id string, task_type string, o ...func(*InferencePutGoogleaistudioRequest)) (*Response, error)

type InferencePutGoogleaistudioRequest struct {
	Body io.Reader

	GoogleaistudioInferenceID string
	TaskType                  string

	Timeout time.Duration

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r InferencePutGoogleaistudioRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f InferencePutGoogleaistudio) WithContext(v context.Context) func(*InferencePutGoogleaistudioRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferencePutGoogleaistudio) WithTimeout(v time.Duration) func(*InferencePutGoogleaistudioRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferencePutGoogleaistudio) WithPretty() func(*InferencePutGoogleaistudioRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferencePutGoogleaistudio) WithHuman() func(*InferencePutGoogleaistudioRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferencePutGoogleaistudio) WithErrorTrace() func(*InferencePutGoogleaistudioRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferencePutGoogleaistudio) WithFilterPath(v ...string) func(*InferencePutGoogleaistudioRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferencePutGoogleaistudio) WithHeader(h map[string]string) func(*InferencePutGoogleaistudioRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferencePutGoogleaistudio) WithOpaqueID(s string) func(*InferencePutGoogleaistudioRequest) {
	_ = "STUB: not implemented"
	return nil
}
