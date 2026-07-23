package esapi

import (
	"context"
	"io"
	"net/http"
	"time"
)

func newInferencePutOpenshiftAiFunc(t Transport) InferencePutOpenshiftAi {
	_ = "STUB: not implemented"
	return *new(InferencePutOpenshiftAi)
}

type InferencePutOpenshiftAi func(body io.Reader, openshiftai_inference_id string, task_type string, o ...func(*InferencePutOpenshiftAiRequest)) (*Response, error)

type InferencePutOpenshiftAiRequest struct {
	Body io.Reader

	OpenshiftaiInferenceID string
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

func (r InferencePutOpenshiftAiRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f InferencePutOpenshiftAi) WithContext(v context.Context) func(*InferencePutOpenshiftAiRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferencePutOpenshiftAi) WithTimeout(v time.Duration) func(*InferencePutOpenshiftAiRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferencePutOpenshiftAi) WithPretty() func(*InferencePutOpenshiftAiRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferencePutOpenshiftAi) WithHuman() func(*InferencePutOpenshiftAiRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferencePutOpenshiftAi) WithErrorTrace() func(*InferencePutOpenshiftAiRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferencePutOpenshiftAi) WithFilterPath(v ...string) func(*InferencePutOpenshiftAiRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferencePutOpenshiftAi) WithHeader(h map[string]string) func(*InferencePutOpenshiftAiRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferencePutOpenshiftAi) WithOpaqueID(s string) func(*InferencePutOpenshiftAiRequest) {
	_ = "STUB: not implemented"
	return nil
}
