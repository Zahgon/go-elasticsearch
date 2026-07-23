package esapi

import (
	"context"
	"io"
	"net/http"
	"time"
)

func newInferencePutAmazonsagemakerFunc(t Transport) InferencePutAmazonsagemaker {
	_ = "STUB: not implemented"
	return *new(InferencePutAmazonsagemaker)
}

type InferencePutAmazonsagemaker func(body io.Reader, amazonsagemaker_inference_id string, task_type string, o ...func(*InferencePutAmazonsagemakerRequest)) (*Response, error)

type InferencePutAmazonsagemakerRequest struct {
	Body io.Reader

	AmazonsagemakerInferenceID string
	TaskType                   string

	Timeout time.Duration

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r InferencePutAmazonsagemakerRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f InferencePutAmazonsagemaker) WithContext(v context.Context) func(*InferencePutAmazonsagemakerRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferencePutAmazonsagemaker) WithTimeout(v time.Duration) func(*InferencePutAmazonsagemakerRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferencePutAmazonsagemaker) WithPretty() func(*InferencePutAmazonsagemakerRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferencePutAmazonsagemaker) WithHuman() func(*InferencePutAmazonsagemakerRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferencePutAmazonsagemaker) WithErrorTrace() func(*InferencePutAmazonsagemakerRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferencePutAmazonsagemaker) WithFilterPath(v ...string) func(*InferencePutAmazonsagemakerRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferencePutAmazonsagemaker) WithHeader(h map[string]string) func(*InferencePutAmazonsagemakerRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferencePutAmazonsagemaker) WithOpaqueID(s string) func(*InferencePutAmazonsagemakerRequest) {
	_ = "STUB: not implemented"
	return nil
}
