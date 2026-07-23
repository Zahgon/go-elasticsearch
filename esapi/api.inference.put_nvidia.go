package esapi

import (
	"context"
	"io"
	"net/http"
	"time"
)

func newInferencePutNvidiaFunc(t Transport) InferencePutNvidia {
	_ = "STUB: not implemented"
	return *new(InferencePutNvidia)
}

type InferencePutNvidia func(body io.Reader, nvidia_inference_id string, task_type string, o ...func(*InferencePutNvidiaRequest)) (*Response, error)

type InferencePutNvidiaRequest struct {
	Body io.Reader

	NvidiaInferenceID string
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

func (r InferencePutNvidiaRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f InferencePutNvidia) WithContext(v context.Context) func(*InferencePutNvidiaRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferencePutNvidia) WithTimeout(v time.Duration) func(*InferencePutNvidiaRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferencePutNvidia) WithPretty() func(*InferencePutNvidiaRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferencePutNvidia) WithHuman() func(*InferencePutNvidiaRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferencePutNvidia) WithErrorTrace() func(*InferencePutNvidiaRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferencePutNvidia) WithFilterPath(v ...string) func(*InferencePutNvidiaRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferencePutNvidia) WithHeader(h map[string]string) func(*InferencePutNvidiaRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferencePutNvidia) WithOpaqueID(s string) func(*InferencePutNvidiaRequest) {
	_ = "STUB: not implemented"
	return nil
}
