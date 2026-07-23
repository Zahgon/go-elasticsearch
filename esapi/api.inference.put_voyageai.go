package esapi

import (
	"context"
	"io"
	"net/http"
	"time"
)

func newInferencePutVoyageaiFunc(t Transport) InferencePutVoyageai {
	_ = "STUB: not implemented"
	return *new(InferencePutVoyageai)
}

type InferencePutVoyageai func(body io.Reader, task_type string, voyageai_inference_id string, o ...func(*InferencePutVoyageaiRequest)) (*Response, error)

type InferencePutVoyageaiRequest struct {
	Body io.Reader

	TaskType            string
	VoyageaiInferenceID string

	Timeout time.Duration

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r InferencePutVoyageaiRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f InferencePutVoyageai) WithContext(v context.Context) func(*InferencePutVoyageaiRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferencePutVoyageai) WithTimeout(v time.Duration) func(*InferencePutVoyageaiRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferencePutVoyageai) WithPretty() func(*InferencePutVoyageaiRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferencePutVoyageai) WithHuman() func(*InferencePutVoyageaiRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferencePutVoyageai) WithErrorTrace() func(*InferencePutVoyageaiRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferencePutVoyageai) WithFilterPath(v ...string) func(*InferencePutVoyageaiRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferencePutVoyageai) WithHeader(h map[string]string) func(*InferencePutVoyageaiRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f InferencePutVoyageai) WithOpaqueID(s string) func(*InferencePutVoyageaiRequest) {
	_ = "STUB: not implemented"
	return nil
}
