package esapi

import (
	"context"
	"io"
	"net/http"
)

func newMLEstimateModelMemoryFunc(t Transport) MLEstimateModelMemory {
	_ = "STUB: not implemented"
	return *new(MLEstimateModelMemory)
}

type MLEstimateModelMemory func(body io.Reader, o ...func(*MLEstimateModelMemoryRequest)) (*Response, error)

type MLEstimateModelMemoryRequest struct {
	Body io.Reader

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

func (r MLEstimateModelMemoryRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f MLEstimateModelMemory) WithContext(v context.Context) func(*MLEstimateModelMemoryRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLEstimateModelMemory) WithPretty() func(*MLEstimateModelMemoryRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLEstimateModelMemory) WithHuman() func(*MLEstimateModelMemoryRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLEstimateModelMemory) WithErrorTrace() func(*MLEstimateModelMemoryRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLEstimateModelMemory) WithFilterPath(v ...string) func(*MLEstimateModelMemoryRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLEstimateModelMemory) WithHeader(h map[string]string) func(*MLEstimateModelMemoryRequest) {
	_ = "STUB: not implemented"
	return nil
}

func (f MLEstimateModelMemory) WithOpaqueID(s string) func(*MLEstimateModelMemoryRequest) {
	_ = "STUB: not implemented"
	return nil
}
