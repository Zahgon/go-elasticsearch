package esapi

import (
	"context"
	"io"
	"net/http"
	"time"
)

func newInferencePutElserFunc(t Transport) InferencePutElser {
	_ = "STUB: not implemented"
	return *new(InferencePutElser)
}

type InferencePutElser func(body io.Reader, elser_inference_id string, task_type string, o ...func(*InferencePutElserRequest)) (*Response, error)

type InferencePutElserRequest struct {
	Body io.Reader

	ElserInferenceID string
	TaskType         string

	Timeout time.Duration

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r InferencePutElserRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f InferencePutElser) WithContext(v context.Context) func(*InferencePutElserRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferencePutElser) WithTimeout(v time.Duration) func(*InferencePutElserRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferencePutElser) WithPretty() func(*InferencePutElserRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferencePutElser) WithHuman() func(*InferencePutElserRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferencePutElser) WithErrorTrace() func(*InferencePutElserRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferencePutElser) WithFilterPath(v ...string) func(*InferencePutElserRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferencePutElser) WithHeader(h map[string]string) func(*InferencePutElserRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferencePutElser) WithOpaqueID(s string) func(*InferencePutElserRequest) {
	_ = "STUB: not implemented"
	return nil
}
