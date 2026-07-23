package esapi

import (
	"context"
	"io"
	"net/http"
	"time"
)

func newInferencePutGooglevertexaiFunc(t Transport) InferencePutGooglevertexai {
	_ = "STUB: not implemented"
	return *new(InferencePutGooglevertexai)
}

type InferencePutGooglevertexai func(body io.Reader, googlevertexai_inference_id string, task_type string, o ...func(*InferencePutGooglevertexaiRequest)) (*Response, error)

type InferencePutGooglevertexaiRequest struct {
	Body io.Reader

	GooglevertexaiInferenceID string
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

func (r InferencePutGooglevertexaiRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f InferencePutGooglevertexai) WithContext(v context.Context) func(*InferencePutGooglevertexaiRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferencePutGooglevertexai) WithTimeout(v time.Duration) func(*InferencePutGooglevertexaiRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferencePutGooglevertexai) WithPretty() func(*InferencePutGooglevertexaiRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferencePutGooglevertexai) WithHuman() func(*InferencePutGooglevertexaiRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferencePutGooglevertexai) WithErrorTrace() func(*InferencePutGooglevertexaiRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferencePutGooglevertexai) WithFilterPath(v ...string) func(*InferencePutGooglevertexaiRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferencePutGooglevertexai) WithHeader(h map[string]string) func(*InferencePutGooglevertexaiRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferencePutGooglevertexai) WithOpaqueID(s string) func(*InferencePutGooglevertexaiRequest) {
	_ = "STUB: not implemented"
	return nil
}
