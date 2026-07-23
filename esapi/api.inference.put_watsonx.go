package esapi

import (
	"context"
	"io"
	"net/http"
	"time"
)

func newInferencePutWatsonxFunc(t Transport) InferencePutWatsonx {
	_ = "STUB: not implemented"
	return *new(InferencePutWatsonx)
}

type InferencePutWatsonx func(body io.Reader, task_type string, watsonx_inference_id string, o ...func(*InferencePutWatsonxRequest)) (*Response, error)

type InferencePutWatsonxRequest struct {
	Body io.Reader

	TaskType           string
	WatsonxInferenceID string

	Timeout time.Duration

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r InferencePutWatsonxRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f InferencePutWatsonx) WithContext(v context.Context) func(*InferencePutWatsonxRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferencePutWatsonx) WithTimeout(v time.Duration) func(*InferencePutWatsonxRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferencePutWatsonx) WithPretty() func(*InferencePutWatsonxRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferencePutWatsonx) WithHuman() func(*InferencePutWatsonxRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferencePutWatsonx) WithErrorTrace() func(*InferencePutWatsonxRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferencePutWatsonx) WithFilterPath(v ...string) func(*InferencePutWatsonxRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferencePutWatsonx) WithHeader(h map[string]string) func(*InferencePutWatsonxRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferencePutWatsonx) WithOpaqueID(s string) func(*InferencePutWatsonxRequest) {
	_ = "STUB: not implemented"
	return nil
}
