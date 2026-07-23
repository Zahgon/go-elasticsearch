package esapi

import (
	"context"
	"io"
	"net/http"
	"time"
)

func newInferencePutFireworksaiFunc(t Transport) InferencePutFireworksai {
	_ = "STUB: not implemented"
	return *new(InferencePutFireworksai)
}

type InferencePutFireworksai func(body io.Reader, fireworksai_inference_id string, task_type string, o ...func(*InferencePutFireworksaiRequest)) (*Response, error)

type InferencePutFireworksaiRequest struct {
	Body io.Reader

	FireworksaiInferenceID string
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

func (r InferencePutFireworksaiRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f InferencePutFireworksai) WithContext(v context.Context) func(*InferencePutFireworksaiRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferencePutFireworksai) WithTimeout(v time.Duration) func(*InferencePutFireworksaiRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferencePutFireworksai) WithPretty() func(*InferencePutFireworksaiRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferencePutFireworksai) WithHuman() func(*InferencePutFireworksaiRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferencePutFireworksai) WithErrorTrace() func(*InferencePutFireworksaiRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferencePutFireworksai) WithFilterPath(v ...string) func(*InferencePutFireworksaiRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferencePutFireworksai) WithHeader(h map[string]string) func(*InferencePutFireworksaiRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferencePutFireworksai) WithOpaqueID(s string) func(*InferencePutFireworksaiRequest) {
	_ = "STUB: not implemented"
	return nil
}
