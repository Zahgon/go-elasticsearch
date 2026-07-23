package esapi

import (
	"context"
	"io"
	"net/http"
	"time"
)

func newInferencePutOpenaiFunc(t Transport) InferencePutOpenai {
	_ = "STUB: not implemented"
	return *new(InferencePutOpenai)
}

type InferencePutOpenai func(body io.Reader, openai_inference_id string, task_type string, o ...func(*InferencePutOpenaiRequest)) (*Response, error)

type InferencePutOpenaiRequest struct {
	Body io.Reader

	OpenaiInferenceID string
	TaskType          string

	Timeout time.Duration

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r InferencePutOpenaiRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f InferencePutOpenai) WithContext(v context.Context) func(*InferencePutOpenaiRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferencePutOpenai) WithTimeout(v time.Duration) func(*InferencePutOpenaiRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferencePutOpenai) WithPretty() func(*InferencePutOpenaiRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferencePutOpenai) WithHuman() func(*InferencePutOpenaiRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferencePutOpenai) WithErrorTrace() func(*InferencePutOpenaiRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferencePutOpenai) WithFilterPath(v ...string) func(*InferencePutOpenaiRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferencePutOpenai) WithHeader(h map[string]string) func(*InferencePutOpenaiRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferencePutOpenai) WithOpaqueID(s string) func(*InferencePutOpenaiRequest) {
	_ = "STUB: not implemented"
	return nil
}
