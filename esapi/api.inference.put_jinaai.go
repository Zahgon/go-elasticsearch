package esapi

import (
	"context"
	"io"
	"net/http"
	"time"
)

func newInferencePutJinaaiFunc(t Transport) InferencePutJinaai {
	_ = "STUB: not implemented"
	return *new(InferencePutJinaai)
}

type InferencePutJinaai func(body io.Reader, jinaai_inference_id string, task_type string, o ...func(*InferencePutJinaaiRequest)) (*Response, error)

type InferencePutJinaaiRequest struct {
	Body io.Reader

	JinaaiInferenceID string
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

func (r InferencePutJinaaiRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f InferencePutJinaai) WithContext(v context.Context) func(*InferencePutJinaaiRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferencePutJinaai) WithTimeout(v time.Duration) func(*InferencePutJinaaiRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferencePutJinaai) WithPretty() func(*InferencePutJinaaiRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferencePutJinaai) WithHuman() func(*InferencePutJinaaiRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferencePutJinaai) WithErrorTrace() func(*InferencePutJinaaiRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferencePutJinaai) WithFilterPath(v ...string) func(*InferencePutJinaaiRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferencePutJinaai) WithHeader(h map[string]string) func(*InferencePutJinaaiRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferencePutJinaai) WithOpaqueID(s string) func(*InferencePutJinaaiRequest) {
	_ = "STUB: not implemented"
	return nil
}
