package esapi

import (
	"context"
	"io"
	"net/http"
	"time"
)

func newInferencePutCohereFunc(t Transport) InferencePutCohere {
	_ = "STUB: not implemented"
	return *new(InferencePutCohere)
}

type InferencePutCohere func(body io.Reader, cohere_inference_id string, task_type string, o ...func(*InferencePutCohereRequest)) (*Response, error)

type InferencePutCohereRequest struct {
	Body io.Reader

	CohereInferenceID string
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

func (r InferencePutCohereRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f InferencePutCohere) WithContext(v context.Context) func(*InferencePutCohereRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferencePutCohere) WithTimeout(v time.Duration) func(*InferencePutCohereRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferencePutCohere) WithPretty() func(*InferencePutCohereRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferencePutCohere) WithHuman() func(*InferencePutCohereRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferencePutCohere) WithErrorTrace() func(*InferencePutCohereRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferencePutCohere) WithFilterPath(v ...string) func(*InferencePutCohereRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferencePutCohere) WithHeader(h map[string]string) func(*InferencePutCohereRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferencePutCohere) WithOpaqueID(s string) func(*InferencePutCohereRequest) {
	_ = "STUB: not implemented"
	return nil
}
