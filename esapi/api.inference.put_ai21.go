package esapi

import (
	"context"
	"io"
	"net/http"
	"time"
)

func newInferencePutAi21Func(t Transport) InferencePutAi21 {
	_ = "STUB: not implemented"
	return *new(InferencePutAi21)
}

type InferencePutAi21 func(body io.Reader, ai21_inference_id string, task_type string, o ...func(*InferencePutAi21Request)) (*Response, error)

type InferencePutAi21Request struct {
	Body io.Reader

	Ai21InferenceID string
	TaskType        string

	Timeout time.Duration

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r InferencePutAi21Request) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f InferencePutAi21) WithContext(v context.Context) func(*InferencePutAi21Request) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferencePutAi21) WithTimeout(v time.Duration) func(*InferencePutAi21Request) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferencePutAi21) WithPretty() func(*InferencePutAi21Request) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferencePutAi21) WithHuman() func(*InferencePutAi21Request) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferencePutAi21) WithErrorTrace() func(*InferencePutAi21Request) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferencePutAi21) WithFilterPath(v ...string) func(*InferencePutAi21Request) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferencePutAi21) WithHeader(h map[string]string) func(*InferencePutAi21Request) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferencePutAi21) WithOpaqueID(s string) func(*InferencePutAi21Request) {
	_ = "STUB: not implemented"
	return nil
}
