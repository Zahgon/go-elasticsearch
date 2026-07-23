package esapi

import (
	"context"
	"io"
	"net/http"
	"time"
)

func newInferencePutFunc(t Transport) InferencePut {
	_ = "STUB: not implemented"
	return *new(InferencePut)
}

type InferencePut func(body io.Reader, inference_id string, o ...func(*InferencePutRequest)) (*Response, error)

type InferencePutRequest struct {
	Body io.Reader

	InferenceID string
	TaskType    string

	Timeout time.Duration

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r InferencePutRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f InferencePut) WithContext(v context.Context) func(*InferencePutRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferencePut) WithTaskType(v string) func(*InferencePutRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferencePut) WithTimeout(v time.Duration) func(*InferencePutRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferencePut) WithPretty() func(*InferencePutRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferencePut) WithHuman() func(*InferencePutRequest) { _ = "STUB: not implemented"; return nil }

func (f InferencePut) WithErrorTrace() func(*InferencePutRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferencePut) WithFilterPath(v ...string) func(*InferencePutRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferencePut) WithHeader(h map[string]string) func(*InferencePutRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferencePut) WithOpaqueID(s string) func(*InferencePutRequest) {
	_ = "STUB: not implemented"
	return nil
}
